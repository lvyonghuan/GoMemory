package store

import (
	"GoMemory/config"
	"GoMemory/proxy"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
)

type VectorData struct {
	Vectors   []Vector `json:"vectors"`
	Namespace string   `json:"namespace"`
}

type Vector struct {
	ID       string         `json:"id"`
	Values   []float32      `json:"values"`
	Metadata map[string]any `json:"metadata"`
}

// StoreResp 储存返回
type StoreResp struct {
	UpsertedCount int64 `json:"upsertedCount"`
}

// PineconeStore 存储数据
func pineconeStore(text, typ, user, namespace string, vector []float32) {
	url := config.Cfg.Url + "/vectors/upsert"

	id := generateRandomString(10)
	if id == "" {
		return
	}
	data := VectorData{
		Vectors: []Vector{
			{
				ID:     id,
				Values: vector,
				Metadata: map[string]any{
					"Type": typ,
					"Text": text,
					"User": user,
				},
			},
		},
		Namespace: namespace,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		log.Println("json格式化错误,", err)
		return
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Api-Key", config.Cfg.PineconeApi)

	client, err := proxy.Client()
	if err != nil {
		log.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var resp StoreResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("pinecone upsert错误:", err)
	} else if resp.UpsertedCount < 1 {
		log.Println(string(body))
		log.Println("pinecone upsert错误:插入向量数量错误")
	}
	logMsg := id + ":" + text
	log.Println("存储成功：" + logMsg)
	write(logMsg)
}

// 随机生成10位的字符串用作id
func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println("生成vector id错误：", err)
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length]
}
