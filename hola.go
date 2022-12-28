package hola

import (
	"errors"
	"fmt"
)

func Hola(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hola %s", name)
	return message, nil
}
