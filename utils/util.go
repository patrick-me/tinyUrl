package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func GetEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)

	if len(value) == 0 {
		return defaultValue
	}

	v, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Error converting %s to int", value)
		return defaultValue
	}
	return v
}
