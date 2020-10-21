package stdout

import (
	"fmt"
	"swiper/models"
)

type stdout struct {
}

func New() models.ISaver {
	return &stdout{}
}

func (s *stdout) Save(data []models.Login) error {
	for _, elem := range data {
		fmt.Println(fmt.Sprintf("Hostname: '%s'\n\tUsername: '%s'\n\tPassword: '%s'",
			elem.Hostname, elem.UsernameField, elem.PasswordField))
	}
	return nil
}
