package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/config/apollo"
	"github.com/go-kratos/kratos/v2/log"
)

type Config struct {
	Server   ServerConfig
	Data     DataConfig
	Logger   LoggerConfig
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

type LoggerConfig struct {
	Level  string
	Output string
}

type ServiceConfig struct {
	Discovery DiscoveryConfig
}

type DiscoveryConfig struct {
	Type      string
	Endpoints []string
}

func NewConfig(path string) (*Config, error) {
	var c Config
	config, err := config.New(
		config.WithSource(
			file.NewSource(path),
			apollo.NewSource(),
		),
		config.WithDecoder(func(kv *config.KeyValue) (interface{}, error) {
			return kv.Value, nil
		}),
	)
	if err != nil {
		return nil, err
	}
	if err := config.Scan(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func NewLogger(c *Config) (log.Logger, error) {
	return log.NewStdLogger(c.Logger.Output), nil
}
