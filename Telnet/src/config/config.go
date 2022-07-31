package config

type Config struct {
	*Gin
}

func Init() (*Config, error) {
	var conf Config
	conf.Gin = ginEngin()
	return &conf, nil
}
