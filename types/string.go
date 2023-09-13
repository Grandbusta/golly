package types

import (
	"fmt"
	"strconv"

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

var messages = map[string]string{
	"string.empty":    "{label} is not allowed to be empty",
	"string.max":      "{label} must be less than or equal to {limit} characters long",
	"string.min":      "{label} length must be at least {limit} characters long",
	"string.alphanum": "{label} must only contain alpha-numeric characters",
	"string.base":     "{label} must be a string",
	"string.minmax":   "max must me greater than min",
	"string.length":   "{label} length must be {limit} characters long",
}

func StringValidate(str string, r *StringRules, label string) error {
	fmt.Println(r.Required, len(str))
	if r.Required {
		if len(str) <= 0 {
			return common.Error(messages["string.empty"], fmt.Sprintf("\"%v\"", label))
		}
	}
	if r.Max != 0 && r.Max < r.Min {
		return common.Error(messages["string.minmax"])
	}
	if r.Min != 0 {
		if len(str) < r.Min {
			return common.Error(messages["string.min"], fmt.Sprintf("\"%v\"", label), strconv.Itoa(r.Min))
		}
	}
	if r.Max != 0 {
		if len(str) > r.Max {
			return common.Error(messages["string.max"], fmt.Sprintf("\"%v\"", label), strconv.Itoa(r.Max))
		}
	}
	if r.Length != 0 && len(str) != r.Length {
		return common.Error(messages["string.length"], fmt.Sprintf("\"%v\"", label), strconv.Itoa(r.Length))
	}
	return nil
}

// messages: {
// 	'string.alphanum': '{{#label}} must only contain alpha-numeric characters',
// 	'string.base': '{{#label}} must be a string',
// 	'string.base64': '{{#label}} must be a valid base64 string',
// 	'string.creditCard': '{{#label}} must be a credit card',
// 	'string.dataUri': '{{#label}} must be a valid dataUri string',
// 	'string.domain': '{{#label}} must contain a valid domain name',
// 	'string.email': '{{#label}} must be a valid email',
// 	'string.empty': '{{#label}} is not allowed to be empty',
// 	'string.guid': '{{#label}} must be a valid GUID',
// 	'string.hex': '{{#label}} must only contain hexadecimal characters',
// 	'string.hexAlign': '{{#label}} hex decoded representation must be byte aligned',
// 	'string.hostname': '{{#label}} must be a valid hostname',
// 	'string.ip': '{{#label}} must be a valid ip address with a {{#cidr}} CIDR',
// 	'string.ipVersion': '{{#label}} must be a valid ip address of one of the following versions {{#version}} with a {{#cidr}} CIDR',
// 	'string.isoDate': '{{#label}} must be in iso format',
// 	'string.isoDuration': '{{#label}} must be a valid ISO 8601 duration',
// 	'string.length': '{{#label}} length must be {{#limit}} characters long',
// 	'string.lowercase': '{{#label}} must only contain lowercase characters',
// 	'string.max': '{{#label}} length must be less than or equal to {{#limit}} characters long',
// 	'string.min': '{{#label}} length must be at least {{#limit}} characters long',
// 	'string.normalize': '{{#label}} must be unicode normalized in the {{#form}} form',
// 	'string.token': '{{#label}} must only contain alpha-numeric and underscore characters',
// 	'string.pattern.base': '{{#label}} with value {:[.]} fails to match the required pattern: {{#regex}}',
// 	'string.pattern.name': '{{#label}} with value {:[.]} fails to match the {{#name}} pattern',
// 	'string.pattern.invert.base': '{{#label}} with value {:[.]} matches the inverted pattern: {{#regex}}',
// 	'string.pattern.invert.name': '{{#label}} with value {:[.]} matches the inverted {{#name}} pattern',
// 	'string.trim': '{{#label}} must not have leading or trailing whitespace',
// 	'string.uri': '{{#label}} must be a valid uri',
// 	'string.uriCustomScheme': '{{#label}} must be a valid uri with a scheme matching the {{#scheme}} pattern',
// 	'string.uriRelativeOnly': '{{#label}} must be a valid relative uri',
// 	'string.uppercase': '{{#label}} must only contain uppercase characters'
// }
