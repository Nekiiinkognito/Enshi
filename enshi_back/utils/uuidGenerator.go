package utils

import (
	"encoding/binary"

	"github.com/google/uuid"
)

func GetUUIDv7AsInt64() (int64, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return -1, err
	}

	return -int64(
		binary.BigEndian.Uint64(uuid[8:]),
	), nil
}
