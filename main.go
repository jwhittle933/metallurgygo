package main

import (
	"fmt"
)

func main() {
	args := ParseFlags()
	files := GetFiles(args.Dir).Filter(args.In)

	for _, file := range files {
		fmt.Println(file.Path)
	}
}
