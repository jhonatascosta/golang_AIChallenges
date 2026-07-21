/*
DESAFIO 3: Wrapping de Erros (errors.Is)
Objetivo: Encapsular um erro para dar contexto, mas ainda conseguir verificar sua origem.

Instruções:
1. Crie uma variável global: 'var ErrConexao = errors.New("falha na conexão com banco")'.
2. Crie uma função 'ConectarDB() error' que apenas retorna 'ErrConexao'.
3. Crie uma função 'IniciarServidor() error' que chama 'ConectarDB()'. Se der erro, use 'fmt.Errorf("erro ao iniciar: %w", err)' para embrulhar o erro.
4. No main, chame 'IniciarServidor()' e use 'errors.Is(err, ErrConexao)' para verificar se o erro original era de conexão.
*/
package main

import (
	"fmt"
	"errors"
)

var ErrConexao = errors.New("falha na conexão com banco")

func ConectarDB() error {
	return ErrConexao
}

func IniciarServidor() error {
	err := ConectarDB()
	if err != nil {
		return fmt.Errorf("erro ao iniciar: %w", err)
	}
	return nil
}

func main() {
	fmt.Println("Iniciando Desafio 3: Wrapping de Erros...")
	err := IniciarServidor()
	if err != nil {
		fmt.Println("Erro retornado:", err)

		if errors.Is(err, ErrConexao) {
			fmt.Println("-> Confirmado: a causa raiz foi falha de conexão com o banco.")
		} else {
			fmt.Println("-> Erro de origem desconhecida.")
		}
	}
}
