package ListaJava11;


public class FuncionarioAssalariado extends Funcionario {
    public double salarioMensal;
    protected double bonus;

    public FuncionarioAssalariado(String nome, int idade, double salarioMensal, double bonus) {
        super(nome, idade);
        this.salarioMensal = salarioMensal;
        this.bonus = bonus;
    }

    @Override
    public double calcularSalario() {
        return salarioMensal + bonus;
    }

    @Override
    public String exibirDetalhes() {
        return String.format("Nome: %s, Idade: %d, Salário Mensal: R$%.2f, Bônus: R$%.2f",
                nome, idade, salarioMensal, bonus);
    }
}
