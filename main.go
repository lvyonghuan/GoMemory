package main

import (
	"GoMemory/config"
	"GoMemory/store"
)

func main() {
	config.ReadCfg()
	store.Init()
}
