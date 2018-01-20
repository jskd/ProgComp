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
   
       	
           
           int i, j;
           int y=0;
           char c;
          while (y != fact) {
          printf("%s\n", a);
          i=1;
          while(a[i] > a[i-1]) i++;
          j=0;
          while(a[j] < a[i])j++;         
      c=a[j];
      a[j]=a[i];
      a[i]=c;
i--;
for (j = 0; j < i; i--, j++) {
  c = a[i];
  a[i] = a[j];
  a[j] = c;
      }
y++;  
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
