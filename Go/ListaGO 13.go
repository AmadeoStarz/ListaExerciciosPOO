/* 13 - Métodos Estáticos Adicione um método estático à classe Matematica que calcula o
fatorial de um número. Em Python, utilize @staticmethod, em Java static, e em Golang crie
ma função regular na struct.*/
package main

import (
	"fmt"
	"log"
)

type Matematica struct{}


func (m Matematica) Fatorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * m.Fatorial(n-1)
}

func main() {
	var numero int
	fmt.Print("Digite um número para calcular o fatorial: ")
	_, err := fmt.Scanf("%d", &numero)
	if err != nil {
		log.Fatalf("Erro ao ler entrada: %v", err)
	}

	if numero < 0 {
		log.Fatalf("O número deve ser não negativo.")
	}

	m := Matematica{}
	resultado := m.Fatorial(numero)
	fmt.Printf("O fatorial de %d é %d\n", numero, resultado)
}
