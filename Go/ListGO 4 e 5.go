/* 4 -Herança Crie uma classe base Animal com métodos como som. Derive classes como
Cachorro e Gato que implementam o método som de maneiras diferentes
5 - Polimorfismo Utilize polimorfismo para criar uma função que receba uma lista de
objetos Animal e chame o método som de cada um, demonstrando comportamento
diferente para cada subclasse*/
package main

import (
    "fmt"
)


type Animal interface {
    Som() string
}


type Cachorro struct {
    Nome string
}


func (c Cachorro) Som() string {
    return fmt.Sprintf("%s faz: Au au!", c.Nome)
}


type Gato struct {
    Nome string
}

func (g Gato) Som() string {
    return fmt.Sprintf("%s faz: Miau!", g.Nome)
}

func FazerSons(animais []Animal) {
    for _, animal := range animais {
        fmt.Println(animal.Som())
    }
}

func main() {
    cachorro := Cachorro{Nome: "Genevaldo"}
    gato := Gato{Nome: "Pilantra"}
    
    listaAnimais := []Animal{cachorro, gato}
    
    FazerSons(listaAnimais)
}
