package main
import "fmt"
func main (){
	var j, m, a int 
	fmt.Println("saisir jour")
	fmt.Scanln(&j)
	fmt.Println("saisir les mois")
	fmt.Scanln(&m)
	fmt.Println("saisir année")
	fmt.Scanln(&a)
	switch{
	case m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10:
		 if j < 31
		 {j = j+1}
		else {m = m+1}	
	case m == 2:
	    if j < 29
		{j = j+1}
		else {m = m+1}
	case m == 4 || m == 6 || m == 9 || m == 11:
	    if j < 30
		{j = j+1}
		else {m == m+1}
	case  m == 12:
	    if j < 31
		{j = j+1}
		else {a = a+1
		      j = 0
			  m = 1}
	Println("Année:",a,"Mois",m,"Jour",j)
	}
