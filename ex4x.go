package main
import "fmt"
func main(){
	var h, s, m int
	fmt.Println("saisir heure")
	fmt.Scanln(&h)
	fmt.Println("saisir minutes")
	fmt.Scanln(&m)
	fmt.Println("saisir secondes")
	fmt.Scanln(&s)
	
if h == 0 && m ==0 && s == 0{
		for i := 1; i < 7200; i++ {
		if s >= 59{
			m=m+1
			s=0
	} else {s = s + 1}
		if m >= 59 {
			h = h+1
			m=0
			s=0
		} else{m=m+1}
		if h >= 24{
			h = 0}
	    else {h=h+1}
	
		fmt.Println("l'heure:",h,"minutes:",m,"secondes",s)
	}
}
		
	}}}
	