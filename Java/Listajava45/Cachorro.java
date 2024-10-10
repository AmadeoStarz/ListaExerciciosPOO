package Listajava45;


public class Cachorro extends Animal {
    public Cachorro(String nome) {
        super(nome);
    }

    @Override
    public String som() {
        return getNome() + " faz: Au au!";
    }
}
