package saver

import (
	"swiper/models"
	"swiper/saver/file"
	"swiper/saver/stdout"
)

type saver struct {
	stdout models.ISaver
	file   models.ISaver
	path   string
}

func New(path string) *saver {
	return &saver{stdout: stdout.New(), file: file.New(path), path: path}
}

func (s *saver) Save(data []models.Login) error {
	if s.path == "" {
		if err := s.stdout.Save(data); err != nil {
			return err
		}
	} else {
		if err := s.file.Save(data); err != nil {
			return err
		}
	}
	return nil
}
