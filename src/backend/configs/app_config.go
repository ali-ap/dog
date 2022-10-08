package configs

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Application struct {
		// Port is the TCP Port which server to bind the HTTP Server to
		Port string `yaml:"port"`
	} `yaml:"application"`
	Dig struct {
		// Host is the DNS address
		Host string `yaml:"host"`
	} `yaml:"dig"`
}

var AppConfig *Config

func NewConfig() (*Config, error) {

	configPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func ParseFlags() (string, error) {
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./configs/app_config.yml", "path to config file")
	flag.Parse()

	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}
