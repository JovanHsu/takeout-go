package main

import (
	"fmt"
	"takeout/nacos"
)

type Config struct {
	Root []*ConfigItem `mapstructure:"root"`
}

type ConfigItem struct {
	Hosts      string   `mapstructure:"hosts"`
	RemoteUser string   `mapstructure:"remote_user"`
	Roles      []string `mapstructure:"roles"`
}

func main() {
	config := nacos.LoadConfig()
	fmt.Println(config)
}
