package ListaJava11;


public class Main {
    public static void main(String[] args) {
        FuncionarioHorista horista = new FuncionarioHorista("João", 28, 160, 25, 20);
        FuncionarioAssalariado assalariado = new FuncionarioAssalariado("Maria", 35, 5000, 500);
        Gerente gerente = new Gerente("Carlos", 45, 8000, 1500, "TI");

        System.out.println("Detalhes do funcionário horista:\n" + horista.exibirDetalhes());
        System.out.printf("Salário do funcionário horista (%s): R$%.2f\n\n", horista.nome, horista.calcularSalario());

        System.out.println("Detalhes do funcionário assalariado:\n" + assalariado.exibirDetalhes());
        System.out.printf("Salário do funcionário assalariado (%s): R$%.2f\n\n", assalariado.nome, assalariado.calcularSalario());

        System.out.println("Detalhes do gerente:\n" + gerente.exibirDetalhes());
        System.out.printf("Salário do gerente (%s): R$%.2f\n", gerente.nome, gerente.calcularSalario());
    }
}
