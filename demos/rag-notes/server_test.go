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
	answer Answer
	err    error
}

func (s stubRunner) Run(_ string) (Answer, error) {
	return s.answer, s.err
}

func TestChatHandlerReturnsAnswerAndSources(t *testing.T) {
	handler := routes(stubRunner{
		answer: Answer{
			Text: "RAG 会先检索再回答。",
			Sources: []SearchResult{{
				File:  "knowledge/rag.md",
				Title: "RAG 入门",
				Score: 8,
			}},
		},
	})
	body := bytes.NewBufferString(`{"message":"RAG 是什么？"}`)
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
	if payload.Answer != "RAG 会先检索再回答。" {
		t.Fatalf("Answer = %q, want RAG answer", payload.Answer)
	}
	if len(payload.Sources) != 1 {
		t.Fatalf("source count = %d, want 1", len(payload.Sources))
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
	body := bytes.NewBufferString(`{"message":"RAG 是什么？"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/chat", body)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusInternalServerError)
	}
}
