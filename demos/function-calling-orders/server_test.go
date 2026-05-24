package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubRunner struct {
	answer string
	err    error
}

func (s stubRunner) Run(_ string) (string, error) {
	return s.answer, s.err
}

func TestChatHandlerReturnsAnswer(t *testing.T) {
	handler := routes(stubRunner{answer: "订单已发货。"})
	body := bytes.NewBufferString(`{"message":"帮我查一下订单 ORD-1001 到哪了"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/chat", body)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}

	var payload ChatResponse
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if payload.Answer != "订单已发货。" {
		t.Fatalf("Answer = %q, want 订单已发货。", payload.Answer)
	}
}

func TestChatHandlerRejectsEmptyMessage(t *testing.T) {
	handler := routes(stubRunner{})
	body := bytes.NewBufferString(`{"message":""}`)
	request := httptest.NewRequest(http.MethodPost, "/api/chat", body)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusBadRequest)
	}
}

func TestChatHandlerReturnsRunnerError(t *testing.T) {
	handler := routes(stubRunner{err: errors.New("缺少 OPENAI_API_KEY")})
	body := bytes.NewBufferString(`{"message":"查订单"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/chat", body)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusInternalServerError)
	}
}
