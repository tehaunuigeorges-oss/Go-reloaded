package main 

import ("fmt"
		"os" 
		"strings") 
func main() { 
	if len(os.Args) < 3 { 
		fmt.Println("Format attendu: go run . <inputFile> <outputFile>") 
		return
	}
	inputFile := os.Args[1] 
	outputFile := os.Args[2] 

	contentBytes, err := os.ReadFile(inputFile)
	if err != nil { 
		fmt.Println("ERREUR DE LECTURE DU FICHIER :", err)
	}

	lines := strings.Split(string(contentBytes), "\n") 
	var processedLines []string 

	for _, line := range lines { 
		words := strings.Fields(line) 
		if len(words) == 0 { 
			processedLines = append(processedLines, "") 
			continue 
		}
		modifiedWords := up(words) 
		processedLines = append(processedLines, strings.Join(modifiedWords, " "))
	}

	finalText := strings.Join(processedLines, "\n") 

	err = os.WriteFile(outputFile, []byte(finalText), 0644) 
	if err != nil { 
		fmt.Println("Erreur écriture :", err) 
		return 
	}

	fmt.Println("Traitement terminé avec succès") 
}