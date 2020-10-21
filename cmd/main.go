package main

import (
	"log"
	"swiper/configurator"
	"swiper/grabber"
	"swiper/saver"
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

	save := saver.New(config.StoreFilePath)
	if err := save.Save(data); err != nil {
		log.Fatal(err)
	}
}
