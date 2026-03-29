package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type PulseProject struct {
	Key  string `toml:"key"`
	Name string `toml:"name"`
}

type PulseAgent struct {
	Host                  string  `toml:"host"`
	Port                  int64   `toml:"port"`
	Binary                string  `toml:"binary"`
	AutoStart             bool    `toml:"auto_start"`
	BufferSize            int     `toml:"buffer_size"`
	FlushIntervalSecs     int     `toml:"flush_interval_secs"`
	FlushThreshold        float64 `toml:"flush_threshold"`
	HeartbeatIntervalSecs int     `toml:"heartbeat_interval_secs"`
}

type PulseIngest struct {
	Endpoint      string `toml:"endpoint"`
	TimeoutSecs   int    `toml:"timeout_secs"`
	RetryAttempts int    `toml:"retry_attempts"`
}

type BackendConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
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
