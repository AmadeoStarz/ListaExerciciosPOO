package ListaJava11;


public class Gerente extends FuncionarioAssalariado {
    private String departamento;

    public Gerente(String nome, int idade, double salarioMensal, double bonus, String departamento) {
        super(nome, idade, salarioMensal, bonus);
        this.departamento = departamento;
    }

    @Override
    public String exibirDetalhes() {
        return String.format("Nome: %s, Idade: %d, Salário Mensal: R$%.2f, Bônus: R$%.2f, Departamento: %s",
                nome, idade, salarioMensal, bonus, departamento);
    }
}
