package ListaJava13;


import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        
        System.out.print("Digite um número para calcular o fatorial: ");
        int numero = scanner.nextInt();
        
        if (numero < 0) {
            System.out.println("O fatorial não está definido para números negativos.");
        } else {
            long resultado = Matematica.fatorial(numero);
            System.out.printf("O fatorial de %d é %d%n", numero, resultado);
        }
        
        scanner.close();
    }
}
