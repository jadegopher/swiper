package models

type IConfigurator interface {
	ReadConfig() (Config, error)
}

type Config struct {
	MozillaPath string `json:"mozilla_path"`
}
