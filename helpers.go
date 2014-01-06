package box

import (
	"fmt"
)

func Log(items ...interface{}) {
	log(items...)
}

func log(items ...interface{}) {
	fmt.Printf("******* %s *******\n", items[0])
	for _, item := range items[1:] {
		fmt.Printf("%s\n", item)
	}
	fmt.Print("\n")
}
