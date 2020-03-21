package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pigeonligh/git-select/utils"
)

var (
	configDir  string
	configPath string
)

func checkPath(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return os.Mkdir(path, os.ModePerm)
	}
	return err
}

func newConfig() *Map {
	config := &Map{
		Data: make(map[string]Config, 0),
	}
	return config
}

// LoadConfig is
func LoadConfig() (*Map, error) {
	if err := checkPath(configDir); err != nil {
		return nil, err
	}
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		config := newConfig()
		if err := SaveConfig(config); err != nil {
			return nil, err
		}
		return config, nil
	} else if err != nil {
		return nil, err
	}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var config Map
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// SaveConfig is
func SaveConfig(config *Map) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(configPath, data, os.ModePerm)
}

func init() {
	homeDir, err := utils.Home()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}
	configDir = homeDir + "/.git-select"
	configPath = configDir + "/config.json"
}
