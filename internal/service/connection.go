package service

import (
	"changeme/internal/define"
	"encoding/json"
	"errors"
	"os"
)

// ConnectionList
func ConnectionList() ([]*define.Connection, error) {
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

// CreateConnection
func CreateConnection(conn *define.Connection) error {
	if conn.Address == "" {
		return errors.New("address cannot be empty")
	}

	// default
	if conn.Name == "" {
		conn.Name = conn.Address
	}
	if conn.Port == "" {
		conn.Port = "6379"
	}

	conf := &define.Config{}

	workDir, _ := os.Getwd()
	data, err := os.ReadFile(workDir + "/" + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		// create
		conf.Connections = []*define.Connection{conn}
	} else {
		// update
		json.Unmarshal(data, conf)
		conf.Connections = append(conf.Connections, conn)
	}
	data, _ = json.Marshal(conf)
	os.WriteFile(workDir+"/"+define.ConfigName, data, 0755)

	return nil
}
