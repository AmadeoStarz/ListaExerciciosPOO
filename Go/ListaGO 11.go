/*11 -Classes Abstratas Crie uma classe abstrata Funcionario com um método abstrato
calcularSalario. Derive classes como FuncionarioHorista e FuncionarioAssalariado que
implementam calcularSalario de formas distintas*/
package main

import (
    "fmt"
)


type Funcionario interface {
    CalcularSalario() float64
    ExibirDetalhes() string
}


type FuncionarioHorista struct {
    Nome            string
    Idade           int
    HorasTrabalhadas float64
    ValorHora       float64
    HorasExtras     float64
}


func (f *FuncionarioHorista) CalcularSalario() float64 {
    salarioBase := f.HorasTrabalhadas * f.ValorHora
    salarioExtras := f.HorasExtras * (f.ValorHora * 1.5)
    return salarioBase + salarioExtras
}


func (f *FuncionarioHorista) ExibirDetalhes() string {
    return fmt.Sprintf("Nome: %s, Idade: %d, Horas Trabalhadas: %.2f, Horas Extras: %.2f, Valor Hora: R$%.2f",
        f.Nome, f.Idade, f.HorasTrabalhadas, f.HorasExtras, f.ValorHora)
}


type FuncionarioAssalariado struct {
    Nome          string
    Idade         int
    SalarioMensal float64
    Bonus         float64
}

func (f *FuncionarioAssalariado) CalcularSalario() float64 {
    return f.SalarioMensal + f.Bonus
}

func (f *FuncionarioAssalariado) ExibirDetalhes() string {
    return fmt.Sprintf("Nome: %s, Idade: %d, Salário Mensal: R$%.2f, Bônus: R$%.2f",
        f.Nome, f.Idade, f.SalarioMensal, f.Bonus)
}

type Gerente struct {
    FuncionarioAssalariado
    Departamento string
}


func (g *Gerente) ExibirDetalhes() string {
    return fmt.Sprintf("Nome: %s, Idade: %d, Salário Mensal: R$%.2f, Bônus: R$%.2f, Departamento: %s",
        g.Nome, g.Idade, g.SalarioMensal, g.Bonus, g.Departamento)
}

func main() {

    horista := FuncionarioHorista{"João", 28, 160, 25, 20}
    fmt.Println("Detalhes do funcionário horista:")
    fmt.Println(horista.ExibirDetalhes())
    fmt.Printf("Salário do funcionário horista (%s): R$%.2f\n\n", horista.Nome, horista.CalcularSalario())


    assalariado := FuncionarioAssalariado{"Maria", 35, 5000, 500}
    fmt.Println("Detalhes do funcionário assalariado:")
    fmt.Println(assalariado.ExibirDetalhes())
    fmt.Printf("Salário do funcionário assalariado (%s): R$%.2f\n\n", assalariado.Nome, assalariado.CalcularSalario())


    gerente := Gerente{
        FuncionarioAssalariado: FuncionarioAssalariado{"Carlos", 45, 8000, 1500},
        Departamento: "TI",
    }
    fmt.Println("Detalhes do gerente:")
    fmt.Println(gerente.ExibirDetalhes())
    fmt.Printf("Salário do gerente (%s): R$%.2f\n", gerente.Nome, gerente.CalcularSalario())
}
