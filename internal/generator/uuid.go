package generator

import (
	"github.com/google/uuid"
)

type UUID struct {
	data []byte
}

func NewUUID(params map[string]interface{}) (Generator, error) {
	s := &UUID{
		data: []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"),
	}

	return s, nil
}

func (s *UUID) Generate() (string, error) {
	return uuid.NewSHA1(uuid.New(), s.data).String(), nil
}
