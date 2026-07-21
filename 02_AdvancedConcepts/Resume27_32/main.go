package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// 1. Struct sem atributos
type FiltroMaiusculo struct{}

// 2. Implementando io.Writer
func (f FiltroMaiusculo) Write(p []byte) (n int, err error) {
	texto := string(p)

	// Se contiver "bug", recusa e retorna erro
	if strings.Contains(texto, "bug") {
		return 0, errors.New("código proibido detectado")
	}

	// Caso contrário, converte para maiúsculas e imprime
	maiusculo := bytes.ToUpper(p)
	fmt.Println(string(maiusculo))

	return len(p), nil
}

func main() {
	fmt.Println("Iniciando a Boss Battle 2: Filtro Customizado io.Writer...")

	var seuFiltro FiltroMaiusculo

	// Primeira chamada: texto limpo
	_, err := fmt.Fprint(seuFiltro, "meu primeiro backend")
	if err != nil {
		fmt.Println("Erro:", err)
	}

	// Segunda chamada: contém "bug"
	_, err = fmt.Fprint(seuFiltro, "tem um bug aqui")
	if err != nil {
		fmt.Println("Erro:", err)
	}
}
