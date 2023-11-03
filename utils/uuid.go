package utils

import "github.com/google/uuid"

func IsValidUUID(str string) (bool, error) {
	_, err := uuid.Parse(str)
	return err == nil, err
}
