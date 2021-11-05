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
	var n float64
	fmt.Println("saisir la note")
	n = lireReel()
	switch {
	case n >= 0 && n < 10:
		fmt.Println("Echec")
	case n >= 10 && n < 12:
		fmt.Println("Passable")
	case n >= 12 && n < 14:
		fmt.Println("Assez bien")
	case n >= 14 && n < 16:
		fmt.Println("Bien")
	case n >= 16 && n <= 20:
		fmt.Println("Très Bien")
	default:
		fmt.Println("Erreur")
	}

}
