package conf

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Options struct {
	Path []string `json:"path"`
}

type option func(*Options)

func WithPath(path ...string) option {
	return func(o *Options) { o.Path = path }
}

func Setup(options ...option) error {
	ops := &Options{}

	for _, opt := range options {
		opt(ops)
	}
	for _, path := range ops.Path {
		viper.AddConfigPath(path)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.String())
	})
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
	return nil
}

var (
	GetBool   = viper.GetBool
	GetString = viper.GetString
	GetInt    = viper.GetInt
	GetInt32  = viper.GetInt32
	Get       = viper.Get
)
