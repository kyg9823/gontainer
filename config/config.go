package config

import "os"

type ConfigList struct {
	ContainerdAddress string
}

var Config *ConfigList

func LoadConfig() {
	Config = &ConfigList{
		ContainerdAddress: getEnv("CONTAINERD_ADDRESS", "/run/containerd/containerd.sock"),
	}
}

func GetContainerdAddress() string {
	return Config.ContainerdAddress
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
