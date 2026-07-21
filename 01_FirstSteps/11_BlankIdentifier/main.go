/*
DESAFIO 9: Ignorando Retornos com '_'
Objetivo: Aprender a usar o underscore (_) para descartar valores não utilizados.

Requisitos:
1. Declare um array de float64 com 4 preços diferentes.
2. Crie um loop 'for range' para iterar sobre esse array.
3. O 'range' sempre retorna o índice e o valor. Desta vez, você NÃO quer usar o índice, apenas os preços.
4. Use o blank identifier '_' no lugar da variável do índice para evitar o erro de "variável declarada e não usada".
5. Imprima apenas os preços.
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 9: Blank Identifier...")
	precos := [4]float64{15.90, 10.30, 17.20, 25.90}
	for _, valor := range precos {
		fmt.Println(valor)
	}
}
