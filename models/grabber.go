package models

import "database/sql"

type IFinder interface {
	FindKeys() ([]Login, error)
}

type IDecrypter interface {
	Decrypt(db *sql.DB, login Login, masterPassword []byte) (Login, error)
}

type FirefoxLogin struct {
	NextId                           int64         `json:"nextId"`
	Logins                           []Login       `json:"logins"`
	VulnerablePasswords              []interface{} `json:"potentiallyVulnerablePasswords"`
	DismissedBreachAlertsByLoginGUID interface{}   `json:"dismissedBreachAlertsByLoginGUID"`
	Version                          int64         `json:"version"`
}

type Login struct {
	Id                  int64       `json:"id"`
	Hostname            string      `json:"hostname"`
	HttpRealm           interface{} `json:"httpRealm"`
	FromSubmitUrl       string      `json:"fromSubmitUrl"`
	UsernameField       string      `json:"usernameField"`
	PasswordField       string      `json:"passwordField"`
	EncryptedUsername   string      `json:"encryptedUsername"`
	EncryptedPassword   string      `json:"encryptedPassword"`
	Guid                string      `json:"guid"`
	EncType             int8        `json:"EncType"`
	TimeCreated         int64       `json:"timeCreated"`
	TimeLastUsed        int64       `json:"timeLastUsed"`
	TimePasswordChanged int64       `json:"timePasswordChanged"`
	TimesUsed           int64       `json:"timesUsed"`
}
