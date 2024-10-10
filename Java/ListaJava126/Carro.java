package ListaJava126;

public class Carro {
    private String marca;
    private String modelo;
    private int ano;
    private Motor motor;
    private int velocidade;

    public Carro(String marca, String modelo, int ano, Motor motor) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.motor = motor;
        this.velocidade = 0; 
    }

    public void exibirDetalhes() {
        System.out.println("Marca: " + marca + ", Modelo: " + modelo + ", Ano: " + ano);
        motor.exibirDetalhes();
    }

    public void acelerar(int incremento) {
        velocidade += incremento;
        System.out.println("O carro acelerou para " + velocidade + " km/h");
    }

    public void frear(int decremento) {
        velocidade -= decremento;
        if (velocidade < 0) {
            velocidade = 0; 
        }
        System.out.println("O carro freou para " + velocidade + " km/h");
    }

    public void exibirVelocidade() {
        System.out.println("Velocidade atual: " + velocidade + " km/h");
    }
}
