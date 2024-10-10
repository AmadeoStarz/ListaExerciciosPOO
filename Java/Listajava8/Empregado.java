package Listajava8;


public class Empregado extends Pessoa {
    private String cargo;
    private double salario;

    public Empregado(String nome, int idade, String cargo, double salario) {
        super(nome, idade);
        this.cargo = cargo;
        this.salario = salario;
    }

    @Override
    public void mostrarInfo() {
        super.mostrarInfo();
        System.out.printf("Cargo: %s, Salário: R$%.2f%n", cargo, salario);
    }
}
