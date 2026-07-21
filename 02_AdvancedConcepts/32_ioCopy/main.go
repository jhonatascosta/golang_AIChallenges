package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Iniciando Desafio 6: io.Copy em ação...")

	// 1. Cria um Reader a partir de uma string
	leitor := strings.NewReader("Mensagem confidencial: O Go é o melhor!\n")

	// 2 e 3. Copia diretamente do leitor (io.Reader) para os.Stdout (io.Writer)
	bytesEscritos, err := io.Copy(os.Stdout, leitor)

	// 4. Trata o erro retornado
	if err != nil {
		fmt.Println("Erro ao copiar dados:", err)
		return
	}

	fmt.Printf("\nCópia concluída! %d bytes transferidos.\n", bytesEscritos)
}
