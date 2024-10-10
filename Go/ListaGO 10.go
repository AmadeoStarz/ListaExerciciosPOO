/* 10 - Sobrecarga de Métodos (Java) / Métodos com Nomes Diferentes (Python, Go)
#Implemente uma classe Calculadora com métodos somar que aceita diferentes números
#de parâmetros (dois ou três números). Em Golang, use funções com nomes diferentes
#para diferentes quantidades de parâmetros. */
package main

import (
    "fmt"
    "log"
    "strconv"
    "strings"
)

type Calculadora struct{}


func (c *Calculadora) Somar(nums []float64) (float64, error) {
    if len(nums) == 2 || len(nums) == 3 {
        var total float64
        for _, num := range nums {
            total += num
        }
        return total, nil
    }
    return 0, fmt.Errorf("o método 'Somar' aceita apenas 2 ou 3 números")
}

func main() {
    var calc Calculadora


    fmt.Print("Digite dois ou três números separados por espaço: ")
    var entrada string
    fmt.Scanln(&entrada)

    partes := strings.Fields(entrada)
    numeros := make([]float64, len(partes))

    for i, parte := range partes {
        num, err := strconv.ParseFloat(parte, 64)
        if err != nil {
            log.Fatalf("Erro ao converter número '%s': %v", parte, err)
        }
        numeros[i] = num
    }


    resultado, err := calc.Somar(numeros)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Soma: %.2f\n", resultado)
    }
}
