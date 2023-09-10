package common

import (
	"errors"
	"fmt"
	"strings"
)

func Error(message string, vars ...[]string) error {
	errSlice := strings.Split(message, " ")
	for i := 0; i < len(errSlice); i++ {
		strSelection := errSlice[i]
		if string(strSelection[0]) == "{" {

		}
	}
	fmt.Println(errSlice)
	return errors.New(message)
}
