#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define MAGIC_HASH 53 //should be prime
#define HASHTABLE 4096
#define INIT_LINE 256
#define MAX_CHAR 256
#define INIT_WORDS 64

typedef struct pair{
	char *word;
	struct pair *next;
} entry;

typedef struct{
	int len, chars[MAX_CHAR];
}anagram;

static entry *hashtable[HASHTABLE];

//"damned mingw-w64" part
char *my_strndup(const char *src, size_t dest_len){
	char *res = malloc(sizeof(char) * (dest_len + 1));
	if (res == NULL)
		return NULL;
	int len = strlen(src), i;
	for(i = 0; i < dest_len + 1; i++)
		res[i] = i < len ? src[i] : '\0';
	return res;
}

//my impl. of getline; ATTN, doesn't include newline char.
//return: -1 allocation error, 0 EOF, 1 newline
int my_getline(char **lineptr, size_t *n, FILE *stream){
	if(!*n){
		*n = INIT_LINE;
		if ((*lineptr = malloc(sizeof(char) * *n)) == NULL)
			return -1;		
	}
	unsigned ptr = 0; //line no longer than int, errors otherwise
	char c;
	while(1){
		c = fgetc(stream);
		if (c == EOF || c == '\n'){
			(*lineptr)[ptr] = '\0';
			return c == EOF ? 0 : 1;
		}
		(*lineptr)[ptr++] = c;
		if (ptr == *n){
			*n += *n/2;
			if ((*lineptr = realloc(*lineptr,sizeof(char) * *n)) == NULL)
				return -1;
		}
	}
}


/*Dictionary part, inspired by (and copied from):
https://stackoverflow.com/questions/4384359/quick-way-to-implement-dictionary-in-c

*/

entry *new_entry(const char *src){
	entry *ret = malloc(sizeof(entry));
	char *s = my_strndup(src, strlen(src));
	if (ret == NULL || s == NULL)
		return NULL;
	ret->word = s;
	ret->next = NULL;
	return ret;
}

unsigned long long hash(const char *s){
    unsigned long long hashval;
    for (hashval = 0; *s != '\0'; s++)
      hashval = *s + MAGIC_HASH * hashval;
    return hashval % HASHTABLE;
}

int check_word(const char *s){
    entry *e;
    for (e = hashtable[hash(s)]; e != NULL; e = e->next)
        if (strcmp(s, e->word) == 0)
          return 1;
    return 0; 
}

int add_word(const char *s){
    entry *e;
	unsigned long long hashval = hash(s);
    if ((e = hashtable[hashval]) == NULL){
		if((hashtable[hashval] = new_entry(s)) == NULL)
			return -1;
	}
	else
		while(1){
			if (!strcmp(e->word, s))
				break;
			if (e->next == NULL){
				if ((e->next = new_entry(s)) == NULL)
					return -1;
				break;
			}
			e = e->next;
		}
	return 1;
}

int file_to_dict(char *filename){
	FILE* f = fopen(filename, "r");
	if (f == NULL)
		return -1;
	int ret;
	char *buf = NULL;
	size_t n = 0;
	while(1){
		ret = my_getline(&buf, &n, f);
		if(ret == -1){
			if (buf != NULL)
				free(buf);
			return -1;
		}
		if (strlen(buf) > 0)
			if (add_word(buf) == -1){
				free(buf);
				return -1;
			}
		if (ret == 0){
			free(buf);
			return 0;
		}
	}
}

void print_dict(){
	int i;
	entry *e;
	for(i = 0; i < HASHTABLE; i++){
		e = hashtable[i];
			while(1){
				if(e == NULL)
					break;
				printf("%s\n", e->word);
				e=e->next;
			}
	}
}

void free_entry(entry *e){
	if (e->next != NULL)
		free_entry(e->next);
	free(e->word);
	free(e);
	return;
}

void free_dict(){
	int i;
	for(i = 0; i < HASHTABLE; i++)
		if(hashtable[i] != NULL)
			free_entry(hashtable[i]);
}

/*Anagram part. Only 1-byte characters are allowed ;-P*/

anagram *new_anagram(const char *s){
	anagram *ret = malloc(sizeof(anagram));
	if (ret == NULL)
		return NULL;
	ret->len = 0;
	int i;
	for(i = 0; i< MAX_CHAR; i++)
		ret->chars[i] = 0;
	i = 0;
	while(1){
		if (s[i] == '\0')
			return ret;
		(ret->chars[s[i++]])++;
		(ret->len)++;
	}
}

void all_anagrams(anagram *a, char* buffer, int ptr){
	int i;
	(a->len)--;
	for(i = 0; i < MAX_CHAR; i++)
		if (a->chars[i] > 0){
			buffer[ptr] = (char)i;
			if(a->len){
				(a->chars[i])--;
				all_anagrams(a, buffer, ptr + 1);
				(a->chars[i])++;
			}
			else
				if (check_word(buffer))
					printf("%s\n", buffer);
		}
	(a->len)++;
	return;
}

int do_word(const char *s){
	anagram *a = new_anagram(s);
	char *buffer = malloc(sizeof(char) * (strlen(s) + 1));
	if(a == NULL || buffer == NULL)
		return -1;
	buffer[a->len] = '\0';
	all_anagrams(a, buffer, 0);
	free(a);
	free(buffer);
	return 0;
}

int main(int argc, char *argv[]){
	if (argc < 2) {
		printf("Args!\n");
		return -1;
	}
	if ((file_to_dict(argv[1])) == -1){
		printf("Fail during reading file, sorry\n");
		return -1;
	}
	int i;
	for(i = 2; i < argc; i++)
		if(do_word(argv[i]) == -1){
			printf("Malloc problems during working with word %s!\n", argv[i]);
			return -1;
		}
	//print_dict();
	free_dict();
	return 0;
}

