package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Messages struct {
	Start          string `mapstructure:"start"`
	EditCategory   string `mapstructure:"edit_category"`
	UnknownCommand string `mapstructure:"unknown_command"`
}

type Db struct {
	Password string
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type News struct {
	ApiKey         string
	Url            string        `mapstructure:"url"`
	DefaultTimeout time.Duration `mapstructure:"defaultTimeout"`
}

type Rkn struct {
	Url            string        `mapstructure:"url"`
	DefaultTimeout time.Duration `mapstructure:"defaultTimeout"`
}

type Consumer struct {
	MailingTimeEnd int           `mapstructure:"mailingTimeEnd"`
	DailySleep     time.Duration `mapstructure:"dailySleep"`
	CategorySleep  time.Duration `mapstructure:"categorySleep"`
	RequestLimit   int8          `mapstructure:"requestLimit"`
}

type Config struct {
	TelegramToken string
	News          News
	Rkn           Rkn
	Db            Db
	Messages      Messages
	Consumer      Consumer
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

	if err := viper.UnmarshalKey("db", &cfg.Db); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("news", &cfg.News); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("rkn", &cfg.Rkn); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("consumer", &cfg.Consumer); err != nil {
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
	cfg.Db.Password = os.Getenv("DB_PASSWORD")
	cfg.News.ApiKey = os.Getenv("NEWS_API_KEY")

	return nil
}
