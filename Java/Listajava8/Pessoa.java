package Listajava8;


public class Pessoa {
    private String nome;
    private int idade;
    private Endereco endereco;

    public Pessoa(String nome, int idade) {
        this.nome = nome;
        this.idade = idade;
        this.endereco = null;
    }

    public void addEndereco(Endereco endereco) {
        this.endereco = endereco;
    }

    public void mostrarInfo() {
        System.out.println("Nome: " + nome + ", Idade: " + idade);
        if (endereco != null) {
            System.out.println("Endereço: " + endereco.getRua() + ", " + endereco.getCidade() + ", " + endereco.getEstado() + ", " + endereco.getCep());
        } else {
            System.out.println("Endereço não disponível");
        }
    }
}
