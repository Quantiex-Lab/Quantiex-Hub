package rpc

import (
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Port string `yaml:"port"`
	Url  string `yaml:"url"`
}

var c *Config

func init() {
	c = &Config{
		Port: ":1235",
		Url:  "http://127.0.0.1:1234",
	}
	c.getConf()
}

func (c *Config) getConf() *Config {
	yamlFile, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		glog.Errorf("yamlFile.Get err   #%v ", err)
		return c
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		glog.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func GetConfig() *Config {
	return c
}
