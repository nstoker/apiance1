package models

import (
	"log"

	"github.com/segmentio/ksuid"
)

// GenKsuid returns a suid
func GenKsuid() string {
	ID := ksuid.New().String()
	log.Printf("id: %s", ID)
	return ID
}
