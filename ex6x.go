package main
import "fmt"
func main (){
	var n, i int
	fmt.Println("saisir un nombre")
	fmt.Scanln(&n)
	n = 0
	for i := 1; i >= 10; i++
	{n=n+i }
	fmt.Prinln("les nombres suivants sont:",n)