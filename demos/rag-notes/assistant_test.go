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
	return &responses.Response{
		ID: "resp_final",
		Output: []responses.ResponseOutputItemUnion{{
			Type:   "message",
			Role:   constant.Assistant("assistant"),
			Status: "completed",
			Content: []responses.ResponseOutputMessageContentUnion{{
				Type: "output_text",
				Text: "RAG 会先检索知识库，再基于上下文回答。\n\n来源：knowledge/rag.md",
			}},
		}},
	}, nil
}

func TestRunInjectsRetrievedContext(t *testing.T) {
	retriever, err := NewRetriever("knowledge")
	if err != nil {
		t.Fatalf("NewRetriever returned error: %v", err)
	}

	client := &recordingResponsesClient{}
	assistant := RAGAssistant{
		client:    client,
		model:     defaultModel,
		retriever: retriever,
		topK:      2,
	}

	answer, err := assistant.Run("RAG 是什么？")
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if !strings.Contains(answer.Text, "RAG") {
		t.Fatalf("answer = %q, want mention RAG", answer.Text)
	}
	if len(answer.Sources) == 0 {
		t.Fatal("expected sources")
	}
	if len(client.requests) != 1 {
		t.Fatalf("request count = %d, want 1", len(client.requests))
	}

	requestJSON, err := json.Marshal(client.requests[0])
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}
	requestText := string(requestJSON)
	if !strings.Contains(requestText, "知识库上下文") {
		t.Fatalf("request does not include RAG context: %s", requestText)
	}
	if !strings.Contains(requestText, "knowledge/rag.md") {
		t.Fatalf("request does not include source path: %s", requestText)
	}
}

func TestRunReturnsNoContextMessageWhenNoResult(t *testing.T) {
	assistant := RAGAssistant{
		client:    &recordingResponsesClient{},
		model:     defaultModel,
		retriever: Retriever{},
		topK:      2,
	}

	answer, err := assistant.Run("完全不存在的问题")
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if !strings.Contains(answer.Text, "没有检索到相关资料") {
		t.Fatalf("answer = %q, want no context message", answer.Text)
	}
	if len(answer.Sources) != 0 {
		t.Fatalf("sources count = %d, want 0", len(answer.Sources))
	}
}

func TestBuildPromptIncludesSources(t *testing.T) {
	prompt := buildPrompt("RAG 是什么？", []SearchResult{{
		File:    "knowledge/rag.md",
		Title:   "RAG 入门",
		Score:   5,
		Content: "RAG 会先检索再回答。",
	}})

	for _, want := range []string{"用户问题", "知识库上下文", "knowledge/rag.md", "RAG 会先检索再回答"} {
		if !strings.Contains(prompt, want) {
			t.Fatalf("prompt = %q, want %q", prompt, want)
		}
	}
}
