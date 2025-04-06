package config

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

type AppConfig struct {
    Color string `json:"color"`
}

func LoadConfig() (*AppConfig, error) {
    godotenv.Load(".env")

    file, err := os.Open("config.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    var config AppConfig
    if err := decoder.Decode(&config); err != nil {
        return nil, err
    }

    fmt.Println("Color desde configuración:", config.Color)
    return &config, nil
}