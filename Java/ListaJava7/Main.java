package ListaJava7;

import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        Escola escola1 = new Escola("Escola Turma da Moranguinho");
        Escola escola2 = new Escola("Escola St Merlin");

        Professor professor1 = new Professor("Antonio Bandeiras");
        Professor professor2 = new Professor("Mario Sergio Cortela");
        Professor professor3 = new Professor("Ana Maria Braga");
        Professor professor4 = new Professor("Joaquim Phoenix");

        professor1.adicionarEscola(escola1);
        professor2.adicionarEscola(escola2);
        professor3.adicionarEscola(escola1);
        professor4.adicionarEscola(escola2);

        System.out.println("Professores na " + escola1 + ": " + getNomesProfessores(escola1));
        System.out.println("Professores na " + escola2 + ": " + getNomesProfessores(escola2));
        System.out.println(professor1 + " leciona em: " + getNomesEscolas(professor1));
        System.out.println(professor2 + " leciona em: " + getNomesEscolas(professor2));
        System.out.println(professor3 + " leciona em: " + getNomesEscolas(professor3));
        System.out.println(professor4 + " leciona em: " + getNomesEscolas(professor4));
    }

    private static String getNomesProfessores(Escola escola) {
        List<Professor> professores = escola.professores; 
        List<String> nomes = new ArrayList<>();
        for (Professor p : professores) {
            nomes.add(p.toString());
        }
        return nomes.toString();
    }

    private static String getNomesEscolas(Professor professor) {
        List<Escola> escolas = professor.escolas; 
        List<String> nomes = new ArrayList<>();
        for (Escola e : escolas) {
            nomes.add(e.toString());
        }
        return nomes.toString();
    }
}
