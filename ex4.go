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
	var x, y, z, t int
	fmt.Println("saisir valeur de x")
	x = lireEntier()
	fmt.Println("saisir valeur de y")
	y = lireEntier()
	fmt.Println("saisir valeur de z")
	z = lireEntier()
	if y < x {
		t = x
		x = y
		y = t} 
		if z < x {
		t = z
		z = y
		y = x
		x = t
	} else if z < y {
		t = z
		z = y
		y = t
	}
	fmt.Println(x, y, z)
	

}
