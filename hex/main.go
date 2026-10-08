package main // Déclare ce fichier comme programme exécutable.

import ("fmt" // Fournit les fonctions d'affichage.
		"os" // Fournit l'accès aux arguments, fichiers et erreurs.
		"strings") // Fournit les fonctions de traitement des chaînes.
func main() { // Point d'entrée du programme.
	if len(os.Args) < 3 { // Vérifie que les fichiers d'entrée et de sortie sont indiqués.
		fmt.Println("Format attendu: go run . <inputFile> <outputFile>") // Affiche la syntaxe d'utilisation.
		return // Arrête le programme si les arguments sont incomplets.
	}
	inputFile := os.Args[1] // Récupère le chemin du fichier à lire.
	outputFile := os.Args[2] // Récupère le chemin du fichier à écrire.

	contentBytes, err := os.ReadFile(inputFile) // Lit tout le contenu du fichier d'entrée.
	if err != nil { // Vérifie si la lecture a échoué.
		fmt.Println("ERREUR DE LECTURE DU FICHIER :", err) // Affiche l'erreur de lecture.
		return // Arrête le programme en cas d'erreur.
	}

	lines := strings.Split(string(contentBytes), "\n") // Sépare le texte en lignes.
	var processedLines []string // Prépare la liste des lignes transformées.

	for _, line := range lines { // Parcourt chaque ligne du fichier.
		words := strings.Fields(line) // Sépare la ligne en mots en ignorant les espaces superflus.
		if len(words) == 0 { // Vérifie si la ligne ne contient aucun mot.
			processedLines = append(processedLines, "") // Conserve la ligne vide dans le résultat.
			continue // Passe directement à la ligne suivante.
		}
		modifiedWords := hex(words) // Applique la conversion hexadécimale aux mots.
		processedLines = append(processedLines, strings.Join(modifiedWords, " ")) // Recompose la ligne avec des espaces.
	}

	finalText := strings.Join(processedLines, "\n") // Recompose le texte avec ses retours à la ligne.

	err = os.WriteFile(outputFile, []byte(finalText), 0644) // Écrit le résultat dans le fichier de sortie.
	if err != nil { // Vérifie si l'écriture a échoué.
		fmt.Println("Erreur écriture :", err) // Affiche l'erreur d'écriture.
		return // Arrête le programme en cas d'erreur.
	}

	fmt.Println("Traitement terminé avec succès") // Confirme que le traitement est terminé.
} // Termine la fonction principale.