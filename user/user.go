package user

import (
	"strconv"

	"github.com/sgdfgo/scraper"
)

type User struct {
	s           *scraper.Scraper
	ID          int
	StructureID int
}

func Connect(login, password string) (*User, error) {
	ID, err := strconv.Atoi(login)
	if err != nil {
		return nil, err
	}

	s := scraper.New()
	structureID, err := s.Connect(login, password)
	if err != nil {
		return nil, err
	}

	return &User{
		s:           s,
		ID:          ID,
		StructureID: structureID,
	}, nil
}
