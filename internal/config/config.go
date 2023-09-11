package config

import (
	"flag"
)

var (
	netAddr = NetAddress{
		Host: "localhost",
		Port: 8080,
	}
	baseURL *string
)

func Init() {
	flag.Var(&netAddr, "a", "Net address Host:Port")
	baseURL = flag.String("b", "http://localhost:8080", "base url")

	flag.Parse()
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
