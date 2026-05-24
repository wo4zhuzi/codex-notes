package main

import (
	"encoding/json"
	"strings"
	"testing"
)

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
