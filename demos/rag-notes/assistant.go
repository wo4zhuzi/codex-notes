package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

const (
	defaultModel        = "gpt-4.1-mini"
	defaultKnowledgeDir = "knowledge"
	defaultTopK         = 3
)

var ErrMissingAPIKey = errors.New("缺少 OPENAI_API_KEY。请先执行：export OPENAI_API_KEY=\"你的 API Key\"")

type Answer struct {
	Text    string         `json:"answer"`
	Sources []SearchResult `json:"sources"`
}

type RAGAssistant struct {
	client    OpenAIResponsesClient
	model     string
	retriever Retriever
	topK      int
}

type OpenAIResponsesClient interface {
	New(context.Context, responses.ResponseNewParams, ...option.RequestOption) (*responses.Response, error)
}

func NewRAGAssistant(retriever Retriever, topK int) (RAGAssistant, error) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		return RAGAssistant{}, ErrMissingAPIKey
	}

	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = defaultModel
	}

	client := openai.NewClient()
	return RAGAssistant{
		client:    &client.Responses,
		model:     model,
		retriever: retriever,
		topK:      normalizeTopK(topK),
	}, nil
}

func (a RAGAssistant) Run(question string) (Answer, error) {
	question = strings.TrimSpace(question)
	if question == "" {
		return Answer{}, errors.New("请输入问题")
	}

	// 先由 Go 检索本地知识库，模型只负责基于检索结果生成答案。
	sources := a.retriever.Search(question, a.topK)
	if len(sources) == 0 {
		return Answer{
			Text:    "知识库没有检索到相关资料，因此不基于模型常识编造答案。请换一个更具体的问题，或补充知识文档。",
			Sources: nil,
		}, nil
	}

	ctx := context.Background()
	response, err := a.client.New(ctx, responses.ResponseNewParams{
		Model: openai.ResponsesModel(a.model),
		Instructions: openai.String(
			"你是 RAG 知识库助手。只能基于用户问题下方提供的知识库上下文回答；如果上下文不足，直接说明不足。回答要简洁，并在末尾列出引用来源文件。",
		),
		Input: responses.ResponseNewParamsInputUnion{
			// 模型收到的是增强后的 prompt，不是原始 knowledge 目录。
			OfString: openai.String(buildPrompt(question, sources)),
		},
	})
	if err != nil {
		return Answer{}, fmt.Errorf("调用 OpenAI Responses API 失败：%w", err)
	}

	text := strings.TrimSpace(response.OutputText())
	if text == "" {
		text = "模型没有返回内容。"
	}
	return Answer{Text: text, Sources: sources}, nil
}

// Prompt Augmentation：把用户问题和 topK 检索结果拼成模型输入。
func buildPrompt(question string, sources []SearchResult) string {
	var b strings.Builder
	b.WriteString("用户问题：\n")
	b.WriteString(question)
	b.WriteString("\n\n知识库上下文：\n")

	for i, source := range sources {
		fmt.Fprintf(&b, "\n[%d] 来源：%s\n标题：%s\n内容：\n%s\n", i+1, source.File, source.Title, source.Content)
	}

	b.WriteString("\n回答要求：\n")
	b.WriteString("- 只使用上面的知识库上下文。\n")
	b.WriteString("- 如果资料不足，明确说明不足。\n")
	b.WriteString("- 末尾用“来源：”列出实际使用的文件路径。\n")
	return b.String()
}

func normalizeTopK(topK int) int {
	if topK <= 0 {
		return defaultTopK
	}
	if topK > 8 {
		return 8
	}
	return topK
}
