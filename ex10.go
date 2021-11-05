package main

import "fmt"

func main() {
	var n, max, pos int

	max = 0
	for i := 1; i < 11; i++ {
		fmt.Println("saisir 10 nombres")
		fmt.Scanln(&n)
		if i == 1 || n > max {
			max = n
			pos =  i
		}
	}
	
	fmt.Println("le max est: ", max, "position :", pos)

}
