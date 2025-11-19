package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jinzhu/configor"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ConfigFile string
	Server     struct {
		Port         string
		SwagerEnable bool
		Timeout      uint32
		Appurl       string
		Appname      string
		LogLevel     uint32 `default:"4"`
	}
	Store struct {
		Host     string `required:"true"`
		Port     int    `required:"true"`
		User     string `required:"true"`
		Password string `required:"true"`
		Dbname   string `required:"true"`
	}
}

func NewConfig(configFile string) (*Config, error) {
	config := &Config{ConfigFile: configFile}
	if err := configor.Load(config, configFile); err != nil {
		return nil, err
	}
	return config, nil
}

func ReadConfig(cfg interface{}, fullPathURL string) error {

	getFormatFile := filePath(fullPathURL)

	switch getFormatFile {
	case ".json":
		fname := fullPathURL
		jsonFile, err := ioutil.ReadFile(fname)
		if err != nil {
			return err
		}
		return json.Unmarshal(jsonFile, cfg)
	default:
		fname := fullPathURL
		yamlFile, err := ioutil.ReadFile(fname)
		if err != nil {
			return err
		}
		return yaml.Unmarshal(yamlFile, cfg)
	}

}

func filePath(root string) string {
	var file string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		file = filepath.Ext(info.Name())
		return nil
	})
	return file
}
