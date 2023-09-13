package golly

import (
	"errors"
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

type field struct {
	value reflect.Value
	rules *Rules
}

func newField(fieldValue reflect.Value, fieldRules *Rules) *field {
	return &field{
		value: fieldValue,
		rules: fieldRules,
	}
}

func validate(data interface{}, rules *Rules, label string) error {
	dataType := reflect.TypeOf(data).String()

	switch dataType {
	case "string":
		return stringValidation(data.(string), rules, label)
	case "int":
		fmt.Println("getting int")
	default:
		fmt.Println(dataType)
	}
	fmt.Println("validated successfully", rules, dataType)
	return nil
}

func Validate(data interface{}, rules *Rules) error {
	return validate(data, rules, "value")
}

type H map[interface{}]*Rules

func ValidateStruct(data interface{}, h H) error {
	value := reflect.ValueOf(data)

	if value.Kind() != reflect.Ptr || !value.IsNil() && value.Elem().Kind() != reflect.Struct {
		// must be a pointer to a struct
		return errors.New("Must be a pointer to a struct")
	}
	if value.IsNil() {
		// treat a nil struct pointer as valid
		return nil
	}
	value = value.Elem()

	for k, r := range h {
		fieldValue := reflect.ValueOf(k)
		if fieldValue.Kind() != reflect.Ptr {
			// return NewInternalError(ErrFieldPointer(i))
		}
		fs := findStructField(value, newField(fieldValue, r))
		err := validate(fieldValue.Elem().Interface(), r, fs.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func findStructField(structValue reflect.Value, field *field) *reflect.StructField {
	fptr := field.value.Pointer()
	for i := 0; i < structValue.NumField(); i++ {
		sf := structValue.Type().Field(i)
		if fptr == structValue.Field(i).UnsafeAddr() {
			if sf.Type == field.value.Elem().Type() {
				// fmt.Println(fptr, sf, sf.Type, field.value.Elem(), field.rules)
				return &sf
			}
		}

		// check anonymous field

	}
	return nil
}

func stringValidation(str string, r *Rules, label string) error {
	// Allowed rules for string
	stringRules := &types.StringRules{
		Min:      r.Min,
		Max:      r.Max,
		Length:   r.Length,
		Required: r.Required,
	}
	return types.StringValidate(str, stringRules, label)
}
