package ListaJava12;


public class Produto {
    private String nome;
    private double preco;

    public Produto(String nome, double preco) {
        this.nome = nome;
        this.preco = preco;
    }

    public Produto add(Produto other) {
        if (other != null) {
            return new Produto("Combo de " + this.nome + " e " + other.nome, this.preco + other.preco);
        }
        return null; 
    }

    @Override
    public boolean equals(Object obj) {
        if (obj instanceof Produto) {
            Produto other = (Produto) obj;
            return this.preco == other.preco;
        }
        return false;
    }

    @Override
    public String toString() {
        return String.format("Produto: %s, Preço: R$%.2f", nome, preco);
    }

    public void aplicarDesconto(double percentual) {
        double desconto = this.preco * (percentual / 100);
        this.preco -= desconto;
        System.out.printf("Desconto de %.2f%% aplicado em %s. Novo preço: R$%.2f%n", percentual, nome, preco);
    }

    public String exibirDetalhes() {
        return String.format("Produto: %s, Preço: R$%.2f", nome, preco);
    }

    public String getNome() {
        return nome;
    }

    public double getPreco() {
        return preco;
    }
}
