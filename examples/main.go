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
	// rules := golly.Rules{
	// 	Min: 3,
	// 	Max: 6,
	// }
	// err := golly.Validate("Bustats", &rules)

	user := User{
		Firstname: "sdff",
		Lastname:  "gra",
	}

	err := golly.ValidateStruct(&user, golly.H{
		&user.Firstname: &golly.Rules{Required: true, Min: 3, Max: 4},
		&user.Lastname:  &golly.Rules{Length: 3},
	})
	fmt.Println(err)

}
