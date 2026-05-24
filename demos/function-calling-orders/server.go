package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Answer string `json:"answer,omitempty"`
	Error  string `json:"error,omitempty"`
}

type AssistantRunner interface {
	Run(question string) (string, error)
}

func routes(runner AssistantRunner) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", indexHandler)
	mux.HandleFunc("POST /api/chat", chatHandler(runner))
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func chatHandler(runner AssistantRunner) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload ChatRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			writeJSON(w, http.StatusBadRequest, ChatResponse{Error: "请求 JSON 格式不正确。"})
			return
		}

		message := strings.TrimSpace(payload.Message)
		if message == "" {
			writeJSON(w, http.StatusBadRequest, ChatResponse{Error: "请输入要查询的问题。"})
			return
		}

		answer, err := runner.Run(message)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, ChatResponse{Error: err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, ChatResponse{Answer: answer})
	}
}

func writeJSON(w http.ResponseWriter, status int, payload ChatResponse) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
