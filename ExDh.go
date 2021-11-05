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
	var mp, r, mt int
	var   nbh float64
	const ph = 2
	fmt.Println("Veuillez saisir le montant payé")
	mp = lireEntier()
	fmt.Println("Veuillez saisir le nombre d'heure (en heure)")
	apayer = int()
	nbh = lireReel()

     if  nbh >= 0.5 {
	fmt.Println("le montant total à rendre est: ", mp-ph*nbh)
	mt = mp-ph*nbh
	fmt.Println("le nombre de 10dhs est: ", mt/10)
	r = mt%10
	fmt.Println("le nombre de 5dhs est: ", r/5)
	r = r%5
	fmt.Println("le nombre de 2dhs est: ", r/2)
	r = r%2
	fmt.Println("le nombre de 1dh est: ", r)}
	else {
	fmt.Println("Erreur")
	}}