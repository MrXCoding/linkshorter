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

func Init() {
	flag.Var(&netAddr, "a", "Net address Host:Port")
	baseURL = flag.String("b", defBaseURL, "base url")

	flag.Parse()

	if envServAddr := os.Getenv("SERVER_ADDRESS"); envServAddr != "" {
		_ = netAddr.Set(envServAddr)
	}
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		baseURL = &envBaseURL
	}
}

type Main struct {
	NetAddr NetAddress
	BaseURL string
}

func (m *Main) Host() string {
	return m.NetAddr.String()
}

func (m *Main) GetBaseURL() string {
	return m.BaseURL + "/"
}

func New() Main {
	return Main{
		NetAddr: netAddr,
		BaseURL: *baseURL,
	}
}
