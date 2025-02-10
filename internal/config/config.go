package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
)

type Messages struct {
	Start          string `mapstructure:"start"`
	EditCategory   string `mapstructure:"editCategory"`
	UnknownCommand string `mapstructure:"unknownCommand"`
}

type Db struct {
	ConnectionUrl string
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

type MessageSender struct {
	MailingTimeEnd   int           `mapstructure:"mailingTimeEnd"`
	MailingTimeStart int           `mapstructure:"mailingTimeStart"`
	CategorySleep    time.Duration `mapstructure:"categorySleep"`
	RequestLimit     int8          `mapstructure:"requestLimit"`
}

type Import struct {
	BatchSize int           `mapstructure:"batchSize"`
	StartHour int           `mapstructure:"startHour"`
	DelayTime time.Duration `mapstructure:"delayTime"`
}

type Config struct {
	TelegramToken string
	News          News
	Rkn           Rkn
	Db            Db
	Messages      Messages
	MessageSender MessageSender
	Import        Import
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
	// Получаем абсолютный путь к исполняемому файлу (message_sender)
	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	// Определяем директорию, в которой лежит бинарник (обычно это .bin/)
	execDir := filepath.Dir(execPath)
	configPath := filepath.Join(execDir, "../configs") // Поднимаемся из .bin/ в корень, затем заходим в configs
	viper.AddConfigPath(configPath)
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

	if err := viper.UnmarshalKey("news", &cfg.News); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("rkn", &cfg.Rkn); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messageSender", &cfg.MessageSender); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("import", &cfg.Import); err != nil {
		return err
	}

	return nil
}

func fromEnv(cfg *Config) error {
	cfg.TelegramToken = os.Getenv("TELEGRAM_API_KEY")
	cfg.Db.ConnectionUrl = os.Getenv("DB_CONNECTION")
	cfg.News.ApiKey = os.Getenv("NEWS_API_KEY")

	return nil
}
