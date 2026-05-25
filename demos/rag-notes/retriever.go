package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Document struct {
	File    string
	Title   string
	Content string
}

type SearchResult struct {
	File    string `json:"file"`
	Title   string `json:"title"`
	Score   int    `json:"score"`
	Content string `json:"content,omitempty"`
}

type Retriever struct {
	docs []Document
}

func NewRetriever(knowledgeDir string) (Retriever, error) {
	docs, err := loadDocuments(knowledgeDir)
	if err != nil {
		return Retriever{}, err
	}
	return Retriever{docs: docs}, nil
}

func loadDocuments(knowledgeDir string) ([]Document, error) {
	entries, err := os.ReadDir(knowledgeDir)
	if err != nil {
		return nil, fmt.Errorf("读取知识库目录失败：%w", err)
	}

	var docs []Document
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}

		path := filepath.Join(knowledgeDir, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("读取知识文档失败：%w", err)
		}

		content := strings.TrimSpace(string(data))
		if content == "" {
			continue
		}

		docs = append(docs, Document{
			File:    filepath.ToSlash(path),
			Title:   extractTitle(content, entry.Name()),
			Content: content,
		})
	}

	sort.Slice(docs, func(i, j int) bool {
		return docs[i].File < docs[j].File
	})
	return docs, nil
}

func extractTitle(content string, fallback string) string {
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "# "))
		}
	}
	return strings.TrimSuffix(fallback, filepath.Ext(fallback))
}

func (r Retriever) Search(query string, topK int) []SearchResult {
	query = strings.TrimSpace(query)
	if query == "" || topK <= 0 {
		return nil
	}

	terms := queryTerms(query)
	if len(terms) == 0 {
		return nil
	}

	var results []SearchResult
	for _, doc := range r.docs {
		score := scoreDocument(doc, terms)
		if score <= 0 {
			continue
		}

		results = append(results, SearchResult{
			File:    doc.File,
			Title:   doc.Title,
			Score:   score,
			Content: compactContent(doc.Content, 1400),
		})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Score == results[j].Score {
			return results[i].File < results[j].File
		}
		return results[i].Score > results[j].Score
	})

	if len(results) > topK {
		return results[:topK]
	}
	return results
}

func scoreDocument(doc Document, terms []string) int {
	file := strings.ToLower(doc.File)
	title := strings.ToLower(doc.Title)
	content := strings.ToLower(doc.Content)
	score := 0

	for _, term := range terms {
		if strings.Contains(file, term) {
			score += 4
		}
		if strings.Contains(title, term) {
			score += 5
		}
		score += strings.Count(content, term)
	}
	return score
}

func queryTerms(query string) []string {
	query = strings.ToLower(query)
	seen := map[string]bool{}
	var terms []string

	add := func(term string) {
		term = strings.TrimSpace(term)
		if term == "" || seen[term] {
			return
		}
		seen[term] = true
		terms = append(terms, term)
	}

	var ascii strings.Builder
	var cjk []rune
	flushASCII := func() {
		if ascii.Len() >= 2 {
			add(ascii.String())
		}
		ascii.Reset()
	}
	flushCJK := func() {
		if len(cjk) == 1 {
			add(string(cjk[0]))
		}
		for i := 0; i+1 < len(cjk); i++ {
			add(string(cjk[i : i+2]))
		}
		cjk = nil
	}

	for _, r := range query {
		switch {
		case isASCIIWord(r):
			flushCJK()
			ascii.WriteRune(r)
		case isCJK(r):
			flushASCII()
			cjk = append(cjk, r)
		default:
			flushASCII()
			flushCJK()
		}
	}
	flushASCII()
	flushCJK()

	return terms
}

func isASCIIWord(r rune) bool {
	return r <= unicode.MaxASCII && (unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_')
}

func isCJK(r rune) bool {
	return unicode.In(r, unicode.Han)
}

func compactContent(content string, maxRunes int) string {
	content = strings.TrimSpace(content)
	if utf8.RuneCountInString(content) <= maxRunes {
		return content
	}

	runes := []rune(content)
	return strings.TrimSpace(string(runes[:maxRunes])) + "\n..."
}
