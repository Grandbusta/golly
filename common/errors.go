package common

import (
	"errors"
	"log"
	"strings"
)

func Error(message string, vars ...string) error {
	varLen := 0
	varPlace := 0
	errSlice := strings.Split(message, " ")

	for i := 0; i < len(errSlice); i++ {
		strSelection := errSlice[i]
		if string(strSelection[0]) == "{" {
			varPlace++
		}
	}

	if len(vars) < varPlace {
		log.Fatal("Insufficient variables")
	}

	for i := 0; i < len(errSlice); i++ {
		strSelection := errSlice[i]
		if string(strSelection[0]) == "{" {
			errSlice[i] = vars[varLen]
			varLen++
		}
	}
	return errors.New(strings.Join(errSlice, " "))
}
