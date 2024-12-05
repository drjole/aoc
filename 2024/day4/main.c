#include "../utils.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BUFFER_SIZE 1024

bool inBounds(int x, int y, int width, int height) {
  return (0 <= x && x < width) && (0 <= y && y < height);
}

bool is(int x, int y, char c, char **grid, int width, int height) {
  return inBounds(x, y, width, height) && grid[x][y] == c;
}

int main(void) {
  char *input = read_file("input.txt");
  if (!input) {
    perror("Failed to read input.txt");
    return EXIT_FAILURE;
  }

  int width = 0;
  int height = 0;
  int i = 0;
  while (input[i] != '\n') {
    width++;
    i++;
  }
  height = (strlen(input)) / width - 1;

  char **grid = malloc(width * sizeof *grid);
  if (!grid) {
    perror("Failed to allocate memory");
    return 1;
  }
  for (int x = 0; x < width; x++) {
    grid[x] = malloc(height * sizeof *grid[x]);
    if (!grid[x]) {
      perror("Failed to allocate memory");
      return 1;
    }
    for (int y = 0; y < height; y++) {
      grid[x][y] = input[y * (width + 1) + x];
    }
  }

  free(input);

  int part1 = 0;
  int part2 = 0;
  for (int x = 0; x < width; x++) {
    for (int y = 0; y < height; y++) {
      for (int i = -1; i <= 1; i++) {
        for (int j = -1; j <= 1; j++) {
          if (i == 0 && j == 0) {
            continue;
          }

          if (is(x + i * 0, y + j * 0, 'X', grid, width, height) &&
              is(x + i * 1, y + j * 1, 'M', grid, width, height) &&
              is(x + i * 2, y + j * 2, 'A', grid, width, height) &&
              is(x + i * 3, y + j * 3, 'S', grid, width, height)) {
            part1++;
          }
        }
      }

      if (is(x, y, 'A', grid, width, height) &&
          ((is(x - 1, y - 1, 'M', grid, width, height) &&
            is(x + 1, y - 1, 'M', grid, width, height) &&
            is(x + 1, y + 1, 'S', grid, width, height) &&
            is(x - 1, y + 1, 'S', grid, width, height)) ||
           (is(x + 1, y - 1, 'M', grid, width, height) &&
            is(x + 1, y + 1, 'M', grid, width, height) &&
            is(x - 1, y + 1, 'S', grid, width, height) &&
            is(x - 1, y - 1, 'S', grid, width, height)) ||
           (is(x + 1, y + 1, 'M', grid, width, height) &&
            is(x - 1, y + 1, 'M', grid, width, height) &&
            is(x - 1, y - 1, 'S', grid, width, height) &&
            is(x + 1, y - 1, 'S', grid, width, height)) ||
           (is(x - 1, y + 1, 'M', grid, width, height) &&
            is(x - 1, y - 1, 'M', grid, width, height) &&
            is(x + 1, y - 1, 'S', grid, width, height) &&
            is(x + 1, y + 1, 'S', grid, width, height)))) {
        part2++;
      }
    }
  }

  for (int x = 0; x < width; x++) {
    free(grid[x]);
  }
  free(grid);

  printf("%d\n", part1);
  printf("%d\n", part2);

  return 0;
}
