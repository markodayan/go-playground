package utils

import (
	"errors"
	"strings"
)

func Capitalize(name string) (string, int, error) {
	handle := func(err error) (string, int, error) {
		return "", 0, err
	}

	if name == "" {
		return handle(errors.New("no name provided"))
	}


	return strings.ToTitle(name), len(name), nil
}