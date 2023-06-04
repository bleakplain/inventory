package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
)

type Config struct {
	HTTP    HTTPConfig    `yaml:"http"`
	GRPC    GRPCConfig    `yaml:"grpc"`
	Thrift  ThriftConfig  `yaml:"thrift"`
	Data    DataConfig    `yaml:"data"`
	Logger  LoggerConfig  `yaml:"logger"`
	Service ServiceConfig `yaml:"service"`
}

type HTTPConfig struct {
	Addr string `yaml:"addr"`
}

type GRPCConfig struct {
	Addr string `yaml:"addr"`
}

type ThriftConfig struct {
	Addr string `yaml:"addr"`
}

type DataConfig struct {
	Database DatabaseConfig `yaml:"database"`
	Cache    CacheConfig    `yaml:"cache"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

type CacheConfig struct {
	Network string `yaml:"network"`
	Addr    string `yaml:"addr"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
}

type ServiceConfig struct {
	Discovery DiscoveryConfig `yaml:"discovery"`
}

type DiscoveryConfig struct {
	Name      string   `yaml:"name"`
	Endpoints []string `yaml:"endpoints"`
}

func LoadConfig(path string) (*Config, error) {
	var c Config
	if err := config.New(
		config.WithSource(file.NewSource(path)),
		config.WithDecoder(func(kv *config.KeyValue, v interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	).Scan(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func NewLogger(c *LoggerConfig) log.Logger {
	var opts []log.Option
	if c.Level != "" {
		opts = append(opts, log.WithLevel(log.Level(c.Level)))
	}
	if c.Output != "" {
		opts = append(opts, log.WithOutput(log.Output(c.Output)))
	}
	return log.NewHelper(log.NewStdLogger(), opts...)
}
