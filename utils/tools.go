package utils

import (
	"github.com/google/uuid"
	uid "k8s.io/apimachinery/pkg/util/uuid"
)

func GenerateUUID() string {
	uuid := uuid.New()
	return uuid.String()
}

func GenerateShortUUID() string {
	id := uid.NewUUID()
	shortID := string(id)[len(id)-8:]
	return shortID
}
