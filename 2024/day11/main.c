#include "../utils.h"
#include <assert.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Pair {
  unsigned long first;
  unsigned long second;
};

struct Node {
  struct Node *prev;
  struct Node *next;
  struct Pair pair;
};

void debug(struct Node *head) {
  while (head != NULL) {
    if (head->pair.second != 0) {
      printf("%lu: %lu |", head->pair.first, head->pair.second);
    }
    head = head->next;
  }
  printf("\n");
}

unsigned long numDigits(unsigned long number) {
  unsigned long digits = 0;

  if (number == 0) {
    digits = 1;
  } else {
    while (number != 0) {
      number /= 10;
      digits++;
    }
  }

  return digits;
}

void splitNumber(unsigned long number, unsigned long *first,
                 unsigned long *second, unsigned long n) {
  *first = number / pow(10, (int)(log10(number) - n + 1));
  *second = number % (int)pow(10, (int)(log10(number) - n + 1));
}

struct Node *increment(struct Node *head, unsigned long key,
                       unsigned long amount) {
  struct Node *node = head;
  struct Node *last = NULL;
  while (node != NULL) {
    if (node->pair.first == key) {
      node->pair.second += amount;
      return head;
    }
    if (node->next == NULL) {
      last = node;
    }
    node = node->next;
  }

  struct Node *new_node = malloc(sizeof(struct Node));
  if (!new_node) {
    perror("Failed to allocate memory");
    return NULL;
  }

  new_node->prev = last;
  new_node->next = NULL;
  new_node->pair.first = key;
  new_node->pair.second = amount;

  if (last != NULL) {
    last->next = new_node;
  }

  return head == NULL ? new_node : head;
}

void decrement(struct Node *head, unsigned long key, unsigned long amount) {
  struct Node *node = head;
  while (node != NULL) {
    if (node->pair.first == key) {
      node->pair.second -= amount;
      return;
    }
    node = node->next;
  }
}

struct Node *copy(struct Node *head) {
  struct Node *new_head = NULL;
  struct Node *node = head;
  while (node != NULL) {
    new_head = increment(new_head, node->pair.first, node->pair.second);
    node = node->next;
  }
  return new_head;
}

unsigned long length(struct Node *head) {
  unsigned long length = 0;
  struct Node *node = head;
  while (node != NULL) {
    length += node->pair.second;
    node = node->next;
  }
  return length;
}

void clear(struct Node *head) {
  struct Node *node = head;
  while (node != NULL) {
    struct Node *next = node->next;
    free(node);
    node = next;
  }
}

int main(void) {
  char *input = read_file("input.txt");
  if (!input) {
    perror("Failed to read input.txt");
    free(input);
    return EXIT_FAILURE;
  }

  struct Node *head = NULL;

  char *token = strtok(input, " ");
  while (token != NULL) {
    unsigned long value = atoll(token);
    head = increment(head, value, 1);
    token = strtok(NULL, " ");
  }

  for (unsigned long i = 1; i <= 75; i++) {
    struct Node *new_head = copy(head);
    struct Node *node = head;
    while (node != NULL) {
      struct Pair pair = node->pair;

      if (pair.second == 0) {
        goto next;
      }

      if (pair.first == 0) {
        decrement(new_head, 0, pair.second);
        increment(new_head, 1, pair.second);
      } else {
        unsigned long digits = numDigits(pair.first);
        if (digits % 2 == 0) {
          unsigned long first = 0;
          unsigned long second = 0;
          splitNumber(pair.first, &first, &second, digits / 2);
          decrement(new_head, pair.first, pair.second);
          increment(new_head, first, pair.second);
          increment(new_head, second, pair.second);
        } else {
          decrement(new_head, pair.first, pair.second);
          increment(new_head, pair.first * 2024, pair.second);
        }
      }

    next:
      node = node->next;
    }

    struct Node *tmp = head;
    head = new_head;
    clear(tmp);

    if (i == 25) {
      printf("%lu\n", length(new_head));
    }

    if (i == 75) {
      printf("%lu\n", length(new_head));
    }
  }

  clear(head);
  free(input);
}
