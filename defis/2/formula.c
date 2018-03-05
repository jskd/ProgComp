#include <stdlib.h>
#include <stdio.h>

char read_char (FILE* f) {
  static char c;
  return c;
}

char* parse_formula (FILE* f) {
  char c;
  fread (&c, 1, 1, f);
  return "f";
}

#define BUFFER_SIZE (16384)

void process (FILE* f) {
  char c[BUFFER_SIZE];
  int count = 0;
  int read = BUFFER_SIZE;
  char* formula = NULL;
  while ((read = fread (&c, 1, BUFFER_SIZE, f))) {
    for (unsigned int i = 0; i < BUFFER_SIZE; i++) {
      if (c[i] == '=') {
	formula = parse_formula (f);
	if (!formula) break;
	count++;
      }
    }
    if (read < BUFFER_SIZE)
      break;
  }
  printf ("%d\n", count);
}

int main (int argc, char** argv) {
  FILE* f = fopen (argv[1], "r");
  process (f);
  fclose (f);
  exit (EXIT_SUCCESS);
}
