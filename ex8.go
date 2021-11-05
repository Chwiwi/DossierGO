package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lireChaine() string {
	var donneeLue string
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	scanner.Scan()             // lancement du scanner
	donneeLue = scanner.Text() // stockage du résultat du scanner dans une variable
	return donneeLue
}
func lireEntier() int {
	var nombre int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nombre, _ = strconv.Atoi(scanner.Text()) // conversion du type string en int
	return nombre
}

func lireReel() float64 {
	var donneeLue string
	var nombreLu float64
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	scanner.Scan()             // lancement du scanner
	donneeLue = scanner.Text() // stockage du résultat du scanner dans une variable
	nombreLu, _ = strconv.ParseFloat(donneeLue, 8)
	return nombreLu
}

func main() {
	var puh, nh, x, y float64
fmt.Println("Veuillez saisir le prix unitaire")
  puh = lireReel()
  fmt.Println("Veuillez saisir le nombre d'heure")
  nh = lireReel()
  x = nh-5
  y = nh-10
   switch{
  case nh > 0 && nh <= 5:
	  fmt.Println("le montant est: ", puh*nh*1.1)
  case nh > 5 && nh <= 10:
	  fmt.Println("le montant est: ", 5*puh*1.1+x*1.25)
 case nh > 10:
	  fmt.Println("le montant est: ",5*puh*1.1+5*puh*1.25+y*puh*1.5 )
  }

}
