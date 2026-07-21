/*
DESAFIO 8: Iterando com Range
Objetivo: Usar a forma mais idiomática do Go para percorrer coleções.

Requisitos:
1. Declare um array de strings já inicializado com 3 nomes. 
   Dica: nomes := [3]string{"Alice", "Bob", "Carlos"}
2. Use a sintaxe 'for indice, valor := range nomes' para percorrer o array.
3. Imprima em cada ciclo algo como: "Índice X tem o valor Y".
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 8: For Range...")
	nomes := [3]string{"Jhon", "Carlos", "Rodrigo"}
	for indice, valor := range nomes {
		fmt.Printf("Índice %d tem o valor %s\n", indice, valor)
	}
}
