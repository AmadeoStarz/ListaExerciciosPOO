/*#3 - Encapsulamento Implemente uma classe ContaBancaria com atributos saldo, titular e
#métodos depositar e sacar. Use encapsulamento para proteger o atributo saldo.
#15 - Exceções/Erros Personalizados Crie uma classe de exceção personalizada
#SaldoInsuficienteException em Java e Python, ou error em Golang, que seja lançada
#quando houver uma tentativa de saque superior ao saldo disponível na classe
#ContaBancaria.*/ 
package main

import (
    "fmt"
)

type SaldoInsuficienteError struct {
    Mensagem string
}

func (e *SaldoInsuficienteError) Error() string {
    return e.Mensagem
}

type ContaBancaria struct {
    titular  string
    saldo    float64
}

func NewContaBancaria(titular string, saldoInicial float64) *ContaBancaria {
    return &ContaBancaria{titular: titular, saldo: saldoInicial}
}

func (c *ContaBancaria) Depositar(valor float64) {
    if valor > 0 {
        c.saldo += valor
        fmt.Printf("Depósito de R$%.2f realizado com sucesso.\n", valor)
    } else {
        fmt.Println("Valor de depósito inválido.")
    }
}

func (c *ContaBancaria) Sacar(valor float64) error {
    if valor > c.saldo {
        return &SaldoInsuficienteError{"Valor de saque superior ao saldo disponível."}
    }
    if valor > 0 {
        c.saldo -= valor
        fmt.Printf("Saque de R$%.2f realizado com sucesso.\n", valor)
        return nil
    }
    return fmt.Errorf("valor de saque inválido")
}

func (c *ContaBancaria) GetSaldo() float64 {
    return c.saldo
}

func (c *ContaBancaria) GetTitular() string {
    return c.titular
}

func main() {
    conta := NewContaBancaria("João", 1000)
    fmt.Printf("Titular: %s\n", conta.GetTitular())
    fmt.Printf("Saldo inicial: R$%.2f\n", conta.GetSaldo())

    conta.Depositar(500)
    fmt.Printf("Saldo após depósito: R$%.2f\n", conta.GetSaldo())

    err := conta.Sacar(300)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("Saldo após saque: R$%.2f\n", conta.GetSaldo())

    err = conta.Sacar(1500) 
    if err != nil {
        fmt.Println(err)
    }
}
