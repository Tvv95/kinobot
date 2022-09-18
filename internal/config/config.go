package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	TgToken       string `env:"TG_TOKEN"`
	KPConsumerKey string `env:"KP_CONSUMER_KEY"`
}

var cfg *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}
		if err := cleanenv.ReadConfig(".env", cfg); err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return cfg
}
