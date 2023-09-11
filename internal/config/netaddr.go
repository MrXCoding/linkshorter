package config

import (
	"errors"
	"strconv"
	"strings"
)

type NetAddress struct {
	host string
	port int
}

func (a NetAddress) String() string {
	return a.host + ":" + strconv.Itoa(a.port)
}

func (a *NetAddress) Set(s string) error {
	hp := strings.Split(s, ":")
	if len(hp) != 2 {
		return errors.New("need address in a form host:port")
	}
	port, err := strconv.Atoi(hp[1])
	if err != nil {
		return err
	}
	a.host = hp[0]
	a.port = port
	return nil
}
