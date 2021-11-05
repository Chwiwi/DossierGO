package main

import "fmt"

func main() {
	var n, max, nombre, s int
    s = 0
	max = 0
	for i := 1; i < 666; i++ {
		fmt.Println("saisir un nombre")
		fmt.Scanln(&n)
		if i == 1 || n > max {
			max = n
		   nombre  =  i
		   s = s + n
		
		
		} else if n == 0 { 
			
			fmt.Println("le max est: ", max, "le nombre :", nombre,"la somme :"),s}
	     
	}
	
	fmt.Println("le max est: ", max, "position :", nombre)

}
