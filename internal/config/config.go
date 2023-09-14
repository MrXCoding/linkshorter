package config

import (
	"flag"
	"os"
)

const defBaseURL = "http://localhost:8080"

var (
	netAddr = NetAddress{
		Host: "localhost",
		Port: 8080,
	}
	baseURL *string
)

func Parse() {
	parseFlag()
	parseEnv()
}

func parseFlag() {
	flag.Var(&netAddr, "a", "Net address Host:Port")
	baseURL = flag.String("b", defBaseURL, "base url")

	flag.Parse()
}

func parseEnv() {
	if envServAddr := os.Getenv("SERVER_ADDRESS"); envServAddr != "" {
		_ = netAddr.Set(envServAddr)
	}
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		baseURL = &envBaseURL
	}
}

type Config struct {
	NetAddr NetAddress
	BaseURL string
}

func (m *Config) Host() string {
	return m.NetAddr.String()
}

func (m *Config) GetBaseURL() string {
	return m.BaseURL + "/"
}

func New() Config {
	return Config{
		NetAddr: netAddr,
		BaseURL: *baseURL,
	}
}
