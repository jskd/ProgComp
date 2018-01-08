#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define TAILLE_MAX 1000

int estdico(char * file,char * mot){
FILE* f=fopen(file,"r");
char chaine[TAILLE_MAX] = "";
while (fgets(chaine, TAILLE_MAX, f) != NULL){
	
	if(!(strcmp(chaine,mot)==0)){
		printf("%s et %s et %d\n", mot,chaine,strcmp(chaine,mot));
	}
}

}
void print_ana(char *head, char *str)
{
    
    static int k = 0;
    size_t sz = strlen(str);
    if(0 == sz) {
        return ;
    }
    if(1 == sz) {
        printf("%s%s\n", head, str);
    } else {
        size_t i;
        char headbuf[64], strbuf[64];
        for(i = 0; i < sz; ++i) {
            sprintf(headbuf, "%s%c", head, str[i]);
            
            strncpy(strbuf, str, i);
            strncpy(strbuf+i, str+i+1, sz-i);
           
            print_ana(headbuf, strbuf);
        }
    }
}
 
int main(int argc,char ** argv) 
{
    
    char head[64] = "";
    char str[64] = "chien";
    print_ana(head, str);    
 	estdico("dico","chien");
    return 0;
}
