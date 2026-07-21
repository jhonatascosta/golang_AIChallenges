/*
DESAFIO 4: Polimorfismo com Interfaces
Objetivo: Criar uma função que aceita diferentes tipos genéricos.

Requisitos:
1. Cria uma interface 'Notificador' com um método 'Enviar(mensagem string)'.
2. Cria duas structs: 'Email' e 'SMS' (podem estar vazias ou ter um campo de endereço/número).
3. Implementa o método 'Enviar' para ambas as structs (faz apenas um Println diferente para cada uma, ex: "A enviar email: ..." e "A enviar SMS: ...").
4. Cria uma função normal chamada 'Alertar' que receba como parâmetro (n Notificador, mensagem string) e chame n.Enviar(mensagem).
5. Na main, cria um Email e um SMS, e passa ambos para a função 'Alertar'.
*/
package main

import "fmt"

// Define a interface, structs e métodos aqui
type Notificador interface{
	Enviar(mensagem string)
}

type Email struct {
	EnderecoEmail string
}

type SMS struct {
	Numero string
}

func (e Email) Enviar(mensagem string){
	fmt.Printf("Enviando Email para [%s]: %s\n", e.EnderecoEmail, mensagem)
}

func (s SMS) Enviar(mensagem string){
	fmt.Printf("Enviando SMS para [%s]: %s\n", s.Numero, mensagem)
}

func Alertar(n Notificador, mensagem string){
	n.Enviar(mensagem)
}

func main() {
	// A tua implementação começa aqui
	fmt.Println("A iniciar Desafio 4: Interfaces...")
	meuEmail := Email{EnderecoEmail: "jhonata@exemplo.com"}
	meuSms := SMS{Numero: "+351912345678"}

	Alertar(meuEmail, "Bem-vindo ao teu novo servidor em Go!")
	Alertar(meuSms, "Alerta de segurança: Login efetuado com sucesso.")
}
