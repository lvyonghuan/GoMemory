package store

import (
	"GoMemory/config"
	"GoMemory/proxy"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ReqsFromOpenai struct {
	Object string `json:"object"`
	Data   []Data `json:"data"`
	Model  string `json:"model"`
	Usage  Usage  `json:"usage"`
}
type Data struct {
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}
type Usage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type req struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

// OpenaiEmbedding 调用openai生成embedding
func openaiEmbedding(msg string) []float32 {
	client, err := proxy.Client()
	if err != nil {
		log.Println(err)
	}
	var request req
	request.Input = msg
	request.Model = "text-embedding-ada-002"
	data, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return nil
	}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/embeddings", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+config.Cfg.OpenaiApi)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var reps ReqsFromOpenai
	err = json.Unmarshal(bodyText, &reps)
	if err != nil || len(reps.Data) == 0 {
		log.Println("embedding生成错误:", err)
		return nil
	}
	return reps.Data[0].Embedding
}
