package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Grandbusta/golly/common"
)

type StringRules struct {
	Min       int
	Max       int
	Length    int
	Alphanum  bool
	Email     bool
	Lowercase bool
	Uppercase bool
	Uri       bool
	Required  bool
}

var stringMessages = map[string]string{
	"empty":     "{label} is not allowed to be empty",
	"email":     "{label} must be a valid email",
	"uppercase": "{label} must only contain uppercase characters",
	"lowercase": "{label} must only contain lowercase characters",
	"max":       "{label} must be less than or equal to {limit} characters long",
	"min":       "{label} length must be at least {limit} characters long",
	"alphanum":  "{label} must only contain alpha-numeric characters",
	"minmax":    "max must me greater than min",
	"length":    "{label} length must be {limit} characters long",
}

func StringValidate(str string, r *StringRules, label string) error {
	// check if required is present
	if r.Required {
		if len(str) <= 0 {
			return common.Error(stringMessages["empty"], fmt.Sprintf("\"%v\"", label))
		}
	}

	// min must be greater than max
	if r.Max != 0 && r.Max < r.Min {
		return common.Error(stringMessages["minmax"])
	}

	// only execute if min is not 0
	if r.Min != 0 {
		if len(str) < r.Min {
			return common.Error(stringMessages["min"], fmt.Sprintf("\"%v\"", label), strconv.Itoa(r.Min))
		}
	}

	// only execute if max is not 0
	if r.Max != 0 {
		if len(str) > r.Max {
			return common.Error(stringMessages["max"], fmt.Sprintf("\"%v\"", label), strconv.Itoa(r.Max))
		}
	}

	if r.Length != 0 && len(str) != r.Length {
		return common.Error(stringMessages["length"], fmt.Sprintf("\"%v\"", label), strconv.Itoa(r.Length))
	}

	if r.Alphanum {
		regex, _ := regexp.Compile(`^[a-zA-Z0-9]+$`)
		if !regex.MatchString(str) {
			return common.Error(stringMessages["alphanum"], fmt.Sprintf("\"%v\"", label))
		}
	}
	if r.Lowercase {
		if !(strings.ToLower(str) == str) {
			return common.Error(stringMessages["lowercase"], fmt.Sprintf("\"%v\"", label))
		}
	}

	if r.Uppercase {
		if !(strings.ToUpper(str) == str) {
			return common.Error(stringMessages["uppercase"], fmt.Sprintf("\"%v\"", label))
		}
	}
	return nil
}
