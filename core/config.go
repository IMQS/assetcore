package core

import "github.com/IMQS/serviceconfigsgo"

type Config struct {
	HttpPort int
}

func (c *Config) Load() error {
	err := serviceconfigsgo.GetConfig("AssetCore", serviceName, serviceConfigVersion, serviceConfigFileName, c)

}
