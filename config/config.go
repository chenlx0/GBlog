package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Blog struct {
	ArticleDir string `yaml:"article_dir"`
	Title      string `yaml:"title"`
	Author     string `yaml:"author"`
}

type Config struct {
	Server Server `yaml:"Server"`
	Blog   Blog   `yaml:"Blog"`
}

func FromYamlFile(path string) (*Config, error) {
	c := &Config{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}

	return c, nil
}
