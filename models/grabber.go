package models

type IFinder interface {
	FindKeys() ([]Auth, error)
}

type Auth struct {
	Key   string `json:"key"`
	Login login  `json:"login"`
}

type FirefoxLogin struct {
	NextId                           int64         `json:"nextId"`
	Logins                           []login       `json:"logins"`
	VulnerablePasswords              []interface{} `json:"potentiallyVulnerablePasswords"`
	DismissedBreachAlertsByLoginGUID interface{}   `json:"dismissedBreachAlertsByLoginGUID"`
	Version                          int64         `json:"version"`
}

type login struct {
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
