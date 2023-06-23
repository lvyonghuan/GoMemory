package main

import (
	"GoMemory/config"
	"GoMemory/proxy"
	"GoMemory/store"
)

func main() {
	config.ReadCfg()
	proxy.InitProxyConfig()
	store.Init()
}
