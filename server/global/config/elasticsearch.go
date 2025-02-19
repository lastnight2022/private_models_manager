package config

type ElasticsearchConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	EnableAuth bool   `json:"enable_auth"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}
