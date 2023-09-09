package golly

import (
	"errors"
	"fmt"
	"reflect"
)

func stringValidation(str string, r *Rules) error {
	fmt.Println(r)
	if r.Required {
		if len(str) == 0 {
			return errors.New("not allowed to be empty")
		}
	}
	if r.Max != 0 && r.Max < r.Min {
		return errors.New("max should be greater than min")
	}
	if r.Min != 0 {
		if len(str) < r.Min {
			return errors.New("length of value is below minimum")
		}
	}
	if r.Max != 0 {
		if len(str) > r.Max {
			return errors.New("lenght of value exceeds maximum")
		}
	}
	return nil
}

func Validate(data interface{}, rules *Rules) error {
	dataType := reflect.TypeOf(data).String()

	switch dataType {
	case "string":
		return stringValidation(data.(string), rules)
	case "int":
		fmt.Println("getting int")
	default:
		fmt.Println(dataType)
	}

	fmt.Println("validated successfully", rules, dataType)
	return nil
}
