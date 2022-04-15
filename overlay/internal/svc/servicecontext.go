package svc

import (
	"github.com/hashicorp/consul/api"
	"github.com/r35krag0th/zero-gtfo-overlay/internal/rundowns"
	"github.com/r35krag0th/zero-gtfo-overlay/overlay/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config       config.Config
	ConsulClient *api.Client
	Rundowns     map[string]*rundowns.RundownDetails
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := api.NewClient(&api.Config{
		Address: c.Consul.Address,
		Scheme:  c.Consul.Scheme,
	})
	if err != nil {
		panic(err)
	}

	allRundowns, loadErr := rundowns.LoadFiles(c.RundownDataFiles)
	if len(loadErr) > 0 {
		for err := range loadErr {
			logx.Errorf("failed to load rundown details: %v", err)
		}
	}

	return &ServiceContext{
		Config:       c,
		ConsulClient: client,
		Rundowns:     allRundowns,
	}
}
