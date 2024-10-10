package Listajava8;

public class Main {
    public static void main(String[] args) {

        Endereco endereco1 = new Endereco("Rua dos Pinhais", "Tokyo", "Ceará", "12345-678");
        Empregado empregado1 = new Empregado("Mario", 38, "Gerente", 8500.00);
        empregado1.addEndereco(endereco1);

        Endereco endereco2 = new Endereco("Paulista", "São Paulo", "São Paulo", "99999-999");
        Empregado empregado2 = new Empregado("Luigi", 25, "Analista", 4500.00);
        empregado2.addEndereco(endereco2);


        Empresa empresa = new Empresa("Empresa ABC", "12.345.678/0001-99");
        empresa.contratarFuncionario(empregado1);
        empresa.contratarFuncionario(empregado2);

        empresa.listarFuncionarios();
    }
}
