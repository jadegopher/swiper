package configurator

import (
	"fmt"
	"log"
	"swiper/configurator/commandline"
	"swiper/configurator/environment"
	"swiper/models"
)

type configurator struct {
	cmd    models.IConfigurator
	env    models.IConfigurator
	config models.Config
}

func New() *configurator {
	return &configurator{
		cmd: commandline.New(),
		env: environment.New(),
	}
}

func (c *configurator) Config() (models.Config, error) {
	var err error
	if c.config, err = c.cmd.ReadConfig(); err != nil {
		log.Println(fmt.Sprintf("commandline: %s", err.Error()))
	} else {
		return c.config, nil
	}
	if c.config, err = c.env.ReadConfig(); err != nil {
		log.Println(fmt.Sprintf("enviroment: %s", err.Error()))
	} else {
		return c.config, nil
	}
	return c.config, err
}
