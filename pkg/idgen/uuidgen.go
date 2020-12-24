package idgen

import "github.com/google/uuid"

type UUIDGenerator struct {
}

func (uuidGen UUIDGenerator) GeneratePassword() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func NewUUIDGenerator() (IDGenerator, error) {
	return UUIDGenerator{}, nil
}
