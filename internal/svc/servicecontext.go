package svc

import (
	"github.com/hashicorp/consul/api"
	"github.com/r35krag0th/zero-gtfo-overlay/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	ConsulClient *api.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := api.NewClient(&api.Config{
		Address: c.Consul.Address,
		Scheme:  c.Consul.Scheme,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:       c,
		ConsulClient: client,
	}
}
