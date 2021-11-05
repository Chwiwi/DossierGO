package main

import "fmt"

func main() {
	var n, max, pos int

	max = 0
	for i := 1; i < 666; i++ {
		fmt.Println("saisir un nombre")
		fmt.Scanln(&n)
		if i == 1 || n > max {
			max = n
			pos =  i
		} else if n == 0 { fmt.Println("le max est: ", max, "position :", pos)}
	     
	}
	
	fmt.Println("le max est: ", max, "position :", pos)

}
