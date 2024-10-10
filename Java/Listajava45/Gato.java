package Listajava45;


public class Gato extends Animal {
    public Gato(String nome) {
        super(nome);
    }

    @Override
    public String som() {
        return getNome() + " faz: Miau!";
    }
}
