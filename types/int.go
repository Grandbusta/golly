package types

type IntRules struct {
	Min int
	Max int
}

var intMessages = map[string]string{}

func IntValidation(num int, r *IntRules, label string) error {
	if num > r.Max {
		// num must be less than or equal to max
	}
	if num < r.Min {
		// num must be greater than or equal to min
	}
	return nil
}
