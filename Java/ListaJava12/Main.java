package ListaJava12;


public class Main {
    public static void main(String[] args) {
        Produto produto1 = new Produto("Notebook", 3000.00);
        Produto produto2 = new Produto("Smartphone", 1500.00);
        Produto produto3 = new Produto("Tablet", 1500.00);

        System.out.println(produto1);
        System.out.println(produto2);
        System.out.println(produto3);

        Produto comboProdutos = produto1.add(produto2);
        System.out.printf("%n%s%n", comboProdutos);

        produto1.aplicarDesconto(10);

        if (produto2.equals(produto3)) {
            System.out.printf("%n%s e %s têm o mesmo preço.%n", produto2.getNome(), produto3.getNome());
        } else {
            System.out.printf("%n%s e %s têm preços diferentes.%n", produto2.getNome(), produto3.getNome());
        }


        Produto somaInvalida = produto1.add(null); 
        if (somaInvalida == null) {
            System.out.println("\nErro: Não é possível somar um produto com um tipo não Produto.");
        }
    }
}
