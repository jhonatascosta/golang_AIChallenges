/*
DESAFIO 4: Sub-Slices e Segurança de Memória
Objetivo: Criar slices a partir de slices e testar o mecanismo de segurança do Go (Bounds Check).

Requisitos:
1. Crie um slice com os números de 1 a 5.
2. Crie um 'subSlice' pegando apenas do índice 1 até o índice 3 (sintaxe slice[inicio:fim]). Imprima o subSlice.
3. Agora, tente imprimir o índice 10 do seu slice original (ex: fmt.Println(slice[10])).
4. Rode o código. O Go vai gerar um "panic: index out of range" (Bounds Check em ação!). 
5. Comente a linha que deu erro para o código compilar, e escreva num comentário logo abaixo explicando por que o Go prefere dar um "panic" e quebrar a aplicação ao invés de ler uma memória aleatória.
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 4: Fatiamento e Bounds Check...")
	slice := []int{1,2,3,4,5}
	subSlice := slice[1:3]
	fmt.Println(subSlice)
	/* fmt.Println(slice[10])
		Go verifica se o comprimento passado está dentro do limite, se sim, irá imprimir normalmente, se não, ele da panic para evitar estouro de buffer.
	*/
}
