package main

import (
	"fmt"
)

func main() {
	//fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	for i := 0; i < 1000000; i++ {
		if i%9 == 0 && i%4 == 1 && i%5 == 4 && i%6 == 3 && i%7 == 5 && i%8 == 1 {
			fmt.Printf("%v\n", i)
		}
	}
	fmt.Println("run over !")
}
