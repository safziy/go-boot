package config

import (
	"path"

	"github.com/safziy/go-boot/web"

	"github.com/spf13/viper"
)

const (
	configPath = "configs/"
	configName = "application.yaml"
)

type BootConfig struct {
	Web *web.Config
}

func InitConfig(rootPath string) (*BootConfig, error) {
	return InitConfigWithFullPath(path.Join(rootPath, configPath, configName))
}

func InitConfigWithFullPath(fullPath string) (*BootConfig, error) {
	// Initialize Viper
	vp := viper.New()

	vp.SetConfigFile(fullPath)

	config := &BootConfig{}
	err := vp.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
