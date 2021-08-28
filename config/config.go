package config

import (
	"github.com/aURORA-JC/headerfaker/model"
	"gopkg.in/ini.v1"
)

var Config = new(model.AppConfig)

// InitConfig Read config from config.ini
func InitConfig(filename string) error {
	return ini.MapTo(Config, filename)
}
