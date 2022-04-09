package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/consul/api"
)

const consulKeyFormat = "gtfo-helper/%s/overlay"

type OverlayData struct {
	Expedition struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"expedition"`
	Sector string `json:"sector"`
}

func GetOverlayData(c *api.Client) (*OverlayData, error) {
	consulKey := fmt.Sprintf(consulKeyFormat, os.Getenv("APP_ENV"))
	kv := c.KV()
	p, _, err := kv.Get(consulKey, nil)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, fmt.Errorf("key does not have a value: %s", consulKey)
	}

	var output OverlayData
	err = json.Unmarshal(p.Value, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func UpdateOverlayData(newData *OverlayData, c *api.Client) error {
	b, err := json.Marshal(newData)
	if err != nil {
		return err
	}
	consulKey := fmt.Sprintf(consulKeyFormat, os.Getenv("APP_ENV"))
	kv := c.KV()
	wp := &api.KVPair{Key: consulKey, Value: b}
	_, err = kv.Put(wp, nil)
	if err != nil {
		return err
	}

	return nil
}
