package main

import "testing"

func TestGetOrderStatusFound(t *testing.T) {
	store := NewOrderStore("orders.json")

	result, err := store.GetOrderStatus("ORD-1001")
	if err != nil {
		t.Fatalf("GetOrderStatus returned error: %v", err)
	}

	if !result.Found {
		t.Fatal("expected order to be found")
	}
	if result.OrderID != "ORD-1001" {
		t.Fatalf("OrderID = %q, want ORD-1001", result.OrderID)
	}
	if result.Status != "已发货" {
		t.Fatalf("Status = %q, want 已发货", result.Status)
	}
}

func TestGetOrderStatusMissing(t *testing.T) {
	store := NewOrderStore("orders.json")

	result, err := store.GetOrderStatus("ORD-404")
	if err != nil {
		t.Fatalf("GetOrderStatus returned error: %v", err)
	}

	if result.Found {
		t.Fatal("expected order to be missing")
	}
	if result.OrderID != "ORD-404" {
		t.Fatalf("OrderID = %q, want ORD-404", result.OrderID)
	}
	if result.Message != "订单不存在。" {
		t.Fatalf("Message = %q, want 订单不存在。", result.Message)
	}
}
