package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Consul struct {
		Scheme  string
		Address string
	}
	CurrentRundown string
}
