package svc

import (
	"github.com/r35krag0th/zero-gtfo-overlay/internal/rundowns"
	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config   config.Config
	Rundowns map[string]*rundowns.RundownDetails
}

func NewServiceContext(c config.Config) *ServiceContext {

	allRundowns, loadErr := rundowns.LoadFiles(c.RundownDataFiles)
	if len(loadErr) > 0 {
		for err := range loadErr {
			logx.Errorf("failed to load rundown details: %v", err)
		}
	}

	return &ServiceContext{
		Config:   c,
		Rundowns: allRundowns,
	}
}
