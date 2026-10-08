package main // Déclare ce fichier comme appartenant au programme exécutable.

import (
	"strconv" // Fournit les fonctions de conversion des nombres.
) // Termine la déclaration des imports.

func hex(words []string) []string { // Convertit en base 10 le mot précédant chaque marqueur "(hex)".
	var result []string // Prépare la liste des mots résultants.
	for _, w := range words { // Parcourt tous les mots reçus.
		if w == "(hex)" { // Repère le marqueur de conversion hexadécimale.
			if len(result) > 0 { // Vérifie qu'un mot précède le marqueur.
				lastwords := result[len(result)-1] // Récupère le dernier mot déjà ajouté.
				if valeur, err := strconv.ParseInt(lastwords, 16, 64); err == nil { // Essaie de lire ce mot comme un nombre hexadécimal.
					result[len(result)-1] = strconv.FormatInt(valeur, 10) // Remplace le mot par sa valeur décimale.
				} else { // Traite le cas où le mot précédent n'est pas un nombre hexadécimal valide.
					result = append(result, w) // Conserve le marqueur sans conversion.
				}
			} else { // Traite le marqueur placé au début de la liste.
				result = append(result, w) // Conserve ce marqueur sans conversion.
			}
		} else { // Traite un mot ordinaire.
			result = append(result, w) // Ajoute le mot au résultat.
		}
	}
	return result // Renvoie les mots après traitement.
} // Termine la fonction de conversion.