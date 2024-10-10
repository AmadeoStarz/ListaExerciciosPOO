/* Associação Crie uma classe Escola e uma classe Professor. A associação deve permitir
que uma escola tenha vários professores, e um professor possa lecionar em várias escolas.*/
package main

import (
    "fmt"
)


type Escola struct {
    Nome       string
    Professores []Professor
}


func (e *Escola) AdicionarProfessor(professor Professor) {
    for _, p := range e.Professores {
        if p.Nome == professor.Nome {
            return 
        }
    }
    e.Professores = append(e.Professores, professor)
}


func (e Escola) String() string {
    return e.Nome
}


type Professor struct {
    Nome    string
    Escolas []Escola
}


func (p *Professor) AdicionarEscola(escola Escola) {
    for _, e := range p.Escolas {
        if e.Nome == escola.Nome {
            return 
        }
    }
    p.Escolas = append(p.Escolas, escola)
    escola.AdicionarProfessor(*p) 
}


func (p Professor) String() string {
    return p.Nome
}

func main() {
    escola1 := Escola{Nome: "Escola Turma da Moranguinho"}
    escola2 := Escola{Nome: "Escola St Merlin"}

    professor1 := Professor{Nome: "Antonio Bandeiras"}
    professor2 := Professor{Nome: "Mario Sergio Cortela"}
    professor3 := Professor{Nome: "Ana Maria Braga"}
    professor4 := Professor{Nome: "Joaquim Phoenix"}

    professor1.AdicionarEscola(escola1)
    professor2.AdicionarEscola(escola2)
    professor3.AdicionarEscola(escola1)
    professor4.AdicionarEscola(escola2)


    fmt.Printf("Professores na %s: ", escola1.String())
    for _, p := range escola1.Professores {
        fmt.Printf("%s ", p.String())
    }
    fmt.Println()

    fmt.Printf("Professores na %s: ", escola2.String())
    for _, p := range escola2.Professores {
        fmt.Printf("%s ", p.String())
    }
    fmt.Println()

    
    fmt.Printf("%s leciona em: ", professor1.String())
    for _, e := range professor1.Escolas {
        fmt.Printf("%s ", e.String())
    }
    fmt.Println()

    fmt.Printf("%s leciona em: ", professor2.String())
    for _, e := range professor2.Escolas {
        fmt.Printf("%s ", e.String())
    }
    fmt.Println()

    fmt.Printf("%s leciona em: ", professor3.String())
    for _, e := range professor3.Escolas {
        fmt.Printf("%s ", e.String())
    }
    fmt.Println()

    fmt.Printf("%s leciona em: ", professor4.String())
    for _, e := range professor4.Escolas {
        fmt.Printf("%s ", e.String())
    }
    fmt.Println()
}
