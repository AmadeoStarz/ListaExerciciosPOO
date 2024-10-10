package ListaJava11;


public abstract class Funcionario {
    protected String nome;
    protected int idade;

    public Funcionario(String nome, int idade) {
        this.nome = nome;
        this.idade = idade;
    }

    public abstract double calcularSalario();
    
    public abstract String exibirDetalhes();
}
