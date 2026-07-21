package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Iniciando Desafio 4: Explorando io.Writer...")

	// 1. Cria (ou sobrescreve) o arquivo log.txt
	arquivo, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}

	// 2. Garante que o arquivo será fechado ao final da função
	defer arquivo.Close()

	// 3. Escreve no arquivo usando o método Write (parte da interface io.Writer)
	bytesEscritos, err := arquivo.Write([]byte("Log gerado com sucesso!"))
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Printf("Arquivo criado com sucesso! %d bytes escritos.\n", bytesEscritos)
}
