#include "../utils.h"
#include <complex.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

bool inBounds(int x, int y, size_t width, size_t height) {
  return (0 <= x && x < (int)width) && (0 <= y && y < (int)height);
}

size_t pathsToHighest(int x, int y, size_t value, unsigned int **grid,
                      size_t width, size_t height, int **memo) {
  if (!inBounds(x, y, width, height)) {
    return 0;
  }

  if (memo[x][y] || grid[x][y] != value) {
    return 0;
  }

  memo[x][y]++;

  if (value == 0) {
    return 1;
  }

  return pathsToHighest(x + 1, y, value - 1, grid, width, height, memo) +
         pathsToHighest(x - 1, y, value - 1, grid, width, height, memo) +
         pathsToHighest(x, y + 1, value - 1, grid, width, height, memo) +
         pathsToHighest(x, y - 1, value - 1, grid, width, height, memo);
}

size_t allPathsToHighest(int x, int y, size_t value, unsigned int **grid,
                         size_t width, size_t height, int **memo) {
  if (!inBounds(x, y, width, height)) {
    return 0;
  }

  if (grid[x][y] != value) {
    return 0;
  }

  memo[x][y]++;

  if (value == 0) {
    return 1;
  }

  return (int)allPathsToHighest(x + 1, y, value - 1, grid, width, height,
                                memo) +
         (int)allPathsToHighest(x - 1, y, value - 1, grid, width, height,
                                memo) +
         (int)allPathsToHighest(x, y + 1, value - 1, grid, width, height,
                                memo) +
         (int)allPathsToHighest(x, y - 1, value - 1, grid, width, height, memo);
}

int main(void) {
  char *input = read_file("input.txt");
  if (!input) {
    perror("Failed to read input.txt");
    return EXIT_FAILURE;
  }

  size_t width = 0;
  while (input[width] != '\n') {
    width++;
  }
  size_t height = strlen(input) / width - 1;

  unsigned int **grid = malloc(width * sizeof(unsigned int *));
  if (!grid) {
    perror("Failed to allocate memory");
    free(input);
    return EXIT_FAILURE;
  }
  for (size_t x = 0; x < width; x++) {
    grid[x] = malloc(height * sizeof(unsigned int));
    if (!grid[x]) {
      perror("Failed to allocate memory");
      free(grid);
      free(input);
      exit(EXIT_FAILURE);
    }
  }

  for (size_t y = 0; y < height; y++) {
    for (size_t x = 0; x < width; x++) {
      size_t index = y * (width + 1) + x;
      grid[x][y] = input[index] - '0';
    }
  }

  int **memo = malloc(width * sizeof(int *));
  if (!memo) {
    perror("Failed to allocate memory");
    for (size_t x = 0; x < width; x++) {
      free(grid[x]);
    }
    free(grid);
    free(input);
    return EXIT_FAILURE;
  }
  for (size_t x = 0; x < width; x++) {
    memo[x] = malloc(height * sizeof(int));
    if (!memo[x]) {
      perror("Failed to allocate memory");
      free(memo);
      for (size_t x = 0; x < width; x++) {
        free(grid[x]);
      }
      free(grid);
      free(input);
    }
  }

  for (size_t x = 0; x < width; x++) {
    for (size_t y = 0; y < height; y++) {
      memo[x][y] = 0;
    }
  }

  size_t part1 = 0;
  size_t part2 = 0;
  for (size_t x = 0; x < width; x++) {
    for (size_t y = 0; y < height; y++) {
      if (grid[x][y] == 9) {
        part1 += pathsToHighest((int)x, (int)y, 9, grid, width, height, memo);
        for (size_t x2 = 0; x2 < width; x2++) {
          for (size_t y2 = 0; y2 < width; y2++) {
            memo[x2][y2] = 0;
          }
        }

        part2 +=
            allPathsToHighest((int)x, (int)y, 9, grid, width, height, memo);
        for (size_t x2 = 0; x2 < width; x2++) {
          for (size_t y2 = 0; y2 < width; y2++) {
            memo[x2][y2] = 0;
          }
        }
      }
    }
  }

  printf("%zu\n", part1);
  printf("%zu\n", part2);

  for (size_t x = 0; x < width; x++) {
    free(memo[x]);
    free(grid[x]);
  }
  free(memo);
  free(grid);
  free(input);
}
