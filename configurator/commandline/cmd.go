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
	var password string
	var save string
	flag.StringVar(&conf, "path", defaultPath, "path to Firefox")
	flag.StringVar(&password, "pwd", defaultPassword, "master password")
	flag.StringVar(&save, "o", "", "path to file where data will store")
	flag.Parse()
	if conf == defaultPath {
		return models.Config{}, FlagNotSetErr
	}
	return models.Config{MozillaPath: conf, MasterPassword: []byte(password), StoreFilePath: save}, nil
}
