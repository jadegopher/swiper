package commandline

import (
	"flag"
	"swiper/models"
)

type cmd struct {
	config models.Config
}

func New() models.IConfigurator {
	return &cmd{}
}

func (c *cmd) ReadConfig() (models.Config, error) {
	var conf string
	flag.StringVar(&conf, "flag", defaultPath, "path to Firefox")
	if conf == defaultPath {
		return models.Config{}, FlagNotSetErr
	}
	return models.Config{MozillaPath: conf}, nil
}
