package configs

const DefaultConfigFilePath = "/etc/loginhub.conf"

var DefaultConfig = &Config{
	Debug: true,
	Port:  8080,
}
