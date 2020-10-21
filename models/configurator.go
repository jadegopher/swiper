package models

type IConfigurator interface {
	ReadConfig() (Config, error)
}

type Config struct {
	MozillaPath    string `json:"mozilla_path"`
	MasterPassword []byte `json:"master_password"`
	StoreFilePath  string `json:"store_file_path"`
}
