import java.io.BufferedReader;
import java.io.FileReader;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.lang.*;
import java.util.*;
import java.util.*;
import java.io.*;
class Anagramme {
  public static void main(String[] argv) throws IOException
  {
  
    BufferedReader lecteur = null;
    String ligne;

    try
      {
									lecteur = new BufferedReader(new FileReader(argv[0]));
      }
    catch(FileNotFoundException exc)
      {
									System.out.println("Erreur d'ouverture");
      }
    
    while ((ligne = lecteur.readLine()) != null){
      			    for(int i=1;i<argv.length;i++){
																String mot = argv[i];
																println(existe(mot,ligne));
													}
    }
    lecteur.close();
  }
  public String existe(String mot, String mot2){
  						int compteur;
  						if(mot.length==mot2.length){
  						for (int i=0; i<mot.length; i++){
         			for (int j=0; j<mot2.length; j++){
         					if(mot[i]==mot[j]){
         								compteur++;
         					}
         			}
  						}
  						if(compteur==mot.length){
  										return mot;
  						}
  						}
  }
}
