package utils

import (
	"strings"

	"github.com/gofrs/uuid"
)

func GenerateUUID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	res := strings.Replace(id.String(), "-", "", -1)
	return res, nil
}
