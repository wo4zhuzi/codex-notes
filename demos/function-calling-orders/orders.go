package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type OrderItem struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Order struct {
	OrderID           string      `json:"order_id"`
	Status            string      `json:"status"`
	Carrier           *string     `json:"carrier"`
	TrackingNumber    *string     `json:"tracking_number"`
	EstimatedDelivery *string     `json:"estimated_delivery"`
	UpdatedAt         string      `json:"updated_at"`
	Items             []OrderItem `json:"items"`
}

type OrderLookupResult struct {
	Found   bool   `json:"found"`
	Message string `json:"message,omitempty"`
	Order
}

type OrderStore struct {
	path string
}

func NewOrderStore(path string) OrderStore {
	return OrderStore{path: path}
}

func (s OrderStore) load() (map[string]Order, error) {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return nil, fmt.Errorf("读取订单文件失败：%w", err)
	}

	var orders map[string]Order
	if err := json.Unmarshal(data, &orders); err != nil {
		return nil, fmt.Errorf("解析订单文件失败：%w", err)
	}
	return orders, nil
}

func (s OrderStore) GetOrderStatus(orderID string) (OrderLookupResult, error) {
	orders, err := s.load()
	if err != nil {
		return OrderLookupResult{}, err
	}

	order, ok := orders[orderID]
	if !ok {
		return OrderLookupResult{
			Found:   false,
			Message: "订单不存在。",
			Order: Order{
				OrderID: orderID,
			},
		}, nil
	}

	return OrderLookupResult{
		Found: true,
		Order: order,
	}, nil
}
