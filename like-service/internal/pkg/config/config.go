package config

import "github.com/spf13/viper"

type Postgres struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type Config struct {
	Postgres Postgres

	LikeServiceHost string
	LikeServicePort int
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
		LikeServiceHost: viper.GetString("service.host"),
		LikeServicePort: viper.GetInt("service.port"),
	}

	return &cfg, nil
}
