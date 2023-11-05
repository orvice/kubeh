package config

import flag "github.com/spf13/pflag"

var configPath string

func Init() {
	flag.StringVar(&configPath, "configPath", "~/.kube/config", "Config Path")
}
