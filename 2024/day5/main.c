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

rule_t *rules;
int rules_count = 0;

bool before(int page, int otherPage) {
  for (int i = 0; i < rules_count; i++) {
    if (rules[i].before == page && rules[i].after == otherPage) {
      return true;
    }
  }

  return false;
}

bool correctOrder(update_t update) {
  for (size_t i = 0; i < update.len - 1; i++) {
    if (!before(update.pages[i], update.pages[i + 1])) {
      return false;
    }
  }
  return true;
}

int order(const void *a, const void *b) {
  int page_a = *((int *)a);
  int page_b = *((int *)b);
  return (int)before(page_a, page_b);
}

int main(void) {
  char *input = read_file("input.txt");

  rules = malloc(2048 * sizeof(rule_t));
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
  while (token != NULL) {
    int before;
    int after;
    if (sscanf(token, "%d|%d", &before, &after) != 2) {
      break;
    }

    rule_t rule = {before, after};
    rules[rules_count++] = rule;

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

  int part1 = 0;
  int part2 = 0;
  for (int i = 0; i < updates_index; i++) {
    if (correctOrder(updates[i])) {
      part1 += updates[i].pages[updates[i].len / 2];
    } else {
      qsort(updates[i].pages, updates[i].len, sizeof(int), order);
      part2 += updates[i].pages[updates[i].len / 2];
    }
  }

  printf("%d\n", part1);
  printf("%d\n", part2);

  free(rules);
  for (int u = 0; u < updates_index; u++) {
    free(updates[u].pages);
  }
  free(updates);
  free(input);

  return EXIT_SUCCESS;
}
