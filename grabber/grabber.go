package grabber

import (
	"log"
	"swiper/grabber/finder"
	"swiper/models"
)

type grabber struct {
	finder models.IFinder
	path   string
}

func New(path string) *grabber {
	return &grabber{
		finder: finder.New(path),
		path:   path,
	}
}

func (g *grabber) Passwords() ([]models.Auth, error) {
	data, err := g.finder.FindKeys()
	if err != nil {
		return nil, err
	}
	log.Println(data)
	return nil, nil
}
