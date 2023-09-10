package golly

import (
	"fmt"
	"reflect"

	"github.com/Grandbusta/golly/types"
)

type Rules struct {
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

func stringValidation(str string, r *Rules) error {
	// Allowed rules for string
	stringRules := &types.StringRules{
		Min:    r.Min,
		Max:    r.Max,
		Length: r.Length,
	}
	return types.StringValidate(str, stringRules)
}
