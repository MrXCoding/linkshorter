package config

import (
	"flag"
)

var (
	netAddr = NetAddress{
		host: "localhost",
		port: 8080,
	}
	baseURL *string
)

func Init() {
	flag.Var(&netAddr, "a", "Net address host:port")
	baseURL = flag.String("b", "http://localhost:8080", "base url")

	flag.Parse()
}

type Main struct {
	netAddr NetAddress
	baseURL string
}

func (m *Main) Host() string {
	return m.netAddr.String()
}

func (m *Main) GetBaseURL() string {
	return m.baseURL + "/"
}

func New() Main {
	return Main{
		netAddr: netAddr,
		baseURL: *baseURL,
	}
}
