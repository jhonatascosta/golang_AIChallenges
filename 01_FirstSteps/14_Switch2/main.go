/*
DESAFIO 3: Switch sem Expressão (Switch True)
Objetivo: Usar o switch como uma alternativa limpa para cadeias longas de if/else.

Requisitos:
1. Declare uma variável 'idade' e atribua uma idade qualquer a ela.
2. Crie um 'switch' vazio (apenas digite 'switch {').
3. Crie os seguintes cases com expressões booleanas lógicas:
   - case idade < 13: imprima "Criança"
   - case idade < 18: imprima "Adolescente"
   - case idade < 60: imprima "Adulto"
   - default: imprima "Idoso"
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 3: Switch sem condição...")
	idade := 37
	switch {
	case idade < 13:
		fmt.Println("Criança")
	case idade < 18:
		fmt.Println("Adolescente")
	case idade < 60:
		fmt.Println("Adulto")
	default:
		fmt.Println("Idoso")
	}
}
