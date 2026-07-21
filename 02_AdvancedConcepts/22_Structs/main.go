/*
DESAFIO 2: Criando a tua primeira Struct
Objetivo: Agrupar dados relacionados num único tipo personalizado.

Requisitos:
1. Fora da função main, define um novo tipo chamado 'Utilizador' usando 'type Utilizador struct { ... }'.
2. A struct deve ter os campos: 'Nome' (string), 'Idade' (int) e 'Ativo' (bool).
3. Na função main, cria duas variáveis do tipo Utilizador:
   - Uma usando a sintaxe de campos nomeados (ex: Utilizador{Nome: "João", ...}).
   - Outra instanciando primeiro a variável (var u Utilizador) e preenchendo os campos um por um (u.Nome = "Maria").
4. Imprima as duas variáveis.
*/
package main

import "fmt"

// Define a tua struct Utilizador aqui
type Utilizador struct {
	Nome string
	Idade int
	Ativo bool
}

func main() {
	// A tua implementação começa aqui
	fmt.Println("A iniciar Desafio 2: Structs...")
	u1 := Utilizador{Nome: "José", Idade: 23, Ativo: true}
	var u2 Utilizador
	u2.Nome = "Maria"
	u2.Idade = 45
	u2.Ativo = false 

	fmt.Println(u1, u2)
}
