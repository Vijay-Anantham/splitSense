package config

type ServerConfig struct {
	Port string
}

func InitConfig() *ServerConfig {
	return &ServerConfig{
		Port: "8080",
	}
}
