package config

type Server struct {
	Port    int    `json:"port"`
	Secret  string `json:"secret"`
	Timeout int    `json:"timeout"`
}
