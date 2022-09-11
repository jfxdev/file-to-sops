package config

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

var cfg Config

type Config struct {
	Files Files `yaml:"files"`
}

type Files struct {
	Extensions []string `yaml:"extensions"`
}

func init() {
	body, err := ioutil.ReadFile(".config.yaml")
	if err != nil {
		fmt.Errorf("unable to read configuration file: %v", err.Error())
	}

	err = yaml.Unmarshal(body, &cfg)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Get() Config {
	return cfg
}

func (f Files) Valid(input string) (err error) {
	for _, v := range f.Extensions {
		if strings.HasSuffix(input, v) {
			return nil
		}
	}
	return fmt.Errorf("this file does not match any extensions in config file: '%s'", input)
}
