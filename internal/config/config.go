package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ObjectTTL struct {
		Video           int `yaml:"video" env-default:"10"`
		SearchResult    int `yaml:"search_result" env-default:"10"`
		VideoStreamPool int `yaml:"video_stream_pool" env-default:"10"`
	} `yaml:"object_ttl"`
	Redis struct {
		Host     string `yaml:"host" env-required:"true"`
		Port     string `yaml:"port" env-required:"true"`
		Username string `yaml:"username" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
		Database int    `yaml:"database" env-required:"true"`
	} `yaml:"redis" env-required:"true"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig(path string) (*Config, error) {
	var err error

	once.Do(func() {
		/*
			logger := logging.GetLogger()
			logger.Info("read application config")
		*/
		instance = &Config{}

		if len(path) == 0 {
			path = "../../config.yml"
		}

		err = cleanenv.ReadConfig(path, instance)

		/*if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
		*/
	})

	return instance, err
}
