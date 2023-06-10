package util

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	Profile             string        `mapstructure:"PROFILE"`
}

func NewConfig() *Config {
	var defaultEnv string
	if os.Getenv("ENV") == "" {
		defaultEnv = "local"
	} else {
		defaultEnv = os.Getenv("ENV")
	}

	return &Config{
		Profile: defaultEnv,
	}
}

func (c *Config) LoadConfig(env string) (err error) {
	path, _ := GetRootPath()

	viper.AddConfigPath(path)
	viper.SetConfigName("app." + env)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}

func GetRootPath() (ex string, err error) {
	ex, _ = os.Getwd()
	_, err = os.Stat(filepath.Join(ex, "go.mod"))

	if err != nil {
		for i := 0; i < 5; i++ {
			ex = filepath.Join(ex, "../")
			_, err = os.Stat(filepath.Join(ex, "go.mod"))

			if err == nil {
				break
			}
		}
		if err != nil {
			log.Println("No env file provided, using only env variables")
		}
	}
	return
}

func GetMigrationsFolder() (ex string, err error) {
	ex, err = GetRootPath()
	if err != nil {
		return
	}

	ex = filepath.Join(ex, "src/infrastructure/db/migrations/")

	return
}
