/*
DESAFIO 5: Type Switch e Assertions
Objetivo: Extrair e testar o tipo de dado guardado numa interface genérica.

Requisitos:
1. Cria uma função 'ProcessarDado' que receba um parâmetro genérico: 'dado any' (ou 'dado interface{}').
2. Dentro da função, usa a sintaxe de "Type Switch" (switch v := dado.(type)) para verificar o tipo.
3. Cria 'cases' para 'int', 'string' e 'bool'.
   - Se for int, imprime "É um número mágico: [v*2]" (multiplica por 2).
   - Se for string, imprime "Recebi um texto: [v]".
   - Adiciona um 'default' para tipos desconhecidos.
4. Na main, chama a função passando um número, um texto, um booleano e um float64 (para cair no default).
*/
package main

import "fmt"

// Cria a função ProcessarDado aqui
func ProcessarDado(dado any){
	switch v := dado.(type) {
	case int:
		fmt.Println("É um número mágico:", v*2)
	case string:
		fmt.Println("Recebi um texto:", v)
	default:
		fmt.Println("Outro tipo!")
	}
}

func main() {
	// A tua implementação começa aqui
	fmt.Println("A iniciar Desafio 5: Type Assertions...")
	ProcessarDado(12)
	ProcessarDado("Qual é?")
	ProcessarDado(true)
	ProcessarDado(16.78)
}

