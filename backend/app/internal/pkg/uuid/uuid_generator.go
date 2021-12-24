package uuid

import (
	"fmt"

	"github.com/google/uuid"
)

type idGenerator struct{}

func NewIDGenerator() *idGenerator {
	return &idGenerator{}
}

func (generator *idGenerator) GenerateID() (ID, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("there is something wrong with id generator : %w", err)
	}

	return ID(uuid.String()), nil
}
