package main

import (
	"fmt"

	"github.com/Grandbusta/golly"
)

type User struct {
	Firstname string
	Lastname  string
}

func main() {
	rules := golly.Rules{
		Min: 3,
		Max: 6,
	}
	err := golly.Validate("Bustats", &rules)

	fmt.Println(err)

	user := User{
		Firstname: "busta",
		Lastname:  "grand",
	}

	golly.ValidateStruct(&user, golly.H{
		&user.Firstname: &golly.Rules{Required: true},
		&user.Lastname:  &golly.Rules{},
	})
}
