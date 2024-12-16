#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *read_file(char *path) {
  FILE *file;
  file = fopen(path, "r");
  if (!file) {
    perror("Failed to read input.txt");
    return NULL;
  }

  size_t buffer_size = 128;
  char *buffer = malloc(buffer_size);
  if (!buffer) {
    perror("Error allocating memory");
    fclose(file);
    return NULL;
  }

  size_t length = 0;
  int c;

  while ((c = fgetc(file)) != EOF) {
    if (length + 1 >= buffer_size) {
      buffer_size *= 2;
      char *new_buffer = realloc(buffer, buffer_size);
      if (!new_buffer) {
        perror("Error allocating memory");
        fclose(file);
        return NULL;
      }
      buffer = new_buffer;
    }
    buffer[length++] = (char)c;
  }
  buffer[length] = '\0';

  fclose(file);

  return buffer;
}

size_t grid_width(char *s) {
  size_t count = 0;
  while (s[count] != '\0') {
    if (s[count] == '\n') {
      break;
    }
    count++;
  }

  return count;
}

size_t grid_height(char *s) {
  size_t count = 0;
  while (*s) {
    if (*s == '\n') {
      count++;
    }
    s++;
  }

  return count;
}

char **make_grid(char *s) {
  size_t width = grid_width(s);
  size_t height = grid_height(s);

  char **grid = malloc(width * sizeof(char *));
  if (!grid) {
    perror("Failed to allocate memory");
    return NULL;
  }
  for (size_t x = 0; x < width; x++) {
    grid[x] = malloc(height * sizeof(char));
    if (!grid[x]) {
      perror("Failed to allocate memory");
      return NULL;
    }
  }

  for (size_t x = 0; x < width; x++) {
    for (size_t y = 0; y < height; y++) {
      size_t index = y * (width + 1) + x;
      grid[x][y] = s[index];
    }
  }

  return grid;
}

void print_grid(char **grid, size_t width, size_t height) {
  for (size_t y = 0; y < height; y++) {
    for (size_t x = 0; x < width; x++) {
      printf("%c", grid[x][y]);
    }
    printf("\n");
  }
}

void free_grid(char **grid, size_t width) {
  for (size_t x = 0; x < width; x++) {
    free(grid[x]);
  }
  free(grid);
}

bool in_bounds(int x, int y, int width, int height) {
  return 0 <= x && x < width && 0 <= y && y < height;
}

bool is_neighbor(int ax, int ay, int bx, int by) {
  return (ax == bx && abs(ay - by) == 1) || (ay == by && abs(ax - bx) == 1);
}
