package store

import (
	"io"
	"log"
	"sync"
)

var mu sync.Mutex

// 按id:text的格式写入log
func write(logMsg string) {
	mu.Lock()
	_, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		log.Println("写日志失败：" + err.Error())
		return
	}
	_, err = file.WriteString(logMsg)
	if err != nil {
		log.Println("写日志失败：" + err.Error())
		return
	}
	mu.Unlock()
}
