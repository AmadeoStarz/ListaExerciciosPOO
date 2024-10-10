package ListaJava126;

public class Main {
    public static void main(String[] args) {
        Motor motor1 = new Motor("Gasolina", 150);
        Carro carro1 = new Carro("Toyota", "Corolla", 2020, motor1);
        
        Carro carro2 = new Carro("Honda", "Civic", 2021, new Motor("Etanol", 140));
        Carro carro3 = new Carro("Ford", "Focus", 2019, new Motor("Diesel", 160));

        carro1.exibirDetalhes();
        carro2.exibirDetalhes();
        carro3.exibirDetalhes();

        carro1.acelerar(30);
        carro1.acelerar(20);
        
        carro1.frear(10);
        carro1.frear(50);
        
        carro1.exibirVelocidade();
    }
}
