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
		Email: true,
	}
	err := golly.Validate("ASAS@m.com", &rules)

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
