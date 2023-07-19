package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Messages struct {
	Start          string `mapstructure:"start"`
	UnknownCommand string `mapstructure:"unknown_command"`
}

type Db struct {
	Host     string `mapstructure:"start"`
	Port     string `mapstructure:"start"`
	Username string `mapstructure:"start"`
	Password string `mapstructure:"start"`
	DBName   string `mapstructure:"start"`
	SSLMode  string `mapstructure:"start"`
}

type Config struct {
	TelegramToken   string
	NewsResourceUrl string `mapstructure:"news_resource_url"`
	Messages        Messages
}

func Init() (*Config, error) {
	var cfg Config

	if err := setUpViper(); err != nil {
		return nil, err
	}

	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := fromEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages", &cfg.Messages); err != nil {
		return err
	}

	return nil
}

func fromEnv(cfg *Config) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg.TelegramToken = os.Getenv("TELEGRAM_API_KEY")

	return nil
}
