package config

import "os"

type Config struct{}

func (c Config) GetConnectionString() string {
	host := os.Getenv("OB_POSTGRES_HOST")
	user := os.Getenv("OB_POSTGRES_USER")
	name := c.GetDB()
	password := os.Getenv("OB_POSTGRES_PASSWORD")
	sslMode := os.Getenv("OB_POSTGRES_SSL_MODE")

	cs := "host=" + host + " user=" + user + " dbname=" + name + " password=" + password + " sslmode=" + sslMode

	return cs
}

func (Config) GetDriver() string {
	return os.Getenv("OB_DB_DRIVER")
}

func (Config) GetDB() string {
	return os.Getenv("OB_DB_NAME")
}
