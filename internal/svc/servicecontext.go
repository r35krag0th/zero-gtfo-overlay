package svc

import (
	"strconv"

	"github.com/hashicorp/consul/api"
	"github.com/r35krag0th/zero-gtfo-overlay/internal/config"
	"github.com/r35krag0th/zero-gtfo-overlay/internal/data/rundowns"
	"github.com/zeromicro/go-zero/core/logx"
)

var rundownsToLoad = []int{5, 6}

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

	allRundowns := make(map[string]*rundowns.RundownDetails)

	for _, rundownDataFile := range c.RundownDataFiles {
		details, err := rundowns.LoadRundownDetailsFromYAML(rundownDataFile)
		if err != nil {
			logx.Errorf("failed to load rundown details %s: %v", rundownDataFile, err)
			continue
		}
		allRundowns[strconv.Itoa(details.ID)] = details
	}

	return &ServiceContext{
		Config:       c,
		ConsulClient: client,
		Rundowns:     allRundowns,
	}
}
