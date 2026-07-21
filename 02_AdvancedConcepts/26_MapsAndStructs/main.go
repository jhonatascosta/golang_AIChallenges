/*
DESAFIO 6: Base de Dados em Memória
Objetivo: Combinar Maps e Structs para indexar dados complexos.

Requisitos:
1. Cria uma struct 'Cliente' com 'ID' (int), 'Nome' (string) e 'Saldo' (float64).
2. Na main, cria um map onde a chave é o ID (int) e o valor é a struct Cliente (map[int]Cliente).
3. Insere 3 clientes no map usando os seus IDs como chave.
4. Cria um loop 'for range' para iterar sobre o map e imprimir o Nome e o Saldo de cada cliente.
5. Lembra-te: O loop em maps no Go é não-ordenado! Ao correres várias vezes, a ordem dos prints pode mudar. (Escreve num comentário se notaste isto a acontecer).
*/
package main

import "fmt"

// Define a struct Cliente aqui
type Cliente struct {
	ID int
	Nome string
	Saldo float64
}

func main() {
	// A tua implementação começa aqui
	fmt.Println("A iniciar Desafio 6: Maps de Structs...")
	client := make(map[int]Cliente)
	client[1] = Cliente{ID: 1, Nome: "José", Saldo: 73892.22}
	client[2] = Cliente{ID: 2, Nome: "Ruan", Saldo: 3882.23}
	client[3] = Cliente{ID: 3, Nome: "Roberto", Saldo: 313.42}

	for id, cliente := range client{
		fmt.Printf("ID da chave: %d | Nome: %s | Saldo: %.2fR$\n", id, cliente.Nome, cliente.Saldo)
	}
}
