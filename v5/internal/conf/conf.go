package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
)

type Config struct {
	Server   ServerConfig
	Data     DataConfig
	Logger   log.Config
	Service  ServiceConfig
}

type ServerConfig struct {
	HTTP  HTTPConfig
	GRPC  GRPCConfig
	Thrift ThriftConfig
}

type HTTPConfig struct {
	Addr string
}

type GRPCConfig struct {
	Addr string
}

type ThriftConfig struct {
	Addr string
}

type DataConfig struct {
	Database DatabaseConfig
	Cache    CacheConfig
}

type DatabaseConfig struct {
	Driver string
	Source string
}

type CacheConfig struct {
	Redis RedisConfig
}

type RedisConfig struct {
	Addr string
}

type ServiceConfig struct {
	Discovery DiscoveryConfig
}

type DiscoveryConfig struct {
	Name      string
	Endpoints []string
}

func NewConfig(path string) (*Config, error) {
	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
	var cfg Config
	if err := c.Scan(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
