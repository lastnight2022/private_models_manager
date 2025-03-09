package config

type Config struct {
	Database    Database    `json:"database"`
	RedisConfig RedisConfig `json:"redis"`
	Server      Server      `json:"server"`
}
