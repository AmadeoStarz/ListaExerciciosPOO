package Listajava3e15;

public class Main {
    public static void main(String[] args) {
        try {
            ContaBancaria conta = new ContaBancaria("João", 1000);
            System.out.println("Titular: " + conta.getTitular());
            System.out.println("Saldo inicial: R$" + conta.getSaldo());

            conta.depositar(500);
            System.out.println("Saldo após depósito: R$" + conta.getSaldo());

            conta.sacar(300);
            System.out.println("Saldo após saque: R$" + conta.getSaldo());

            conta.sacar(1500); // Isso irá lançar a exceção
        } catch (SaldoInsuficienteException e) {
            System.out.println(e.getMessage());
        }
    }
}

