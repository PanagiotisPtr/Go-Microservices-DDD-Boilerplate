package config

type Configuration struct {
	Service  ServiceConfig
	Database DatabaseConfig
}

type ServiceConfig struct {
	Port               string
	PublicKeyLocation  string
	PrivateKeyLocation string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}
