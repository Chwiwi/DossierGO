package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
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
	var a, b, c, delta int
	fmt.Println("saisir valeur de a")
	a = lireEntier()
	fmt.Println("saisir valeur de b")
	b = lireEntier()
	fmt.Println("saisir valeur de c")
	c = lireEntier()
	if a == 0 && b == 0 && c == 0{
		fmt.Println("infinité de solution")
	} else if a == 0 && b == 0 && c != 0{
		 fmt.Println("0 solution")}
	if a == 0 && b != 0{
    fmt.Println("1 solution -b/c: "), -b/c}
	else if a != 0
	{delta = b*b-4*a*c}
	if delta = 0
	{fmt.Println("1 solution -b/2a: "), -b/2*a}
	else if delta > 0
	{fmt.Println("2 solutions")}
	else
    {{fmt.Println("pas de solution")}
    

}
