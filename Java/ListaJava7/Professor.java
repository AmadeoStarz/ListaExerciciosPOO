package ListaJava7;

import java.util.ArrayList;
import java.util.List;

class Professor {
    private String nome;
    List<Escola> escolas;

    public Professor(String nome) {
        this.nome = nome;
        this.escolas = new ArrayList<>();
    }

    public void adicionarEscola(Escola escola) {
        if (!escolas.contains(escola)) {
            escolas.add(escola);
            escola.adicionarProfessor(this);
        }
    }

    @Override
    public String toString() {
        return nome;
    }
}
