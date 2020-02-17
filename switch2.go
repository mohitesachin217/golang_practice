package main

import (
	"fmt"
	"time"
	)


func main(){
		switch time.Now().Weekday(){
			case time.Saturday, time.Sunday:
				fmt.Println("It's a weekend")
				fmt.Println("and It's a",time.Now().Weekday())
			default:
				fmt.Println("It's a weekday")
				fmt.Println("And IT's a ",time.Now().Weekday())
			}
	}
