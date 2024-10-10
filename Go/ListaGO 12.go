/* 12 - Sobrecarga de Operadores (Python) / Métodos de Conveniência (Java, Go) Em Python,
sobrecarregue o operador + para somar dois objetos Produto baseados no preço. Em Java
e Golang, crie métodos que permitam essa funcionalidade.*/
package main

import (
    "fmt"
)


type Produto struct {
    Nome  string
    Preco float64
}


func (p *Produto) Somar(outro Produto) Produto {
    return Produto{
        Nome:  fmt.Sprintf("Combo de %s e %s", p.Nome, outro.Nome),
        Preco: p.Preco + outro.Preco,
    }
}


func (p *Produto) AplicarDesconto(percentual float64) {
    desconto := p.Preco * (percentual / 100)
    p.Preco -= desconto
    fmt.Printf("Desconto de %.2f%% aplicado em %s. Novo preço: R$%.2f\n", percentual, p.Nome, p.Preco)
}


func (p *Produto) String() string {
    return fmt.Sprintf("Produto: %s, Preço: R$%.2f", p.Nome, p.Preco)
}

func main() {
    produto1 := Produto{"Notebook", 3000.00}
    produto2 := Produto{"Smartphone", 1500.00}

    fmt.Println(produto1.String())
    fmt.Println(produto2.String())

    comboProdutos := produto1.Somar(produto2)
    fmt.Println("\n" + comboProdutos.String())

    produto1.AplicarDesconto(10)
}
