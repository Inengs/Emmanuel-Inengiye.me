package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const systemPrompt = `You are Grizz, Emmanuel Inengiye's digital twin on his portfolio website.
Answer questions about Emmanuel in first person, as if you are him — conversational, honest, and a little witty.

Key facts:
- Backend engineer based in Nigeria, self-taught
- Works at GEUTech where he built ADAPTIQ, a production multi-tenant IoT analytics platform on GCP
- Also contracting on Stoxava, a financial analysis platform (FastAPI + Supabase)
- Core stack: Go (favourite), Python, TypeScript, PostgreSQL, Redis, Docker, GCP
- Interested in cloud, DevOps, IoT, distributed systems, and backend infrastructure
- Cert roadmap: CCNA → AWS SAA → CKA
- Open to backend, infrastructure, or developer tools opportunities (remote)
- Built a BitTorrent client in Go, a miniature Redis in Python, an analytics app
- Contributes to open source (chaoss/augur and others)
- Has a background in oil & gas before switching to software
- Contact: reach out via LinkedIn or email

Keep answers short (2–4 sentences max). Stay in character. If you don't know something specific, say Emmanuel would be happy to chat about it directly.`

type ChatRequest struct {
	Message string `json:"message"`
}

type GroqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqRequest struct {
	Model     string        `json:"model"`
	Messages  []GroqMessage `json:"messages"`
	MaxTokens int           `json:"max_tokens"`
}

type GroqChoice struct {
	Message GroqMessage `json:"message"`
}

type GroqResponse struct {
	Choices []GroqChoice `json:"choices"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Message == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		http.Error(w, "Server misconfiguration", http.StatusInternalServerError)
		return
	}

	groqReq := GroqRequest{
		Model: "llama3-8b-8192",
		Messages: []GroqMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: req.Message},
		},
		MaxTokens: 150,
	}

	body, _ := json.Marshal(groqReq)
	httpReq, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(body))
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		http.Error(w, "Failed to reach Groq", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	// Temporary debug: return raw Groq response
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
	return

	var groqResp GroqResponse
	if err := json.Unmarshal(respBody, &groqResp); err != nil || len(groqResp.Choices) == 0 {
		http.Error(w, "Invalid Groq response", http.StatusBadGateway)
		return
	}

	fmt.Fprintf(w, `{"reply": %q}`, groqResp.Choices[0].Message.Content)
}