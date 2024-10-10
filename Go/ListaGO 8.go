/* 8 - Agregação Modele uma classe Empresa que agregue uma lista de objetos Empregado.
Cada empregado deve ter atributos como nome, cargo, e salario. */
package main

import (
    "fmt"
)


type Endereco struct {
    Rua    string
    Cidade string
    Estado string
    Cep    string
}


type Pessoa struct {
    Nome     string
    Idade    int
    Endereco *Endereco
}


func (p *Pessoa) AddEndereco(endereco Endereco) {
    p.Endereco = &endereco
}


func (p Pessoa) MostrarInfo() {
    fmt.Printf("Nome: %s, Idade: %d\n", p.Nome, p.Idade)
    if p.Endereco != nil {
        fmt.Printf("Endereço: %s, %s, %s, %s\n", p.Endereco.Rua, p.Endereco.Cidade, p.Endereco.Estado, p.Endereco.Cep)
    } else {
        fmt.Println("Endereço não disponível")
    }
}


type Empregado struct {
    Pessoa
    Cargo   string
    Salario float64
}


func (e Empregado) MostrarInfo() {
    e.Pessoa.MostrarInfo()
    fmt.Printf("Cargo: %s, Salário: R$%.2f\n", e.Cargo, e.Salario)
}


type Empresa struct {
    Nome        string
    CNPJ        string
    Funcionarios []Empregado
}


func (emp *Empresa) ContratarFuncionario(funcionario Empregado) {
    emp.Funcionarios = append(emp.Funcionarios, funcionario)
}

func (emp Empresa) ListarFuncionarios() {
    fmt.Printf("Funcionários da empresa %s:\n", emp.Nome)
    for _, funcionario := range emp.Funcionarios {
        funcionario.MostrarInfo()
    }
}

func main() {
    
    endereco1 := Endereco{"Rua dos Pinhais", "Tokyo", "Ceará", "12345-678"}
    empregado1 := Empregado{Pessoa: Pessoa{"Mario", 38}, Cargo: "Gerente", Salario: 8500.00}
    empregado1.AddEndereco(endereco1)

    endereco2 := Endereco{"Paulista", "São Paulo", "São Paulo", "99999-999"}
    empregado2 := Empregado{Pessoa: Pessoa{"Luigi", 25}, Cargo: "Analista", Salario: 4500.00}
    empregado2.AddEndereco(endereco2)


    empresa := Empresa{"Empresa ABC", "12.345.678/0001-99"}
    empresa.ContratarFuncionario(empregado1)
    empresa.ContratarFuncionario(empregado2)


    empresa.ListarFuncionarios()
}
