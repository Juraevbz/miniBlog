package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
)

type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
}

func InitConfig(logger zerolog.Logger) (dsn string, err error) {
	var config struct {
		DB DBConfig `json:"db"`
	}

	file, err := os.Open("config/config.json")
	if err != nil {
		logger.Error().Err(err).Msg("error pening 'config.json' file")
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		logger.Error().Err(err).Msg("error reading file contents")
		return "", err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		logger.Error().Err(err).Msg("error unmarshaling data to config")
		return "", err
	}

	dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DBName)

	return dsn, nil
}
