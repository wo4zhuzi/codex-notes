package main

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
	"github.com/openai/openai-go/v3/shared/constant"
)

type recordingResponsesClient struct {
	requests []responses.ResponseNewParams
}

func (c *recordingResponsesClient) New(
	_ context.Context,
	params responses.ResponseNewParams,
	_ ...option.RequestOption,
) (*responses.Response, error) {
	c.requests = append(c.requests, params)

	if len(c.requests) == 1 {
		return mustResponse(tResponseFunctionCall), nil
	}

	return &responses.Response{
		ID: "resp_final",
		Output: []responses.ResponseOutputItemUnion{{
			Type:   "message",
			Role:   constant.Assistant("assistant"),
			Status: "completed",
			Content: []responses.ResponseOutputMessageContentUnion{{
				Type: "output_text",
				Text: "订单 ORD-1001 已发货。",
			}},
		}},
	}, nil
}

const tResponseFunctionCall = `{
  "id": "resp_first",
  "object": "response",
  "output": [
    {
      "type": "function_call",
      "call_id": "call_123",
      "name": "get_order_status",
      "arguments": "{\"order_id\":\"ORD-1001\"}"
    }
  ]
}`

func mustResponse(raw string) *responses.Response {
	var response responses.Response
	if err := json.Unmarshal([]byte(raw), &response); err != nil {
		panic(err)
	}
	return &response
}

func TestRunSendsFunctionOutputWithoutPreviousResponseID(t *testing.T) {
	client := &recordingResponsesClient{}
	assistant := OrderAssistant{
		client: client,
		model:  defaultModel,
		store:  NewOrderStore("orders.json"),
	}

	answer, err := assistant.Run("帮我查一下订单 ORD-1001 到哪了")
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if answer != "订单 ORD-1001 已发货。" {
		t.Fatalf("answer = %q, want 订单 ORD-1001 已发货。", answer)
	}
	if len(client.requests) != 2 {
		t.Fatalf("request count = %d, want 2", len(client.requests))
	}

	followUp := client.requests[1]
	if followUp.PreviousResponseID.Value != "" {
		t.Fatalf("PreviousResponseID = %q, want empty", followUp.PreviousResponseID.Value)
	}

	items := followUp.Input.OfInputItemList
	if len(items) != 3 {
		t.Fatalf("follow-up input item count = %d, want 3", len(items))
	}
	if items[0].OfMessage == nil {
		t.Fatal("expected follow-up input to include original user message")
	}
	if items[1].OfFunctionCall == nil {
		t.Fatal("expected follow-up input to include function_call item")
	}
	if items[2].OfFunctionCallOutput == nil {
		t.Fatal("expected follow-up input to include function_call_output item")
	}

	outputItemJSON, err := items[2].MarshalJSON()
	if err != nil {
		t.Fatalf("marshal function_call_output: %v", err)
	}
	var outputItem struct {
		CallID string `json:"call_id"`
		Output string `json:"output"`
	}
	if err := json.Unmarshal(outputItemJSON, &outputItem); err != nil {
		t.Fatalf("unmarshal function_call_output: %v; json=%s", err, outputItemJSON)
	}
	if outputItem.CallID != "call_123" {
		t.Fatalf("function output call_id = %q, want call_123; json=%s", outputItem.CallID, outputItemJSON)
	}
	if !strings.Contains(outputItem.Output, `"order_id":"ORD-1001"`) {
		t.Fatalf("function output = %q, want order JSON", outputItem.Output)
	}
}

func TestCallFunctionReturnsOrderJSON(t *testing.T) {
	assistant := OrderAssistant{
		store: NewOrderStore("orders.json"),
	}

	output, err := assistant.callFunction(toolName, `{"order_id":"ORD-1001"}`)
	if err != nil {
		t.Fatalf("callFunction returned error: %v", err)
	}

	var result OrderLookupResult
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}
	if !result.Found {
		t.Fatal("expected order to be found")
	}
	if result.OrderID != "ORD-1001" {
		t.Fatalf("OrderID = %q, want ORD-1001", result.OrderID)
	}
}

func TestCallFunctionRejectsMissingOrderID(t *testing.T) {
	assistant := OrderAssistant{
		store: NewOrderStore("orders.json"),
	}

	_, err := assistant.callFunction(toolName, `{}`)
	if err == nil {
		t.Fatal("expected missing order_id error")
	}
	if !strings.Contains(err.Error(), "order_id") {
		t.Fatalf("error = %q, want mention order_id", err.Error())
	}
}
