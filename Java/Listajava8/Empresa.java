package Listajava8;

import java.util.ArrayList;
import java.util.List;


public class Empresa {
    private String nome;
    @SuppressWarnings("unused")
    private String cnpj;
    private List<Empregado> funcionarios;

    public Empresa(String nome, String cnpj) {
        this.nome = nome;
        this.cnpj = cnpj;
        this.funcionarios = new ArrayList<>();
    }

    public void contratarFuncionario(Empregado funcionario) {
        funcionarios.add(funcionario);
    }

    public void listarFuncionarios() {
        System.out.println("Funcionários da empresa " + nome + ":");
        for (Empregado funcionario : funcionarios) {
            funcionario.mostrarInfo();
        }
    }
}
