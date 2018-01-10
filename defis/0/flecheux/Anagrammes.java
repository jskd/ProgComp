import java.util.ArrayList;
import java.util.Arrays;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Anagrammes {

    //Trie le mot dans l'ordre alphabetique
    private static String ordonnerMot(String m){
        char tmpTab[] = m.toCharArray();
        Arrays.sort(tmpTab);
        return new String(tmpTab);
    }

    //trie les anagrammes trouves par ordre alphabetique
    private static void trierResultats(ArrayList<ArrayList<String>> anagrammes){
        for (ArrayList<String> list : anagrammes)
            list.sort(String::compareToIgnoreCase); //provient de https://stackoverflow.com/questions/5815423/sorting-arraylist-in-alphabetical-order-case-insensitive
    }

    private static void afficherResultats(String[] args, ArrayList<ArrayList<String>> anagrammes){
        for (int i = 1; i < args.length; ++i){
            /* Il y a une liste de moins dans anagrammes qu'il y a d'arguments dans args, le premier argument etant le dictionnaire.
             * On doit donc faire anagrammes.get(i-1) pour obtenir la liste d'anagrammes correspondant a l'argument i.
             */
             if (!anagrammes.get(i-1).isEmpty()){
                System.out.println(args[i] + ":");
                for (String s : anagrammes.get(i-1))
                    System.out.println(s);
             }
        }
    }

    private static ArrayList<ArrayList<String>> rechercheAnagrammes(Scanner sc, String[] args){
    //On fait une liste des arguments avec leurs caracteres tries et on cree les listes de resultats correspondants
        ArrayList<String> motsTries = new ArrayList<String>();
        ArrayList<ArrayList<String>> anagrammes = new ArrayList<ArrayList<String>>();
        for (int i = 1; i < args.length; ++i){
            motsTries.add(ordonnerMot(args[i]));
            anagrammes.add(new ArrayList<String>());
        }

        while(sc.hasNextLine()){
            String ligne = sc.nextLine();
            String ligneTriee = ordonnerMot(ligne);
            for (int i = 0; i < motsTries.size(); ++i){
                if(ligneTriee.equals(motsTries.get(i)))
                    anagrammes.get(i).add(ligne);
            }
        }
        return anagrammes;
    }

    public static void main (String[] args){
        if (args.length < 2){
            System.err.println("Syntaxe: java Anagrammes dictionnaire [mots]");
            return;
        }

        FileInputStream dictionnaire;
        try{
            dictionnaire = new FileInputStream(args[0]);
        } catch (FileNotFoundException e) {
            System.err.println("chemin du dictionnaire incorrect");
            return;
        }

        Scanner sc;
        try{
            sc = new Scanner(dictionnaire, "UTF-8");
        } catch (IllegalArgumentException e) {
            System.err.println("erreur lors de la creation du Scanner");
            return;
        }

        ArrayList<ArrayList<String>> anagrammes = rechercheAnagrammes(sc, args);
        trierResultats(anagrammes);
        afficherResultats(args, anagrammes);
    }
}
