package my_uuid

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

func DemoMyUUID() {
	for i := 0; i < 10; i++ {
		uuids, e := New()
		log.Println(fmt.Sprintf("%x", uuids), e)

	}
}

type UUID [16]byte

var rander = rand.Reader

// New generates a new uuid.
func New() (UUID, error) {
	var uuid [16]byte

	_, err := io.ReadFull(rander, uuid[:])
	if err != nil {
		return [16]byte{}, err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10

	return uuid, nil
}
