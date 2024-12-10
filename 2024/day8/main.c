#include "../utils.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_FREQUENCIES 26 + 26 + 10

typedef struct {
  int x;
  int y;
} point_t;

bool inBounds(point_t point, size_t width, size_t height) {
  return (0 <= point.x && point.x < (int)width) &&
         (0 <= point.y && point.y < (int)height);
}

bool contains(point_t *points, point_t point, size_t count) {
  for (size_t i = 0; i < count; i++) {
    if (points[i].x == point.x && points[i].y == point.y) {
      return true;
    }
  }
  return false;
}

int main(void) {
  char *input = read_file("input.txt");

  size_t width = 0;
  while (input[width] != '\n') {
    width++;
  }
  size_t height = strlen(input) / width - 1;

  char **grid = malloc(width * sizeof(char *));
  if (!grid) {
    perror("Failed to allocate memory");
    exit(EXIT_FAILURE);
  }
  for (size_t x = 0; x < width; x++) {
    grid[x] = malloc(height * sizeof(char));
    if (!grid[x]) {
      perror("Failed to allocate memory");
      exit(EXIT_FAILURE);
    }
  }

  for (size_t x = 0; x < width; x++) {
    for (size_t y = 0; y < height; y++) {
      size_t index = y * (width + 1) + x;
      grid[x][y] = input[index];
    }
  }

  char *frequencies = malloc(MAX_FREQUENCIES * sizeof(char));
  size_t frequencies_count = 0;
  if (!frequencies) {
    perror("Failed to allocate memory");
    exit(EXIT_FAILURE);
  }
  for (size_t i = 0; i < MAX_FREQUENCIES; i++) {
    frequencies[i] = '\0';
  }
  for (size_t x = 0; x < width; x++) {
    for (size_t y = 0; y < height; y++) {
      char frequency = grid[x][y];
      if (frequency == '.') {
        continue;
      }
      if (!strchr(frequencies, frequency)) {
        frequencies[frequencies_count++] = frequency;
      }
    }
  }

  point_t antennas[width * height];
  size_t antennas_count = 0;
  point_t antinodes[width * height];
  size_t antinodes_count = 0;
  size_t part1 = 0;
  for (size_t i = 0; i < frequencies_count; i++) {
    antennas_count = 0;
    char frequency = frequencies[i];
    for (size_t x = 0; x < width; x++) {
      for (size_t y = 0; y < height; y++) {
        if (grid[x][y] == frequency) {
          antennas[antennas_count++] = (point_t){x, y};
        }
      }
    }

    for (size_t j = 0; j < antennas_count; j++) {
      point_t antenna = antennas[j];
      for (size_t k = 0; k < antennas_count; k++) {
        if (j == k) {
          continue;
        }

        point_t other_antenna = antennas[k];

        int dx = other_antenna.x - antenna.x;
        int dy = other_antenna.y - antenna.y;

        point_t antinode =
            (point_t){other_antenna.x + dx, other_antenna.y + dy};
        if (inBounds(antinode, width, height) &&
            !contains(antinodes, antinode, antinodes_count)) {
          part1++;
          antinodes[antinodes_count++] = antinode;
        }
        antinode = (point_t){antenna.x - dx, antenna.y - dy};
        if (inBounds(antinode, width, height) &&
            !contains(antinodes, antinode, antinodes_count)) {
          part1++;
          antinodes[antinodes_count++] = antinode;
        }
      }
    }
  }

  antinodes_count = 0;
  size_t part2 = 0;
  for (size_t i = 0; i < frequencies_count; i++) {
    antennas_count = 0;
    char frequency = frequencies[i];
    for (size_t x = 0; x < width; x++) {
      for (size_t y = 0; y < height; y++) {
        if (grid[x][y] == frequency) {
          antennas[antennas_count++] = (point_t){x, y};
        }
      }
    }

    for (size_t j = 0; j < antennas_count; j++) {
      point_t antenna = antennas[j];
      for (size_t k = 0; k < antennas_count; k++) {
        if (j == k) {
          continue;
        }

        point_t other_antenna = antennas[k];

        int dx = other_antenna.x - antenna.x;
        int dy = other_antenna.y - antenna.y;
        point_t antinode;
        size_t loop = 0;
        do {
          antinode = (point_t){other_antenna.x + loop * dx,
                               other_antenna.y + loop * dy};
          if (!inBounds(antinode, width, height)) {
            break;
          }

          if (!contains(antinodes, antinode, antinodes_count)) {
            part2++;
            antinodes[antinodes_count++] = antinode;
          }
          loop++;
        } while (1);

        loop = 0;
        do {
          antinode = (point_t){antenna.x - loop * dx, antenna.y - loop * dy};
          if (!inBounds(antinode, width, height)) {
            break;
          }

          if (!contains(antinodes, antinode, antinodes_count)) {
            part2++;
            antinodes[antinodes_count++] = antinode;
          }
          loop++;
        } while (1);
      }
    }
  }

  printf("%zu\n", part1);
  printf("%zu\n", part2);

  free(frequencies);
  for (size_t x = 0; x < width; x++) {
    free(grid[x]);
  }
  free(grid);
  free(input);
}
