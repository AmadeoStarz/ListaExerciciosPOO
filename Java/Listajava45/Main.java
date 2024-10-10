package Listajava45;


public class Main {
    public static void fazerSons(Animal[] animais) {
        for (Animal animal : animais) {
            System.out.println(animal.som());
        }
    }

    public static void main(String[] args) {
        Cachorro cachorro = new Cachorro("Genevaldo");
        Gato gato = new Gato("Pilantra");
        
        Animal[] listaAnimais = {cachorro, gato};
        
        fazerSons(listaAnimais);
    }
}
