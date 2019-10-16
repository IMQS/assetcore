package core

import (
	"github.com/IMQS/nf/nfdb"
	serviceconfig "github.com/IMQS/serviceconfigsgo"
)

type Config struct {
	HttpPort int
	DB       nfdb.DBConfig
}

func (c *Config) Load() error {
	return serviceconfig.GetConfig("", "assetcore", 0, "assetcore.json", c)
}
