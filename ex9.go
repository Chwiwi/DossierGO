package main

import "fmt"

func main() {
	var n, max, s int
	s = 0
	max = 0
	for i := 1; i < 11; i++ {	
		fmt.Println("saisir un nombre")
	    fmt.Scanln(&n)
		if i == 1 || n > max {                     
			max = n 
			s = s + i		
		}
	fmt.Println("le max est: ",max,"la somme est: ",s)
		
	}
}