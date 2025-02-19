package config

type Config struct {
	Database            Database            `json:"database"`
	RedisConfig         RedisConfig         `json:"redis"`
	ElasticsearchConfig ElasticsearchConfig `json:"elasticsearch"`
	Server              Server              `json:"server"`
}
