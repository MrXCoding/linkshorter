package config

import (
	"flag"
)

var (
	netAddr = NetAddress{
		Host: "localhost",
		Port: 8080,
	}
	baseUrl = "http://localhost:8080/"
)

func Init() {
	flag.Var(&netAddr, "addr", "Net address host:port")
	baseUrl = *flag.String("baseUrl", baseUrl, "base url")

	flag.Parse()
}

type Main struct {
	netAddr NetAddress
	baseUrl string
}

func (m *Main) Host() string {
	return m.netAddr.String()
}

func (m *Main) GetBaseUrl() string {
	return m.baseUrl
}

func New() Main {
	return Main{
		netAddr: netAddr,
		baseUrl: baseUrl,
	}
}
