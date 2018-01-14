import java.io.BufferedReader;
import java.io.FileReader;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.util.*;
class Anagramme {
public static void main(String[] argv) throws IOException{
    for(int i=1;i<argv.length;i++){	//pour tous les parametres
    	String ligne;
    	HashSet<String> permutations = new HashSet<String>(); 
    	String mot = argv[i];
    	anagrammesPossibles(0,mot,permutations);//toutes les permutations possibles
        BufferedReader lecteur = null;
    	try{
    		lecteur = new BufferedReader(new FileReader(argv[0]));
        }
    	catch(FileNotFoundException exc)
        {
    		System.out.println("Erreur d'ouverture");
        }
    	while ((ligne = lecteur.readLine()) != null){ //tant qu'il y a des mots dans le fichier
			for(Object perm : permutations) {
			    String element = (String) perm;
			    if(element.equals(ligne)){	//si ligne est egale a un element de permutation
			    	System.out.println(mot + "----"+ ligne);
			    }
			}
		}
        lecteur.close();
    }
  }

  static void anagrammesPossibles(int pivot,String mot,HashSet<String> permutations) {
	    char[] motEnTab=mot.toCharArray();//tableau des caracteres du mot
	    if(pivot==mot.length())	//si le pivot de permutation est egale a la longueur du mot
	    	permutations.add(mot);	//ajout de la permutaion au tableau de permutation
	    for(int i=pivot;i<mot.length();i++) {	//permutation des caracteres commencant du pivot
	        char c=motEnTab[i];
	        motEnTab[i]=motEnTab[pivot];
	        motEnTab[pivot]=c;
	        anagrammesPossibles(pivot+1,new String(motEnTab),permutations);//appel recursif de la fonction
	    }   
	}
}
