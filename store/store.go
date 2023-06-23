package store

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var file *os.File

type memory struct {
	msg  string
	flag string
}

func Init() {
	var err error
	file, err = os.Create(fmt.Sprintf("./store/log/%s.txt", time.Now().Format("2006-01-02_15_04-05")))
	if err != nil {
		log.Fatal(err)
	}
	wg = sync.WaitGroup{}
	read()
	log.Println("end read")
	wg.Wait()
	file.Close()
	log.Println("end store")
}

func read() {
	file, err := os.Open("./store/text")
	if err != nil {
		log.Fatal("open text err")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var flag string //消息类型确认，A代表记忆类型，2代表知识类型
	for i := 1; ; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if strings.TrimSpace(line) == "A" {
			flag = "memory"
		} else if strings.TrimSpace(line) == "B" {
			flag = "knowledge"
		} else if line == "end" {
			return
		} else {
			msg := memory{
				msg:  line,
				flag: flag,
			}
			wg.Add(1)
			go msg.store()
		}
	}
}

func (msg memory) store() {
	defer wg.Done()
	message := strings.Split(msg.msg, " ")
	embed := openaiEmbedding(message[0])
	if embed == nil {
		return
	}
	pineconeStore(message[1], msg.flag, "creator", "test1", embed)
}
