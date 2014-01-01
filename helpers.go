package box

import (
	"fmt"
)

func log(items ...interface{}) {
	fmt.Print("*******\n")
	for _, item := range items {
		fmt.Printf("%s\n", item)
	}
	fmt.Print("\n")
}
