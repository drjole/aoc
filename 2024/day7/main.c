#include "../utils.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

unsigned long long long_long_pow(unsigned long long base,
                                 unsigned long long exponent) {
  unsigned long long result = 1;
  for (unsigned long long i = 0; i < exponent; i++) {
    result *= base;
  }
  return result;
}

int main(void) {
  char *input = read_file("input.txt");

  unsigned long long part1 = 0;
  unsigned long long part2 = 0;

  size_t values[32];
  size_t values_count = 0;

  char *line_saveptr;
  char *line = strtok_r(input, "\n", &line_saveptr);
  while (line != NULL) {
    unsigned long long expected_result;
    sscanf(line, "%llu: ", &expected_result);
    line = strchr(line, ' ') + 1;

    char *token_saveptr;
    char *token = strtok_r(line, " ", &token_saveptr);
    values_count = 0;
    while (token != NULL) {
      size_t value = atoi(token);
      values[values_count++] = value;
      token = strtok_r(NULL, " ", &token_saveptr);
    }

    for (size_t i = 0; i < long_long_pow(2, values_count - 1); i++) {
      unsigned long long result = values[0];
      // printf("%llu = %llu", expected_result, result_1);

      for (size_t j = 0; j < values_count - 1; j++) {
        size_t operator=(i >> j) & 1;

        if (operator== 0) {
          // printf(" + %llu", values[j + 1]);
          result += values[j + 1];
        } else if (operator== 1) {
          // printf(" * %llu", values[j + 1]);
          result *= values[j + 1];
        }
      }

      if (result == expected_result) {
        part1 += expected_result;
        break;
      }
    }

    for (size_t i = 0; i < long_long_pow(3, values_count - 1); i++) {
      unsigned long long result = values[0];
      // printf("%llu = %llu", expected_result, result_1);

      for (size_t j = 0; j < values_count - 1; j++) {
        size_t num = i;
        for (size_t k = 0; k < j; k++) {
          num /= 3;
        }
        size_t operator= num % 3;

        if (operator== 0) {
          // printf(" + %llu", values[j + 1]);
          result += values[j + 1];
        } else if (operator== 1) {
          // printf(" * %llu", values[j + 1]);
          result *= values[j + 1];
        } else if (operator== 2) {
          // printf(" || %llu", values[j + 1]);
          size_t digits = 0;
          size_t temp = values[j + 1];
          while (temp > 0) {
            temp /= 10;
            digits++;
          }
          result *= long_long_pow(10, digits);
          result += (unsigned long long)values[j + 1];
        }
      }

      if (result == expected_result) {
        part2 += expected_result;
        break;
      }
    }

    line = strtok_r(NULL, "\n", &line_saveptr);
  }

  printf("%llu\n", part1);
  printf("%llu\n", part2);

  free(input);
}
