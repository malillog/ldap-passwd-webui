package app

import (
	"os"
	"strconv"
)

func getTitle() string {
	return envStr("LPW_TITLE", "Change your password on example.org")
}

func envStr(key, defaultValue string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return defaultValue
}

func envInt(key string, defaultValue int) int {
	val := os.Getenv(key)
	if val != "" {
		i, err := strconv.Atoi(val)
		if err != nil {
			return defaultValue
		}
		return i
	}
	return defaultValue
}

func envBool(key string, defaultValue bool) bool {
	val := os.Getenv(key)
	if val != "" {
		b, err := strconv.ParseBool(val)
		if err != nil {
			return defaultValue
		}
		return b
	}
	return defaultValue
}
