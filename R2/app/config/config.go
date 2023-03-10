package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type env struct {
	Port string `mapstructure:"PORT"`
}

var ENV *env

func Load() {
	log.Println("loading app configs")

	ENV = &env{Port: ""}

	env := strings.ToLower(os.Getenv("GO_ENVIRONMENT"))
	viper.SetConfigFile(fmt.Sprintf("%s.env", env))
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var vipErr viper.ConfigFileNotFoundError
		if ok := errors.As(err, &vipErr); ok {
			log.Fatalln(fmt.Errorf("config file not found. %w", err))
		} else {
			log.Fatalln(fmt.Errorf("unexpected error loading config file. %w", err))
		}
	}

	if err := viper.Unmarshal(ENV); err != nil {
		log.Fatalln(fmt.Errorf("failed to unmarshal config. %w", err))
	}

	log.Println("app configs loaded")
}
