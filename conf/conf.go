package conf

import (
	"log"

	"github.com/spf13/viper"
)

type Options struct {
	Path string `json:"path"`
}

type option func(*Options)

func WithPath(path string) option {
	return func(o *Options) { o.Path = path }
}

func Setup(options ...option) error {
	ops := &Options{}

	for _, opt := range options {
		opt(ops)
	}

	viper.AddConfigPath(ops.Path)

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
