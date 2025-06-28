package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
)

// TODO: we can use here a validaotr liberary to validate the required fields
type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() (*Env, error) {
	env := Env{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("can't find .env: %w", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		return nil, fmt.Errorf("can't load env: %w", err)
	}

	//TODO: replace with a validator on the struct
	// Validate required fields
	if env.DBHost == "" {
		return nil, fmt.Errorf("missing required env var: DB_HOST")
	}
	if env.DBName == "" {
		return nil, fmt.Errorf("missing required env var: DB_NAME")
	}
	if env.DBPass == "" {
		return nil, fmt.Errorf("missing required env var: DB_PASS")
	}
	if env.AccessTokenSecret == "" {
		return nil, fmt.Errorf("missing required env var: ACCESS_TOKEN_SECRET")
	}

	return &env, nil
}
