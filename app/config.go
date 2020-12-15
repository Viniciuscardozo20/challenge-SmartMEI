package app

import (
	"encoding/json"
	"os"

	"challenge-SmartMEI/database"
)

type Config struct {
	LogLevel int      `json:"log_level"`
	Port     int      `json:"port"`
	Database Database `json:"database"`
}

type Database struct {
	Config         database.Config `json:"config"`
	UserCollection string          `json:"user_collection"`
}

func NewConfigFile(filename string) error {
	err := generateConfigFile(filename, configSample())
	if err != nil {
		return err
	}
	return nil
}

func generateConfigFile(filename string, config Config) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func configSample() Config {
	return Config{
		Port: 9000,
		Database: Database{
			Config: database.Config{
				Host:     "http://mongo.service.com.br",
				Port:     8529,
				User:     "root",
				Password: "123qwe",
				Database: "smartMei-db",
			},
			UserCollection: "user-collection",
		},
	}
}
