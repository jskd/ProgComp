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

void process (FILE* f) {
  static char c;
  int count = 0;
  char* formula = NULL;
  while (fread (&c, 1, 1, f)) {
    if (c == '=') {
      formula = parse_formula (f);
      if (!formula) break;
      count++;
    }
  }
  printf ("%d\n", count);
}

int main (int argc, char** argv) {
  FILE* f = fopen (argv[1], "r");
  process (f);
  fclose (f);
  exit (EXIT_SUCCESS);
}
