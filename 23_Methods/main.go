/*
DESAFIO 3: Métodos (Valor vs Ponteiro)
Objetivo: Dar comportamento a uma struct e alterar os seus dados internos.

Requisitos:
1. Cria uma struct 'Produto' com 'Nome' (string) e 'Preco' (float64).
2. Cria um método chamado 'AplicarDesconto' que recebe um PONTEIRO para Produto (*Produto) e um valor de desconto em percentagem. O método deve alterar o Preco original da struct.
3. Cria um método chamado 'ObterResumo' que recebe Produto por VALOR (Produto) e retorna uma string formatada: "O produto [Nome] custa [Preco]".
4. Na main, cria um Produto, imprime o resumo, aplica um desconto de 10% chamando o método, e imprime o resumo novamente.
*/
package main

import "fmt"

// Define a struct Produto e os seus métodos aqui
type Produto struct {
	Nome string
	Preco float64
}

func (p *Produto) AplicarDesconto(percentagem float64){
	valorDesconto := p.Preco * (percentagem/100)
	p.Preco = p.Preco - valorDesconto
}

func (p Produto) ObterResumo() string{
	return fmt.Sprintf("O produto %s custa %.2f", p.Nome, p.Preco)
}
func main() {
	// A tua implementação começa aqui
	fmt.Println("A iniciar Desafio 3: Methods...")
	meuProduto := Produto{
		Nome: "Teclado",
		Preco: 150.00,
	}
	fmt.Println(meuProduto.ObterResumo())
	meuProduto.AplicarDesconto(10.0)
	fmt.Println(meuProduto.ObterResumo())
}
