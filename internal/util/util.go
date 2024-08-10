package util

import "github.com/google/uuid"

func GenerateUUID() string {
	uid, err := uuid.NewV7()
	if err != nil {
		return uuid.NewString()
	}
	return uid.String()
}
