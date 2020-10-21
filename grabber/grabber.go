package grabber

import (
	"swiper/grabber/finder"
	"swiper/models"
)

type grabber struct {
	finder models.IFinder
	path   string
}

func New(path string, masterPassword []byte) *grabber {
	return &grabber{
		finder: finder.New(path, masterPassword),
		path:   path,
	}
}

func (g *grabber) Passwords() ([]models.Login, error) {
	data, err := g.finder.FindKeys()
	if err != nil {
		return nil, err
	}
	return data, nil
}
