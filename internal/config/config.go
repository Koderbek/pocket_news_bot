package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

type Messages struct {
	RateLimit           int           `mapstructure:"rateLimit"`
	RateLimitInterval   time.Duration `mapstructure:"rateLimitInterval"`
	StartCommand        string        `mapstructure:"startCommand"`
	EditCategoryCommand string        `mapstructure:"editCategoryCommand"`
	UnknownCommand      string        `mapstructure:"unknownCommand"`
	ManyRequestsCommand string        `mapstructure:"manyRequestsCommand"`
}

type Db struct {
	ConnectionUrl string
	RootPath      string
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

type Cleaner struct {
	Period time.Duration `mapstructure:"period"`
}

type Stmp struct {
	Password string
	Server   string `mapstructure:"server"`
	Port     int    `mapstructure:"port"`
	From     string `mapstructure:"from"`
	To       string `mapstructure:"to"`
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
	RootPath      string
	News          News
	Rkn           Rkn
	Cleaner       Cleaner
	Stmp          Stmp
	Db            Db
	Messages      Messages
	MessageSender MessageSender
	Import        Import
}

func Init() (*Config, error) {
	var cfg Config
	if err := cfg.fromEnv(); err != nil {
		return nil, err
	}

	if err := cfg.setUpViper(); err != nil {
		return nil, err
	}

	if err := cfg.unmarshal(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) setUpViper() error {
	viper.AddConfigPath(cfg.RootPath + "/configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}

func (cfg *Config) unmarshal() error {
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

	if err := viper.UnmarshalKey("cleaner", &cfg.Cleaner); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("stmp", &cfg.Stmp); err != nil {
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

func (cfg *Config) fromEnv() error {
	cfg.TelegramToken = os.Getenv("TELEGRAM_API_KEY")
	cfg.RootPath = os.Getenv("ROOT_PATH")
	cfg.Db.ConnectionUrl = os.Getenv("DB_CONNECTION")
	cfg.Db.RootPath = os.Getenv("DB_ROOT_PATH")
	cfg.News.ApiKey = os.Getenv("NEWS_API_KEY")
	cfg.Stmp.Password = os.Getenv("STMP_PASSWORD")

	return nil
}
