package main

import (
	"fmt"
	)

func main(){

	i := 2

	fmt.Print(" Write ", i, " As ")
	switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
	}

}
