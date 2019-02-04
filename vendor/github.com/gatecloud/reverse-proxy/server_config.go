package proxy

import (
	"encoding/json"
	"io/ioutil"
)

// ServerConfig representing an array of Server
type ServerConfig []Server

// Default creates a ServerConfig handler
func Default(fileName string) (*ServerConfig, error) {
	sc := &ServerConfig{}
	if err := sc.loadFile(fileName); err != nil {
		return nil, err
	}
	return sc, nil
}

// LoadFile loads routing and forwarding files
func (sc *ServerConfig) loadFile(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, sc)
	if err != nil {
		return err
	}
	return nil
}
