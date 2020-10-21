package file

import (
	"encoding/json"
	"os"
	"swiper/models"
)

type file struct {
	path string
}

func New(path string) models.ISaver {
	return &file{path: path}
}

func (f *file) Save(data []models.Login) error {
	file, err := os.Create(f.path)
	if err != nil {
		return err
	}
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if _, err := file.Write(byteData); err != nil {
		return err
	}
	return nil
}
