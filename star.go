package main

import "fmt"

func main(){

	var  i,j,level int = 0,0,40

	for i=0;i<level;i++{
		j = i + 1

		for row:=1;row<=j;row++{
			fmt.Printf("*")
		}
		fmt.Println("")
	}

}
