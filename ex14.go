package main

import "fmt"

func main() {
	var f, s int
	
	fmt.Println(" saisir la borne basse")
	fmt.Scanln(&f)
	fmt.Println("saisir la borne haute")
	fmt.Scanln(&s)
	for i := f; i < s-f; i++ {
		if f%2 != 0 { f = f +1
	    } else {
		f = f + 2 }
	fmt.Println("les nombres paires sont: ",f)
	}

}