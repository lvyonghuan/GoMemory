package config

import (
	"github.com/spf13/viper"
	"log"
)

type config struct {
	OpenaiApi   string `mapstructure:"openai_api"`
	PineconeApi string `mapstructure:"pinecone_api"`
	Url         string `mapstructure:"url"`
}

var Cfg config

func ReadCfg() {
	viper.SetConfigName("config.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
}
