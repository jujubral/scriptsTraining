package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	apiURL := "http://localhost:5127/"
	wordlistFilePath := "wordlists/endpoints.txt"

	wordlist, err := readLines(wordlistFilePath)
	if err != nil {
		fmt.Printf("Erro ao ler a wordlist: %s\n", err)
		os.Exit(1)
	}

	outputFilePath := "recon.txt"
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo de saída: %s\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	for _, endpoint := range wordlist {
		fullURL := fmt.Sprintf("%s%s", apiURL, endpoint)

		resp, err := http.Get(fullURL)
		if err != nil {
			fmt.Printf("Erro ao acessar o endpoint %s: %s\n", fullURL, err)
			continue
		}
		if resp.StatusCode != 404 {
			fmt.Printf("Endpoint encontrado: %s (status %d)\n", fullURL, resp.StatusCode)
			fmt.Fprintf(outputFile, "%s (status %d)\n", fullURL, resp.StatusCode)
		}
		resp.Body.Close()
	}

	fmt.Println("Enumeração completa")
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
