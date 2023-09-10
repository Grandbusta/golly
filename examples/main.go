package main

import (
	"fmt"

	"github.com/Grandbusta/golly"
)

func main() {
	err := golly.Validate("Bu", &golly.Rules{
		Min: 3,
	},
	)
	fmt.Println(err)
}
