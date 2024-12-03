#include <memory.h>
#include <regex.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

#include "../utils.h"

int main(void) {
  char *input = read_file("input.txt");
  if (!input) {
    perror("Failed to read input.txt");
    return EXIT_FAILURE;
  }

  char *multiplicationRegexString =
      "mul\\([[:digit:]]{1,3},[[:digit:]]{1,3}\\)";
  char *doRegexString = "do\\(\\)";
  char *dontRegexString = "don't\\(\\)";

  regex_t multiplicationRegex;
  regex_t doRegex;
  regex_t dontRegex;

  regmatch_t multiplicationMatches[1];
  regmatch_t doDontMatches[1];

  if (regcomp(&multiplicationRegex, multiplicationRegexString, REG_EXTENDED)) {
    perror("Failed to compile regex");
    return EXIT_FAILURE;
  }

  if (regcomp(&doRegex, doRegexString, REG_EXTENDED)) {
    perror("Failed to compile regex");
    return EXIT_FAILURE;
  }

  if (regcomp(&dontRegex, dontRegexString, REG_EXTENDED)) {
    perror("Failed to compile regex");
    return EXIT_FAILURE;
  }

  long part1 = 0;
  long part2 = 0;
  char *p = input;
  int enabled = 1;

  while (1) {
    if (regexec(&multiplicationRegex, p, 1, multiplicationMatches, 0)) {
      break;
    }
    if (!regexec(&doRegex, p, 1, doDontMatches, 0) &&
        doDontMatches[0].rm_so == 0) {
      enabled = 1;
    }
    if (!regexec(&dontRegex, p, 1, doDontMatches, 0) &&
        doDontMatches[0].rm_so == 0) {
      enabled = 0;
    }

    if (multiplicationMatches[0].rm_so == 0) {
      int a;
      int b;
      sscanf(p, "mul(%d,%d)", &a, &b);
      part1 += a * b;
      if (enabled) {
        part2 += a * b;
      }
    }
    p++;
  }

  regfree(&multiplicationRegex);

  printf("%ld\n", part1);
  printf("%ld\n", part2);

  free(input);

  return EXIT_SUCCESS;
}
