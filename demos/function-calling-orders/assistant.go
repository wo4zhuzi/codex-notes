package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

const (
	defaultModel = "gpt-4.1-mini"
	toolName     = "get_order_status"
)

var ErrMissingAPIKey = errors.New("缺少 OPENAI_API_KEY。请先执行：export OPENAI_API_KEY=\"你的 API Key\"")

type OrderAssistant struct {
	client OpenAIResponsesClient
	model  string
	store  OrderStore
}

type OpenAIResponsesClient interface {
	New(context.Context, responses.ResponseNewParams, ...option.RequestOption) (*responses.Response, error)
}

func NewOrderAssistant(store OrderStore) (OrderAssistant, error) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		return OrderAssistant{}, ErrMissingAPIKey
	}

	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = defaultModel
	}

	client := openai.NewClient()
	return OrderAssistant{
		client: &client.Responses,
		model:  model,
		store:  store,
	}, nil
}

func (a OrderAssistant) Run(question string) (string, error) {
	ctx := context.Background()
	tools := orderTools()

	response, err := a.client.New(ctx, responses.ResponseNewParams{
		Model: openai.ResponsesModel(a.model),
		Instructions: openai.String(
			"你是订单查询助手。需要订单实时状态时，必须调用 get_order_status，不要编造订单状态。",
		),
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String(question),
		},
		Tools: tools,
	})
	if err != nil {
		return "", fmt.Errorf("调用 OpenAI Responses API 失败：%w", err)
	}

	for _, item := range response.Output {
		if item.Type != "function_call" {
			continue
		}

		toolCall := item.AsFunctionCall()
		output, err := a.callFunction(toolCall.Name, toolCall.Arguments)
		if err != nil {
			return "", err
		}

		finalResponse, err := a.client.New(ctx, responses.ResponseNewParams{
			Model:              openai.ResponsesModel(a.model),
			PreviousResponseID: openai.String(response.ID),
			Input: responses.ResponseNewParamsInputUnion{
				OfInputItemList: []responses.ResponseInputItemUnionParam{{
					OfFunctionCallOutput: &responses.ResponseInputItemFunctionCallOutputParam{
						CallID: toolCall.CallID,
						Output: responses.ResponseInputItemFunctionCallOutputOutputUnionParam{
							OfString: openai.String(output),
						},
					},
				}},
			},
			Tools: tools,
		})
		if err != nil {
			return "", fmt.Errorf("回传函数执行结果失败：%w", err)
		}

		return finalResponse.OutputText(), nil
	}

	return response.OutputText(), nil
}

func (a OrderAssistant) callFunction(name string, rawArguments string) (string, error) {
	if name != toolName {
		result := map[string]string{
			"message": fmt.Sprintf("未知函数：%s", name),
		}
		output, _ := json.Marshal(result)
		return string(output), nil
	}

	var args struct {
		OrderID string `json:"order_id"`
	}
	if err := json.Unmarshal([]byte(rawArguments), &args); err != nil {
		return "", fmt.Errorf("解析函数参数失败：%w", err)
	}
	if args.OrderID == "" {
		return "", errors.New("函数参数缺少 order_id")
	}

	result, err := a.store.GetOrderStatus(args.OrderID)
	if err != nil {
		return "", err
	}

	output, err := json.Marshal(result)
	if err != nil {
		return "", fmt.Errorf("序列化订单查询结果失败：%w", err)
	}
	return string(output), nil
}

func orderTools() []responses.ToolUnionParam {
	return []responses.ToolUnionParam{{
		OfFunction: &responses.FunctionToolParam{
			Name:        toolName,
			Description: openai.String("根据订单号查询订单状态、物流信息、预计送达时间和订单商品。"),
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"order_id": map[string]string{
						"type":        "string",
						"description": "订单号，例如 ORD-1001。",
					},
				},
				"required":             []string{"order_id"},
				"additionalProperties": false,
			},
			Strict: openai.Bool(true),
		},
	}}
}
