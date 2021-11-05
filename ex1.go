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
	var x, a, b int
	fmt.Println("Veuillez saisir a")
	a = lireEntier()
	fmt.Println("Veuillez saisir b")
	b = lireEntier()
	fmt.Println("Veuillez saisir x")
	x = lireEntier()
	if x >= a && x <= b {
		fmt.Println("Vrai")
	} else if x <= a && x >= b {
		fmt.Println("Vrai")
	} else {
		fmt.Println("Faux")
	}
}
