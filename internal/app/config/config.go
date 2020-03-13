package config

type Config struct {
	DB   Database `yaml:"database"`
	GRPC GRPC     `yaml:"grpc"`
}
