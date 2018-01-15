#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define TAILLE_MAX 1000

int estdico(char * file,char * mot){
FILE* f=fopen(file,"r");
char chaine[TAILLE_MAX];

while (fgets(chaine, TAILLE_MAX, f) != NULL){
	
	if(!(strcmp(chaine,mot)==0)){
		return 1;
	}
	
		
}
return 0;

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
	printf("-------");
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
 
int main(int argc,char * argv[]) 
{	
    int i;
    for(i=2;i<argc;i++){
    if(estdico(argv[1],strcat(argv[i],"\n"))){
    char head[64] = "";
    char str[64] = "foo";
    print_ana(head, str);    }}
 	
    return 0;
}
