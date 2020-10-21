package main

import (
	"fmt"
	"log"
	"swiper/configurator"
	"swiper/grabber"
)

func main() {
	c := configurator.New()

	config, err := c.Config()
	if err != nil {
		log.Fatal(err)
	}

	grab := grabber.New(config.MozillaPath, config.MasterPassword)

	data, err := grab.Passwords()
	if err != nil {
		log.Fatal(err)
	}

	if config.StoreFilePath == "" {
		for _, elem := range data {
			log.Println(fmt.Sprintf("Hostname: '%s'\n\tUsername: '%s'\n\tPassword: '%s'", elem.Hostname, elem.UsernameField, elem.PasswordField))
		}
	}
}
