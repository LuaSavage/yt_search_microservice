package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ObjectTTL struct {
		Video           string `yaml:"video" env-default:"10"`
		SearchResult    string `yaml:"search_result" env-default:"10"`
		VideoStreamPool string `yaml:"video_stream_pool" env-default:"10"`
	} `yaml:"object_ttl"`
	Redis struct {
		Host     string `yaml:"host" env-required:"true"`
		Port     string `yaml:"port" env-required:"true"`
		Username string `yaml:"username" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
		Database string `yaml:"database" env-required:"true"`
	} `yaml:"redis" env-required:"true"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() (*Config, error) {
	var err error

	once.Do(func() {
		/*
			logger := logging.GetLogger()
			logger.Info("read application config")
		*/
		instance = &Config{}
		err = cleanenv.ReadConfig("../../config.yml", instance)

		/*if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
		*/
	})

	return instance, err
}
