package main

import (
	"bufio"
	"os"
	"strconv"
)

func lireChaine() string {
	var donneeLue string
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	scanner.Scan()                        // lancement du scanner
	donneeLue = scanner.Text()            // stockage du résultat du scanner dans une variable
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
	scanner.Scan()                        // lancement du scanner
	donneeLue = scanner.Text()            // stockage du résultat du scanner dans une variable
	nombreLu, _ = strconv.ParseFloat(donneeLue, 8)
	return nombreLu
}

func main() {
	var M, Dh200, Dh100, Dh50, Dh20, Dh10, Dh5, Dh2, Dh1, R1, R2, R3, R4, R5, R6, R7 int
	println("Veuillez saisir le montant")
	M = lireEntier()
	Dh200 = M / 200
	println("les nombres de billet de 200 est:", Dh200)
	R1 = M % 200
	Dh100 = R1 / 100
	println("les nombres de billet de 100 est:", Dh100)
	R2 = R1 % 100
	Dh50 = R2 / 50
	println("les nombres de billet de 50 est:", Dh50)
	R3 = R2 % 50
	Dh20 = R3 / 20
	println("les nombres de billet de 20 est:", Dh20)
	R4 = R3 % 20
	Dh10 = R4 / 10
	println("les nombres de piece de 10 est:", Dh10)
	R5 = R4 % 10
	Dh5 = R5 / 5
	println("les nombres de piece de 5dh est:", Dh5)
	R6 = R5 % 5
	Dh2 = R6 / 2
	println("les nombres de piece de 2dh est:", Dh2)
	R7 = R6 % 2
	Dh1 = R7
	println("les nombres de piece de 1dh est:", Dh1)

}
