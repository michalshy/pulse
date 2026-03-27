package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type PulseProject struct {
	Key  string
	Name string
}

type PulseAgent struct {
	Host                  string
	Port                  int64
	Binary                string
	AutoStart             bool
	BufferSize            int
	FlushIntervalSecs     int
	FlushThreshold        float64
	HeartbeatIntervalSecs int
}

type PulseIngest struct {
	Endpoint      string
	TimeoutSecs   int
	RetryAttempts int
}

type BackendConfig struct {
	Port int
}

type AppConfig struct {
	Project PulseProject
	Agent   PulseAgent
	Ingest  PulseIngest
	Backend BackendConfig
}

func ParseConfig(path string) (AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return AppConfig{}, err
	}

	var conf AppConfig
	_, err = toml.Decode(string(data), &conf)
	if err != nil {
		return AppConfig{}, err
	}
	return conf, nil
}
