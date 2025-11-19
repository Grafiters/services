package lib

import (
    "fmt"
    "os"
)

func GetVarEnv(key string) (string, error) {
    value := os.Getenv(key)
    if value == "" {
        return "", fmt.Errorf("environment variable %s is not set", key)
    }
    return value, nil
}