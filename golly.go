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
		return stringValidation(data.(string), rules, "value")
	case "int":
		fmt.Println("getting int")
	default:
		fmt.Println(dataType)
	}

	fmt.Println("validated successfully", rules, dataType)
	return nil
}

type H map[interface{}]*Rules

func ValidateStruct(data interface{}, h H) {
	// dataType := reflect.TypeOf(data).String()
	// // fmt.Println(data, h, dataType)
	v := reflect.ValueOf(data).Elem()

	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Field(i).UnsafeAddr(), v.Field(i), v.Type().Field(i).Anonymous)
	}

	for k, _ := range h {
		kptr := reflect.ValueOf(k).Pointer()
		fmt.Println("gvd", reflect.ValueOf(k).Elem().Type(), kptr)
	}
	// t := reflect.TypeOf(data)
	// for i := 0; i < t.NumField(); i++ {
	// 	fmt.Printf("%+v\n", t.Field(i))
	// 	fmt.Println(t.Field(i).Type)
	// }
}

func stringValidation(str string, r *Rules, label string) error {
	// Allowed rules for string
	stringRules := &types.StringRules{
		Min:    r.Min,
		Max:    r.Max,
		Length: r.Length,
	}
	return types.StringValidate(str, stringRules, label)
}
