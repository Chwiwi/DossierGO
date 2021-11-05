package main

import "fmt"

func main() {
	var nbheure float64
	var Apayer, MontantPaye, Arendre, NP10, NP5, NP2, NP1, Rest int
	const PrixHeure = 2
	fmt.Println("Saisir nombre d'heure")
	fmt.Scanln(&nbheure)
	Apayer = int(nbheure * PrixHeure)
	fmt.Println("Saisir le montant payé")
	fmt.Scanln(&MontantPaye)
	Arendre = MontantPaye - Apayer
	NP10 = Arendre / 10
	Rest = Arendre % 10
	NP5 = Rest / 5
	Rest = Rest % 5
	NP2 = Rest / 2
	NP1 = Rest % 2
	fmt.Println("Le montant à rendre est: ", Arendre)
	fmt.Println("Nombre de pieces de 10: ", NP10, "de 5dhs: ", NP5, "de 2dhs ", NP2, "de 1dh: ", NP1)
}
