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
	ConnectionUrl     string
	TestConnectionUrl string
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
	testMode      bool
	TelegramToken string
	News          News
	Rkn           Rkn
	Db            Db
	Messages      Messages
	MessageSender MessageSender
	Import        Import
}

func Init(testMode bool) (*Config, error) {
	var cfg Config
	cfg.testMode = testMode

	if err := cfg.setUpViper(); err != nil {
		return nil, err
	}

	if err := cfg.unmarshal(); err != nil {
		return nil, err
	}

	if err := cfg.fromEnv(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) setUpViper() error {
	var cfgPath string
	if cfg.testMode {
		cfgPath = os.Getenv("MAIN_CONFIG_PATH_TEST")
	} else {
		cfgPath = os.Getenv("MAIN_CONFIG_PATH")
	}

	viper.AddConfigPath(cfgPath)
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
	cfg.Db.ConnectionUrl = os.Getenv("DB_CONNECTION")
	cfg.Db.TestConnectionUrl = os.Getenv("DB_CONNECTION_TEST")
	cfg.News.ApiKey = os.Getenv("NEWS_API_KEY")

	return nil
}
