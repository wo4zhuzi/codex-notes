package main

import "testing"

func TestSearchReturnsRelevantMarkdown(t *testing.T) {
	retriever, err := NewRetriever("knowledge")
	if err != nil {
		t.Fatalf("NewRetriever returned error: %v", err)
	}

	results := retriever.Search("Function Calling 和 RAG 有什么区别？", 3)
	if len(results) == 0 {
		t.Fatal("expected search results")
	}
	if results[0].File != "knowledge/function-calling.md" {
		t.Fatalf("top result = %q, want knowledge/function-calling.md", results[0].File)
	}
	if results[0].Score <= 0 {
		t.Fatalf("score = %d, want positive", results[0].Score)
	}
}

func TestSearchHonorsTopK(t *testing.T) {
	retriever, err := NewRetriever("knowledge")
	if err != nil {
		t.Fatalf("NewRetriever returned error: %v", err)
	}

	results := retriever.Search("Codex RAG MCP Function Calling", 2)
	if len(results) != 2 {
		t.Fatalf("result count = %d, want 2", len(results))
	}
}

func TestSearchRejectsEmptyQuery(t *testing.T) {
	retriever, err := NewRetriever("knowledge")
	if err != nil {
		t.Fatalf("NewRetriever returned error: %v", err)
	}

	results := retriever.Search("  ", 3)
	if len(results) != 0 {
		t.Fatalf("result count = %d, want 0", len(results))
	}
}

func TestQueryTermsIncludesASCIIAndChineseTerms(t *testing.T) {
	terms := queryTerms("RAG 如何检索 Markdown 文档？")
	termSet := map[string]bool{}
	for _, term := range terms {
		termSet[term] = true
	}

	for _, want := range []string{"rag", "如何", "检索", "markdown"} {
		if !termSet[want] {
			t.Fatalf("terms = %#v, want %q", terms, want)
		}
	}
}
