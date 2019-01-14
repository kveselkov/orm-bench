package config

type Config struct {}

func (Config) GetConnectionString() string {
	return "host=localhost user=docker dbname=test password=dockerpass sslmode=disable"
}

func (Config) GetDriver() string {
	return "postgres"
}

func (Config) GetDB() string {
	return "test"
}