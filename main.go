package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	var height, width, haiku_count int
	haiku_count = len(os.Args[1:])
	haiku_ary := make([]string, haiku_count)
	haiku_space := ""
	for _, arg := range os.Args[1:] {
		height += len(arg) - 1
		haiku_ary[haiku_count-width-1] = haiku_space + arg
		for i := 0; i < len(arg)/2; i++ {
			haiku_space += " "
		}
		width++
	}
	for i := 0; i < len(haiku_ary[0]); i++ {
		for j := 0; j < len(haiku_ary); j++ {
			// fmt.Println(haiku_ary[j])
			if i < len([]rune(haiku_ary[j])) {
				// fmt.Print(haiku_ary[j][i])
				// fmt.Print(haiku_ary[j][i])
				// fmt.Print([]rune(haiku_ary[j])[i])
				fmt.Print(string([]rune(haiku_ary[j])[i]))
			} else {
				fmt.Print("  ")
			}

		}
		fmt.Println()
	}
}
