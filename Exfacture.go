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
	var pu, nba, r, ar float64
	const tva = 0.2
	fmt.Println("Veuillez saisir le montant HT")
	pu = lireReel()
	fmt.Println("Veuillez saisir le nombre d'article")
	nba = lireReel()
	fmt.Println("Le total HT avant remise est: ", pu*nba)
	fmt.Println("Veuillez saisir le taux de remise (par exemple pour 20% ecrire 0.2)")
	r = lireReel()
	fmt.Println("Le total après remise est: ", pu*nba-pu*nba*r)
	ar = pu*nba - pu*nba*r
	fmt.Println("Le montant de la tva est: ", ar*tva)
	fmt.Println("Le total TTC est: ", ar+ar*tva)
}
