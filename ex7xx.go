package main

import "fmt"

func main() {
	var n, s int
	fmt.Println("saisir un nombre")
	fmt.Scanln(&n)
	s = 0
	for i := 1; i < n+1; i++ {
		s = s + i
	}
	fmt.Println("la somme est: ",s)
	
	}
