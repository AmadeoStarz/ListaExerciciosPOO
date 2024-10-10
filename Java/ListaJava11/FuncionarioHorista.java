package ListaJava11;


public class FuncionarioHorista extends Funcionario {
    private int horasTrabalhadas;
    private double valorHora;
    private int horasExtras;

    public FuncionarioHorista(String nome, int idade, int horasTrabalhadas, double valorHora, int horasExtras) {
        super(nome, idade);
        this.horasTrabalhadas = horasTrabalhadas;
        this.valorHora = valorHora;
        this.horasExtras = horasExtras;
    }

    @Override
    public double calcularSalario() {
        double salarioBase = horasTrabalhadas * valorHora;
        double salarioExtras = horasExtras * (valorHora * 1.5);
        return salarioBase + salarioExtras;
    }

    @Override
    public String exibirDetalhes() {
        return String.format("Nome: %s, Idade: %d, Horas Trabalhadas: %d, Horas Extras: %d, Valor Hora: R$%.2f",
                nome, idade, horasTrabalhadas, horasExtras, valorHora);
    }
}
