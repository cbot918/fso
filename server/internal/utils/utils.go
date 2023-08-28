package utils

import (
	"io"
	"os"
)

func ReadMap(path string) ([]byte, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	byteValue, err := io.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	return byteValue, nil
}
