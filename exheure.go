package main
import "fmt"
func main (){
	var h, m int
	fmt.Println("Veuillez saisir l'heure")
	fmt.Scanln(&h)
	fmt.Println("Veuillez saisir les minutes")
	fmt.Scanln(&m)
	if  m >= 59{
		h = h+1 
		m = 0
	} else {
		m = m+1
		}
	fmt.Println("l'heure est:",h,"les minutes",m)
}