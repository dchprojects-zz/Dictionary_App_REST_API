package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "50.87.139.17",
			Port:     3306,
			Username: "dchproje_N0",
			Password: "Future$Projects@2021",
			Name:     "dchproje_dbmdwords",
			Charset:  "utf8",
		},
	}
}
