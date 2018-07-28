package global

type Config struct {
	Sys SystemConfig
	Db DatabaseConfig
}

type SystemConfig struct {
	Authurl string
	Logfile string
}

type DatabaseConfig struct {
	Username string
	Password string
	Host string
	Port string
}

func initConfig() *Config {
	sys := SystemConfig{
		Authurl: "/user/login",
		Logfile: "/var/log/goblog.log",
	}
	return &Config {
		Sys: sys,
	}
}