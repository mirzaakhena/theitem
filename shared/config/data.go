package config

type Config struct {
	Servers  map[string]Server `json:"servers"`
	Database Database          `json:"database"`
}

type Server struct {
	Address string `json:"address,omitempty"`
}

type Database struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	DBName   string `json:"db_name"`
}
