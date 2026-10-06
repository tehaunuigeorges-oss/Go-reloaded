package main

import ("os" 
 		"fmt"
		"strings"
	)

func main() {
	if len(os.Args)<3 {
		fmt.Println("Format attendu: go run . <inputFile> <outputFile>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	contentBytes, err:= os.ReadFile(inputFile)
	if err !=nil {
		fmt.Println("ERREUR DE LECTURE DU FICHIER", err)
		return
	}
	words := strings.Fields(string(contentBytes))
	modifieWords := hex(words)
	finalText := strings.Join(modifiedWords, " ")
	err = os.WriteFile(outputFile, []byte(finalText), 0644)
	if err != nil {
		fmt.Println("Erreur écriture :", err)
		return
	}
	fmt.Println("traitement terminé avec succés")
}
