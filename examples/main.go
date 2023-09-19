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
		Alphanum:  true,
		Uppercase: true,
	}
	err := golly.Validate("ASAS", &rules)

	// user := User{
	// 	Firstname: "sdff",
	// 	Lastname:  "gr@",
	// }

	// err = golly.ValidateStruct(&user, golly.H{
	// 	&user.Firstname: &golly.Rules{Required: true, Min: 3, Max: 4},
	// 	&user.Lastname:  &golly.Rules{Length: 3, Alphanum: true},
	// })
	fmt.Println(err)

}
