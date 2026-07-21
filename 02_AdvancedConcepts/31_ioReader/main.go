package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println("Iniciando Desafio 5: Explorando io.Reader em pedaços...")

	// 1. Cria um Reader a partir de uma string
	reader := strings.NewReader("Golang é incrivelmente rápido e poderoso para Backends!")

	// 2. Buffer pequeno de 8 bytes
	buffer := make([]byte, 8)

	// 3. Loop infinito lendo pedaço por pedaço
	for {
		n, err := reader.Read(buffer)

		if n > 0 {
			// Imprime apenas os bytes efetivamente lidos (buffer[:n])
			fmt.Printf("Lido %d bytes: %q\n", n, string(buffer[:n]))
		}

		// 4. Trata o erro EOF para encerrar o loop
		if err == io.EOF {
			fmt.Println("Fim do arquivo (EOF) alcançado!")
			break
		}
		if err != nil {
			fmt.Println("Erro inesperado:", err)
			break
		}
	}
}
