package config

import "time"

type (
	Config struct {
		AppName string
		Env     string
		HTTP    HTTP
		PubSub  PubSub
	}

	HTTP struct {
		Port             int
		GracefulShutdown time.Duration
		ReadTimeout      time.Duration
		WriteTimeout     time.Duration
		IdleTimeout      time.Duration
	}

	PubSub struct {
		ProjectA                    string
		ProjectB                    string
		EmployeeHiredSubscriptionID string
	}
)

func Read() *Config {
	return &Config{
		AppName: "go-structure-demo",
		Env:     "prod",
		HTTP: HTTP{
			Port:             8080,
			GracefulShutdown: time.Second,
			ReadTimeout:      3 * time.Second,
			WriteTimeout:     3 * time.Second,
			IdleTimeout:      3 * time.Second,
		},
		PubSub: PubSub{
			ProjectA:                    "",
			ProjectB:                    "",
			EmployeeHiredSubscriptionID: "the_id",
		},
	}
}
