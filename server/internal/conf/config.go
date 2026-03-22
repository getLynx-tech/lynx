package conf

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"reflect"
)

type Config struct {
	DBHost            string `mapstructure:"DB_HOST" validate:"required"`
	DBPort            int    `mapstructure:"DB_PORT" validate:"required"`
	DBUser            string `mapstructure:"DB_USER" validate:"required"`
	DBPassword        string `mapstructure:"DB_PASSWORD" validate:"required"`
	DBName            string `mapstructure:"DB_NAME" validate:"required"`
	BasicAuthUser     string `mapstructure:"BASIC_AUTH_USER" validate:"required"`
	BasicAuthPassword string `mapstructure:"BASIC_AUTH_PASSWORD" validate:"required"`
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfgType := reflect.TypeOf(config)
	for i := 0; i < cfgType.NumField(); i++ {
		if env := cfgType.Field(i).Tag.Get("mapstructure"); env != "" {
			if err := viper.BindEnv(env); err != nil {
				return nil, err
			}
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return &config, err
	}

	if err := validator.New().Struct(config); err != nil {
		return &config, err
	}

	return &config, nil
}
