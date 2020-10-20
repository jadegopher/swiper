package environment

import (
	"os"
	"swiper/models"
)

type environment struct {
}

func New() models.IConfigurator {
	return &environment{}
}

func (e *environment) ReadConfig() (models.Config, error) {
	path, in := os.LookupEnv("FIREFOX")
	if !in {
		return models.Config{}, KeyNotFoundError
	}
	return models.Config{MozillaPath: path}, nil
}
