#include "../utils.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

enum DIRECTION { NORTH, EAST, SOUTH, WEST };

typedef struct {
  int x;
  int y;
  enum DIRECTION d;
} vector_t;

vector_t rotate(vector_t position) {
  return (vector_t){position.x, position.y, (position.d + 1) % 4};
}

vector_t move(vector_t position) {
  switch (position.d) {
  case NORTH:
    return (vector_t){position.x, position.y - 1, NORTH};
  case EAST:
    return (vector_t){position.x + 1, position.y, EAST};
  case SOUTH:
    return (vector_t){position.x, position.y + 1, SOUTH};
  case WEST:
    return (vector_t){position.x - 1, position.y, WEST};
  }
  perror("Invalid direction");
  exit(EXIT_FAILURE);
}

void reset(bool ***grid, size_t width, size_t height) {
  for (size_t x = 0; x < width; x++) {
    for (size_t y = 0; y < height; y++) {
      for (size_t d = 0; d < 4; d++) {
        grid[x][y][d] = false;
      }
    }
  }
}

bool inBounds(vector_t position, size_t width, size_t height) {
  return (0 <= position.x && position.x < (int)width) &&
         (0 <= position.y && position.y < (int)height);
}

int main(void) {
  char *input = read_file("input.txt");
  if (!input) {
    perror("Failed to read input.txt");
    return EXIT_FAILURE;
  }

  size_t width = 0;
  size_t height = 0;
  while (input[width] != '\n') {
    width++;
  }
  height = strlen(input) / width - 1;

  bool ***blocked = malloc(width * sizeof(bool **));
  if (!blocked) {
    perror("Failed to allocate memory");
    return EXIT_FAILURE;
  }
  for (size_t x = 0; x < width; x++) {
    blocked[x] = malloc(height * sizeof(bool *));
    if (!blocked[x]) {
      perror("Failed to allocate memory");
      return EXIT_FAILURE;
    }
    for (size_t y = 0; y < height; y++) {
      blocked[x][y] = malloc(4 * sizeof(bool));
      if (!blocked[x][y]) {
        perror("Failed to allocate memory");
        return EXIT_FAILURE;
      }
    }
  }

  vector_t start;
  for (size_t y = 0; y < height; y++) {
    for (size_t x = 0; x < width; x++) {
      size_t index = y * (width + 1) + x;
      blocked[x][y][0] = (input[index] == '#');
      if (input[index] == '^') {
        start = (vector_t){x, y, NORTH};
      }
    }
  }

  bool ***visited = malloc(width * sizeof(bool **));
  if (!visited) {
    perror("Failed to allocate memory");
    return EXIT_FAILURE;
  }
  for (size_t x = 0; x < width; x++) {
    visited[x] = malloc(height * sizeof(bool *));
    if (!visited[x]) {
      perror("Failed to allocate memory");
      return EXIT_FAILURE;
    }
    for (size_t y = 0; y < height; y++) {
      visited[x][y] = malloc(4 * sizeof(bool));
      if (!visited[x][y]) {
        perror("Failed to allocate memory");
        return EXIT_FAILURE;
      }
    }
  }
  reset(visited, width, height);
  vector_t guard = start;
  size_t part1 = 0;
  while (inBounds(guard, width, height)) {
    if (!visited[guard.x][guard.y][0]) {
      visited[guard.x][guard.y][0] = true;
      part1++;
    }

    vector_t next = move(guard);
    if (!inBounds(next, width, height)) {
      break;
    }
    if (blocked[next.x][next.y][0]) {
      guard = rotate(guard);
    }
    guard = move(guard);
  }

  size_t part2 = 0;
  for (size_t x = 0; x < width; x++) {
    for (size_t y = 0; y < height; y++) {
      if (blocked[x][y][0] || ((int)x == start.x && (int)y == start.y)) {
        continue;
      }

      blocked[x][y][0] = true;
      guard = start;
      reset(visited, width, height);

      while (inBounds(guard, width, height)) {
        if (visited[guard.x][guard.y][guard.d]) {
          part2++;
          break;
        }

        visited[guard.x][guard.y][guard.d] = true;

        vector_t next = move(guard);
        if (!inBounds(next, width, height)) {
          break;
        }
        if (blocked[next.x][next.y][0]) {
          guard = rotate(guard);
        } else {
          guard = move(guard);
        }
      }

      blocked[x][y][0] = false;
    }
  }

  printf("%zu\n", part1);
  printf("%zu\n", part2);

  for (size_t x = 0; x < width; x++) {
    for (size_t y = 0; y < height; y++) {
      free(blocked[x][y]);
    }
    free(blocked[x]);
  }
  free(blocked);
  for (size_t x = 0; x < width; x++) {
    for (size_t y = 0; y < height; y++) {
      free(visited[x][y]);
    }
    free(visited[x]);
  }
  free(visited);
  free(input);
}
