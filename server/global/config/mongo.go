package config

type MongoDb struct {	
	Username         string     `json:"username"`
    Password         string     `json:"password"`
    Hosts            []MongoHost `json:"hosts"`
    Database         string     `json:"database"`
    AuthSource       string     `json:"auth_source"`
    ReplicaSet       string     `json:"replica_set"`
    ReadPreference   string     `json:"read_preference"`
    MaxPoolSize      uint64     `json:"max_pool_size"`
    MinPoolSize      uint64     `json:"min_pool_size"`
    SocketTimeoutSec int        `json:"socket_timeout_sec"`
    ConnectTimeoutSec int      `json:"connect_timeout_sec"`
}

type MongoHost struct {
    IP   string `json:"ip"`
    Port int    `json:"port"`
}