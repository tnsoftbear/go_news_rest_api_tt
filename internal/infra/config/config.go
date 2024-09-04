package config

import (
	"fmt"
	"log"
	"os"

	"frr-news/internal/infra/env"

	"gopkg.in/yaml.v3"
)

type App struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Name         string `yaml:"name"`
	ServerHeader string `yaml:"server_header"`
}

type Jwt struct {
	Expiration string `yaml:"expiration"`
	Tokenkey   string `yaml:"tokenkey"`
}

type Auth struct {
	Jwt Jwt `yaml:"jwt"`
}

type MysqlStorage struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Config struct {
	App          App          `yaml:"app"`
	Auth         Auth         `yaml:"auth"`
	MysqlStorage MysqlStorage `yaml:"mysql_storage"`
}

func NewConfig(configPath string) (*Config, error) {
	config, err := loadFromYaml(&Config{}, configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file \"%s\", %s", configPath, err)
	}
	populateFromEnv(config)
	log.Printf("%v", config)
	return config, nil
}

func loadFromYaml(config *Config, configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func populateFromEnv(config *Config) {
	env.Setup()
	config.App.Host = env.GetStringEnv("APP_HOST", config.App.Host)
	config.App.Port = env.GetIntEnv("APP_PORT", config.App.Port)
	config.Auth.Jwt.Tokenkey = env.GetStringEnv("AUTH_JWT_TOKENKEY", config.Auth.Jwt.Tokenkey)
	config.MysqlStorage.Host = env.GetStringEnv("DB_HOST", config.MysqlStorage.Host)
	config.MysqlStorage.Port = env.GetIntEnv("DB_PORT", config.MysqlStorage.Port)
	config.MysqlStorage.Username = env.GetStringEnv("MYSQL_USER", config.MysqlStorage.Username)
	config.MysqlStorage.Password = env.GetStringEnv("MYSQL_PASSWORD", config.MysqlStorage.Password)
	config.MysqlStorage.Database = env.GetStringEnv("MYSQL_DATABASE", config.MysqlStorage.Database)
}
