package config

import "os"

type Config struct {
    Port        string
    UserSvcAddr string
}

func LoadConfig() *Config {
    return &Config{
        Port:        getEnv("PORT", ":8080"),
        UserSvcAddr: getEnv("USER_SERVICE_ADDR", "localhost:50051"),
    }
}

func getEnv(key, defaultVal string) string {
    if val := os.Getenv(key); val != "" {
        return val
    }
    return defaultVal
}
