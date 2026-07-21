/*
DESAFIO 2: Structs como Erros (Custom Errors)
Objetivo: Criar um erro personalizado que carrega mais informações além da mensagem.

Instruções:
1. Crie uma struct chamada 'ErroAPI' contendo 'StatusCode int' e 'Mensagem string'.
2. Faça 'ErroAPI' implementar a interface 'error' (basta criar o método 'Error() string' usando pointer receiver).
3. Crie uma função 'BuscarUsuario(id int) error' que, se o id for menor que 1, retorna um '&ErroAPI{StatusCode: 404, Mensagem: "usuário não encontrado"}'.
4. No main, chame a função e imprima o erro.
*/
package main

import (
	"fmt"
)

type ErroAPI struct {
	StatusCode int
	Mensagem string
}

func (e *ErroAPI) Error() string {
	return fmt.Sprintf("erro %d: %s", e.StatusCode, e.Mensagem)
}

func BuscarUsuario(id int) error {
	if id < 1{
		return &ErroAPI{StatusCode: 404, Mensagem: "usuário não encontrado"}
	}
	return nil
}

func main() {
	fmt.Println("Iniciando Desafio 2: Erros Customizados...")
	// Seu código de teste aqui
	err := BuscarUsuario(-1)
	fmt.Println("Ocorreu um erro:", err)	
	err = BuscarUsuario(5)
	if err == nil {
		fmt.Println("Usuário encontrado com sucesso!")
	}
}
