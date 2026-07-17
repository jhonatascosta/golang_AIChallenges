/*
DESAFIO 10: O Combo Final
Objetivo: Juntar Variáveis, Constantes, Tipos, Funções, Arrays e Loops em um único programa!

Requisitos:
1. Declare uma constante chamada 'Limite' com o valor 10.
2. Crie uma função chamada 'ehMaiorQueLimite' que recebe um 'int' e retorna um 'bool' indicando se o número é maior que a constante 'Limite'.
3. Na função main, crie um array de tamanho 5 com os números: 2, 8, 15, 23, 7.
4. Use um loop 'for range' (ignorando o índice com '_') para percorrer o array.
5. Dentro do loop, passe o número atual para a função 'ehMaiorQueLimite'. Se retornar true, imprima "[Número] é MAIOR que o limite". Se false, imprima "[Número] é MENOR ou IGUAL ao limite".
*/
package main

import "fmt"

// Declare sua função aqui
const Limite int = 10

func ehMaiorQueLimite (numero int) bool {
	if numero > Limite {
		return  true
	} else {
		return false
	}
}
func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 10: Boss Battle...")
	numeros := [5]int{2,8,15,23,7}
	for _, valor := range numeros {
		resultado := ehMaiorQueLimite(valor)
		if resultado == true {
			fmt.Printf("%d é MAIOR que o limite\n", valor)
		} else {
			fmt.Printf("%d é MENOR ou IGUAL ao limite\n", valor)
		}
	}
}
