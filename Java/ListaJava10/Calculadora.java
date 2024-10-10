package ListaJava10;


public class Calculadora {
    public double somar(double... args) throws IllegalArgumentException {
        
        if (args.length == 2 || args.length == 3) {
            double soma = 0;
            for (double num : args) {
                soma += num;
            }
            return soma;
        } else {
            throw new IllegalArgumentException("O método 'somar' aceita apenas 2 ou 3 números.");
        }
    }
}
