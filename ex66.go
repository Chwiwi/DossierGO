package main

import "fmt"

func main() {
	var n int
	fmt.Println("Saisir un nombre")
	fmt.Scanln(&n)
	for i := 1; i < 11; i++ {
		n = n + 1
		fmt.Println(n)
	}

}
