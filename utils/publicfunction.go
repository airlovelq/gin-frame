package utils

import (
	"os"
)

func GetHomeDir() string {
	return os.Getenv("HOME")
}

func GetEnvDefault(key string, default_value string) string {
	if res := os.Getenv(key); res != "" {
		return res
	} else {
		return default_value
	}
}
