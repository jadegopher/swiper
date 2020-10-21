package models

type ISaver interface {
	Save([]Login) error
}
