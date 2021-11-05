package main

import "fmt"

func main() {
	var n, f int
	fmt.Println("saisir un nombre")
	fmt.Scanln(&n)
	f = 1
	for i := 2; i < n+1; i++ {
		f = f * i
	}
	fmt.Println("la factorielle est: ", f)
}
