package config

import "time"

type Configuration struct {
	HeartbeatExpirePeriod time.Duration
}

func DefaultConfiguration() *Configuration {
	return &Configuration{}
}
