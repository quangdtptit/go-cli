package config

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type (
	Config struct {
		App      App      `mapstructure:"APP"`
		Postgres Postgres `mapstructure:"POSTGRES"`
	}

	App struct {
		Name    string `mapstructure:"NAME"`
		Mode    string `mapstructure:"MODE"`
		Version string `mapstructure:"VERSION"`
	}

	Postgres struct {
		PoolMax           int           `mapstructure:"POOL_MAX"`
		PoolMin           int           `mapstructure:"POOL_MIN"`
		MaxConnLifetime   time.Duration `mapstructure:"MAX_CONN_LIFETIME"`
		MaxConnIdleTime   time.Duration `mapstructure:"MAX_CONN_IDLE_TIME"`
		HealthCheckPeriod time.Duration `mapstructure:"HEALTH_CHECK_PERIOD"`
		Host              string        `mapstructure:"HOST"`
		Port              int           `mapstructure:"PORT"`
		Username          string        `mapstructure:"USERNAME"`
		Password          string        `mapstructure:"PASSWORD"`
		Database          string        `mapstructure:"DATABASE"`
	}
)

func NewConfig() (*Config, error) {
	return LoadConfig()
}

func LoadConfig() (*Config, error) {
	conf := &Config{}
	_ = godotenv.Load() // load from .env
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	for _, key := range viper.AllKeys() {
		str := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))

		err := viper.BindEnv(key, str)
		if err != nil {
			return nil, err
		}
	}

	err := conf.binding(viper.GetViper())
	if err != nil {
		return nil, err
	}
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		for _, key := range viper.AllKeys() {
			str := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
			err := viper.BindEnv(key, str)
			if err != nil {
				return
			}
		}

		if err := conf.binding(viper.GetViper()); err != nil {
			log.Println("binding error:", err)
		}
		log.Printf("config: %+v", conf)
	})
	return conf, nil
}

func (c *Config) binding(v *viper.Viper) error {
	if err := v.Unmarshal(&c); err != nil {
		log.Println("failed to unmarshal config: ", err)
		return err
	}
	return nil
}
