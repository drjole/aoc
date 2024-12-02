#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

bool all_ascending(int *numbers, int length, int skip) {
  for (int i = 0; i < length - 1; i++) {
    if (i == skip) {
      continue;
    }

    int b = i + 1;
    if (b == skip) {
      b += 1;
    }
    if (b >= length) {
      break;
    }

    if (numbers[b] <= numbers[i]) {
      return false;
    }
  }

  return true;
}

bool all_descending(int *numbers, int length, int skip) {
  for (int i = 0; i < length - 1; i++) {
    if (i == skip) {
      continue;
    }

    int b = i + 1;
    if (b == skip) {
      b += 1;
    }
    if (b >= length) {
      break;
    }

    if (numbers[b] >= numbers[i]) {
      return false;
    }
  }

  return true;
}

bool safe_distances(int *numbers, int length, int skip) {
  for (int i = 0; i < length - 1; i++) {
    if (i == skip) {
      continue;
    }

    int b = i + 1;
    if (b == skip) {
      b += 1;
    }
    if (b >= length) {
      break;
    }

    int diff = abs(numbers[b] - numbers[i]);
    if (diff < 1 || diff > 3) {
      return false;
    }
  }

  return true;
}

bool safe(int *numbers, int length, int skip) {
  return ((all_ascending(numbers, length, skip) ||
           all_descending(numbers, length, skip)) &&
          safe_distances(numbers, length, skip));
}

int main(void) {
  char buffer[100];
  FILE *file = fopen("input.txt", "r");

  int part1 = 0;
  int part2 = 0;

  while (fgets(buffer, 100, file)) {
    char *token = strtok(buffer, " ");
    int numbers[100];
    int index = 0;
    while (token != NULL) {
      numbers[index] = atoi(token);
      index++;
      token = strtok(NULL, " ");
    }

    if (safe(numbers, index, -1)) {
      part1++;
      part2++;
    } else {
      for (int i = 0; i < index; i++) {
        if (safe(numbers, index, i)) {
          part2++;
          break;
        }
      }
    }
  }

  fclose(file);

  printf("%d\n", part1);
  printf("%d\n", part2);

  return 0;
}
