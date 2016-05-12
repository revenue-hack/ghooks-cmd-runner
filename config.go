package main

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type config struct {
	Port      int    `toml:"port"`
	Host      string `toml:"host"`
	Secret    string `toml:"secret"`
	Logfile   string `toml:"logfile"`
	Pidfile   string `toml:"pidfile"`
	Daemonize bool   `toml:"daemonize"`
	Hook      []hook
}

type hook struct {
	Event string `toml:"event"`
	Cmd   string `toml:"command"`
}

type hooks struct {
	Hook []hook
}

func loadFile(filename string) (string, error) {
	var err error
	buf, err := ioutil.ReadFile(filename)

	return string(buf), err
}

func loadToml(filename string, c config) (config, error) {
	var config config
	buf, err := loadFile(filename)
	if err != nil {
		return config, err
	}

	_, err = toml.Decode(string(buf), &config)
	if err != nil {
		return config, err
	}

	if config.Port == 0 {
		config.Port = c.Port
	}

	if config.Host == "" {
		config.Host = c.Host
	}

	if !config.Daemonize {
		config.Daemonize = c.Daemonize
	}

	if config.Logfile == "" {
		config.Logfile = c.Logfile
	}

	if config.Pidfile == "" {
		config.Pidfile = c.Pidfile
	}

	return config, err
}
