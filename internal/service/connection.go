package service

import (
	"changeme/internal/define"
	"encoding/json"
	"errors"
	"os"
)

// ConnectionList
func ConnectionList() ([]*define.Connection, error) {
	// mock
	return []*define.Connection{{
		Identity: "local",
		Name:     "Local",
		Address:  "127.0.0.1",
		Port:     "6379",
		User:     "",
		Password: "",
	}}, nil

	workDir, _ := os.Getwd()
	data, err := os.ReadFile(workDir + "/" + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	conf := &define.Config{}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}

	return conf.Connections, nil
}
