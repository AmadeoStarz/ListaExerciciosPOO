/*#1 - Classes e Objetos Básicos Crie uma classe Carro com atributos como marca, modelo, e
#ano. Instancie três objetos dessa classe e exiba os detalhes de cada um
#2 - Métodos Adicione um método acelerar e frear à classe Carro que altere um atributo
#velocidade. Crie um método para exibir a velocidade atual.*/
package main

import (
    "fmt"
)

type Motor struct {
    Tipo     string
    Potencia int
}


func (m Motor) ExibirDetalhes() {
    fmt.Printf("Tipo do Motor: %s, Potência: %d cv\n", m.Tipo, m.Potencia)
}

type Carro struct {
    Marca     string
    Modelo    string
    Ano       int
    Motor     Motor
    Velocidade int
}


func (c Carro) ExibirDetalhes() {
    fmt.Printf("Marca: %s, Modelo: %s, Ano: %d\n", c.Marca, c.Modelo, c.Ano)
    c.Motor.ExibirDetalhes()
}


func (c *Carro) Acelerar(incremento int) {
    c.Velocidade += incremento
    fmt.Printf("O carro acelerou para %d km/h\n", c.Velocidade)
}


func (c *Carro) Frear(decremento int) {
    c.Velocidade -= decremento
    if c.Velocidade < 0 {
        c.Velocidade = 0
    }
    fmt.Printf("O carro freou para %d km/h\n", c.Velocidade)
}


func (c Carro) ExibirVelocidade() {
    fmt.Printf("Velocidade atual: %d km/h\n", c.Velocidade)
}

func main() {
    motor1 := Motor{Tipo: "Gasolina", Potencia: 150}

    carro1 := Carro{Marca: "Toyota", Modelo: "Corolla", Ano: 2020, Motor: motor1}

    carro1.ExibirDetalhes()

    carro1.Acelerar(30)
    carro1.Acelerar(20)

    carro1.Frear(10)
    carro1.Frear(50)

    carro1.ExibirVelocidade()

    
    motor2 := Motor{Tipo: "Álcool", Potencia: 130}
    carro2 := Carro{Marca: "Honda", Modelo: "Civic", Ano: 2019, Motor: motor2}

    motor3 := Motor{Tipo: "Elétrico", Potencia: 200}
    carro3 := Carro{Marca: "Tesla", Modelo: "Model 3", Ano: 2021, Motor: motor3}


    carro2.ExibirDetalhes()
    carro3.ExibirDetalhes()
}
