/*
DESAFIO 6: O Combo Final (If, Switch, Defer)
Objetivo: Juntar os três conceitos em uma simulação de processamento de dados.

Requisitos:
1. Declare uma função chamada 'processarArquivo' que recebe um 'nomeArquivo' (string) e um 'tamanhoMB' (float64).
2. Logo no início da função, coloque um defer para imprimir: "Arquivo [nomeArquivo] fechado com segurança."
3. Use um 'if' simples para checar se 'tamanhoMB' é maior que 100.0. Se for, imprima "Erro: Arquivo muito grande!" e dê um 'return' antecipado (early return) para sair da função.
4. Se o arquivo tiver o tamanho válido, crie um 'switch' no 'nomeArquivo' (você pode testar strings inteiras!).
   - case "relatorio.pdf": imprima "Processando PDF..."
   - case "dados.csv": imprima "Lendo planilha..."
   - default: imprima "Formato desconhecido."
5. Na função main, chame 'processarArquivo' três vezes testando cenários diferentes (um arquivo grande, um PDF pequeno, e um formato desconhecido pequeno).
*/
package main

import "fmt"

// Crie a função processarArquivo aqui
func processarArquivo (nomeArquivo string, tamanhoMB float64) {
	defer fmt.Printf("Arquivo %s fechado com segurança.\n", nomeArquivo)
	if tamanhoMB > 100.0 {
		fmt.Println("Erro: Arquivo muito grande!")
		return
	}
	switch nomeArquivo {
	case "relatorio.pdf":
		fmt.Println("Processando PDF...")
	case "dados.csv":
		fmt.Println("Lendo planilha...")
	default:
		fmt.Println("Formato desconhecido.")
	}
}

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 6: Boss Battle de Controle de Fluxo...")
	processarArquivo("relatorio.pdf", 15.3)
	processarArquivo("dados.csv", 18.9)
	processarArquivo("teste.txt", 150.9)
}
