package main

import (
	"fmt"

	"github.com/Grandbusta/golly"
)

func main() {
	err := golly.Validate("Busta", &golly.Rules{
		Required: false,
		Min:      3,
	},
	)
	fmt.Println(err)
}
