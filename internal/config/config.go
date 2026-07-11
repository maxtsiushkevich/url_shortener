package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database   DbConfig         `yaml:"database"`
	HTTPServer HTTPServerConfig `yaml:"http_server"`
	Redis      RedisConfig      `yaml:"redis"`
}

type DbConfig struct {
	Address  string `yaml:"db_address"`
	User     string `yaml:"db_user"`
	Password string `yaml:"db_pass"`
	DbName   string `yaml:"db_name"`
}

type HTTPServerConfig struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	User     string `yaml:"redis_user"`
	Password string `yaml:"redis_password"`
}

func Load(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	config := Config{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
