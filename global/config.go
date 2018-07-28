package global

type Config struct {
	Sys SystemConfig
	Db DatabaseConfig
}

type SystemConfig struct {
	LogFile string
}

type DatabaseConfig struct {
	Username string
	Password string
	Host string
	Port string
}

func initConfig() *Config {
	sys := SystemConfig{
		LogFile: "/var/log/goblog.log",
	}
	return &Config {
		Sys: sys,
	}
}