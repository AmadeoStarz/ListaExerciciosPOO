package ListaJava7;

import java.util.ArrayList;
import java.util.List;

class Escola {
    private String nome;
    List<Professor> professores;

    public Escola(String nome) {
        this.nome = nome;
        this.professores = new ArrayList<>();
    }

    public void adicionarProfessor(Professor professor) {
        if (!professores.contains(professor)) {
            professores.add(professor);
        }
    }

    @Override
    public String toString() {
        return nome;
    }
}
