package main

import "fmt"

func main(){
	var a [5]int
	fmt.Println(a)

	 a[4] = 100
	 a[0] = 200
	 a[3] = 300
	 a[2] = 400
	 a[1] = 500

	fmt.Println(a)

	var b [2][3]int

	fmt.Println(b)


	for i:=0;i<2;i++ {
		for j:=0;j<3;j++ {
			b[i][j] = i + j
		}
	}


	fmt.Println(b)

	}
