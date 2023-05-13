package config

import (
	"os"
	"strconv"
)

type AppConf struct {
	Environment string
	Name        string
}

type HttpConf struct {
	Port       string
	XRequestID string
	Timeout    int
}

type LogConf struct {
	Name string
}

type RedisConf struct {
	Host string
	Port string
}

type NatsConf struct {
	NatsHost   string
	NatsStatus string
}

// Config ...
type Config struct {
	App   AppConf
	Http  HttpConf
	Log   LogConf
	Redis RedisConf
	Nats  NatsConf
}

// NewConfig ...
func Make() Config {
	app := AppConf{
		Environment: os.Getenv("APP_ENV"),
		Name:        os.Getenv("APP_NAME"),
	}

	http := HttpConf{
		Port:       os.Getenv("HTTP_PORT"),
		XRequestID: os.Getenv("HTTP_REQUEST_ID"),
	}

	log := LogConf{
		Name: os.Getenv("LOG_NAME"),
	}

	redis := RedisConf{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}

	nats := NatsConf{
		NatsHost:   os.Getenv("NATS_HOST"),
		NatsStatus: os.Getenv("NATS_STATUS"),
	}

	// set default env to local
	if app.Environment == "" {
		app.Environment = "LOCAL"
	}

	// set default port for HTTP
	if http.Port == "" {
		http.Port = "8080"
	}

	httpTimeout, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT"))
	if err == nil {
		http.Timeout = httpTimeout
	}

	config := Config{
		App:   app,
		Http:  http,
		Log:   log,
		Redis: redis,
		Nats:  nats,
	}

	return config
}
