package ListaJava10;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Calculadora calc = new Calculadora();
        Scanner scanner = new Scanner(System.in);

        System.out.print("Digite dois ou três números separados por espaço: ");
        String entrada = scanner.nextLine();
        
        String[] partes = entrada.split(" ");
        double[] numeros = new double[partes.length];

        for (int i = 0; i < partes.length; i++) {
            numeros[i] = Double.parseDouble(partes[i]);
        }

        try {
            double resultado = calc.somar(numeros);
            System.out.printf("Soma: %.2f%n", resultado);
        } catch (IllegalArgumentException e) {
            System.out.println(e.getMessage());
        } finally {
            scanner.close();
        }
    }
}
