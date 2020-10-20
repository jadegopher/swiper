package main

import (
	"log"
	"swiper/configurator"
)

func main() {
	c := configurator.New()
	config, err := c.Config()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(config)
}
