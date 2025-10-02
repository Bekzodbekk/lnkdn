package config

import "github.com/spf13/viper"

type Postgres struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type Redis struct {
	Host string
	Port int
}

type Config struct {
	Postgres Postgres
	Redis Redis

	AuthServiceHost string
	AuthServicePort int

	JWTSecretKey string
}

func LOAD(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		Postgres: Postgres{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetInt("postgres.port"),
			Database: viper.GetString("postgres.dbname"),
			Username: viper.GetString("postgres.username"),
			Password: viper.GetString("postgres.password"),
		},
		Redis: Redis{
			Host: viper.GetString("redis.host"),
			Port: viper.GetInt("redis.port"),
		},
		AuthServiceHost: viper.GetString("service.host"),
		AuthServicePort: viper.GetInt("service.port"),
		
		JWTSecretKey: viper.GetString("secret-key"),
	}

	return &cfg, nil
}
