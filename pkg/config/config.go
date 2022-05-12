package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Database Postgres
		Port     string
	}

	Postgres struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
)

func New() *Config {
	v := viper.New()
	v.SetConfigName("config")         // name of config file (without extension)
	v.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath("/etc/chi-test/") // path to look for the config file in
	v.AddConfigPath("$HOME/chi-test") // call multiple times to add many search paths
	v.AddConfigPath(".")              // optionally look for config in the working directory
	err := v.ReadInConfig()           // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	return read(v)
}

func read(v *viper.Viper) *Config {
	pgHost := v.GetString("postgres_host")
	if pgHost == "" {
		panic(errors.New("postgres_host is required"))
	}

	pgPort := v.GetInt("postgres_port")
	if pgPort == 0 {
		panic(errors.New("postgres_port is required"))
	}

	pgUser := v.GetString("postgres_user")
	if pgUser == "" {
		panic(errors.New("postgres_user is required"))
	}

	pgPassword := v.GetString("postgres_password")
	if pgPassword == "" {
		panic(errors.New("postgres_password is required"))
	}

	pgDBName := v.GetString("postgres_dbname")
	if pgDBName == "" {
		panic(errors.New("postgres_dbname is required"))
	}

	port := v.GetString("port")
	if port == "" {
		panic(errors.New("port is required"))
	}

	return &Config{
		Database: Postgres{
			Host:     pgHost,
			Port:     pgPort,
			User:     pgUser,
			Password: pgPassword,
			DBName:   pgDBName,
		},
		Port: port,
	}
}
