#include "../utils.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  int before;
  int after;
} rule_t;

typedef struct {
  int *pages;
  size_t len;
} update_t;

int main() {
  char *input = read_file("input.txt");

  rule_t *rules = malloc(2048 * sizeof(rule_t));
  if (!rules) {
    perror("Failed to allocate memory");
    return EXIT_FAILURE;
  }

  update_t *updates = malloc(2048 * sizeof(update_t));
  if (!updates) {
    perror("Failed to allocate memory");
    return EXIT_FAILURE;
  }

  char *input_saveptr;
  char *token = strtok_r(input, "\n", &input_saveptr);
  int rules_index = 0;
  while (token != NULL) {
    int before;
    int after;
    if (sscanf(token, "%d|%d", &before, &after) != 2) {
      break;
    }

    rule_t rule = {before, after};
    rules[rules_index++] = rule;

    token = strtok_r(NULL, "\n", &input_saveptr);
  }

  int updates_index = 0;
  while (token != NULL) {
    int *pages = malloc(2048 * sizeof(int));
    if (!pages) {
      perror("Failed to allocate memory");
      break;
    }

    char *line_saveptr;
    char *line_token = strtok_r(token, ",", &line_saveptr);
    int p = 0;
    while (line_token != NULL) {
      int page = atoi(line_token);
      pages[p++] = page;
      line_token = strtok_r(NULL, ",", &line_saveptr);
    }

    update_t update = {pages, p};
    updates[updates_index++] = update;

    token = strtok_r(NULL, "\n", &input_saveptr);
  }

  free(rules);
  for (int u = 0; u < updates_index; u++) {
    free(updates[u].pages);
  }
  free(updates);
  free(input);

  return EXIT_SUCCESS;
}
