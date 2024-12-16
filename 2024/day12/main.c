#include "../utils.h"
#include <stdio.h>
#include <stdlib.h>

int directions[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

bool **make_set(size_t width, size_t height) {
  bool **set = malloc(width * sizeof(bool *));
  if (!set) {
    return NULL;
  }
  for (size_t x = 0; x < width; x++) {
    set[x] = malloc(height * sizeof(bool));
    if (!set[x]) {
      for (size_t i = 0; i < x; i++) {
        free(set[i]);
      }
      free(set);
      return NULL;
    }
  }
  return set;
}

void init_set(bool **set, size_t width, size_t height) {
  for (size_t x = 0; x < width; x++) {
    for (size_t y = 0; y < height; y++) {
      set[x][y] = false;
    }
  }
}

void free_set(bool **set, size_t width) {
  for (size_t x = 0; x < width; x++) {
    free(set[x]);
  }
  free(set);
}

size_t count_corners(char **grid, int x, int y, char plant, size_t width,
                     size_t height) {
  size_t corners = 0;

  int corner_offsets[4][6] = {{0, -1, -1, 0, -1, -1},
                              {0, -1, 1, 0, 1, -1},
                              {0, 1, -1, 0, -1, 1},
                              {0, 1, 1, 0, 1, 1}};

  for (int i = 0; i < 4; i++) {
    int nx1 = x + corner_offsets[i][0];
    int ny1 = y + corner_offsets[i][1];
    int nx2 = x + corner_offsets[i][2];
    int ny2 = y + corner_offsets[i][3];
    int ndx = x + corner_offsets[i][4];
    int ndy = y + corner_offsets[i][5];

    char a1 = (nx1 >= 0 && ny1 >= 0 && nx1 < (int)width && ny1 < (int)height)
                  ? grid[nx1][ny1]
                  : '\0';
    char a2 = (nx2 >= 0 && ny2 >= 0 && nx2 < (int)width && ny2 < (int)height)
                  ? grid[nx2][ny2]
                  : '\0';
    char d = (ndx >= 0 && ndy >= 0 && ndx < (int)width && ndy < (int)height)
                 ? grid[ndx][ndy]
                 : '\0';

    if ((a1 == plant && a2 == plant && d != plant) ||
        (a1 != plant && a2 != plant)) {
      corners++;
    }
  }

  return corners;
}

void flood_fill(char **grid, int x, int y, char plant, bool **visited,
                size_t *area, size_t *perimeter, size_t *corners, size_t width,
                size_t height) {
  *area = 0;

  size_t stack[width * height][2];
  int top = -1;

  stack[++top][0] = x;
  stack[top][1] = y;
  visited[x][y] = true;

  while (top >= 0) {
    int cx = stack[top][0];
    int cy = stack[top--][1];

    (*area)++;
    *corners += count_corners(grid, cx, cy, plant, width, height);

    for (int i = 0; i < 4; i++) {
      int nx = cx + directions[i][0];
      int ny = cy + directions[i][1];

      if (nx < 0 || ny < 0 || nx >= (int)width || ny >= (int)height ||
          grid[nx][ny] != plant) {
        (*perimeter)++;
        continue;
      }

      if (!visited[nx][ny]) {
        stack[++top][0] = nx;
        stack[top][1] = ny;
        visited[nx][ny] = true;
      }
    }
  }
}

int main(void) {
  char *input = read_file("input.txt");
  if (!input) {
    perror("Failed to read input.txt");
    return EXIT_FAILURE;
  }

  size_t width = grid_width(input);
  size_t height = grid_height(input);

  char **grid = make_grid(input);
  if (!grid) {
    perror("Failed to make grid");
    free(input);
    return EXIT_FAILURE;
  }

  bool **visited = make_set(width, height);
  if (!visited) {
    perror("Failed to make set");
    free(input);
    return EXIT_FAILURE;
  }

  bool **horizontal_sides = make_set(width, height + 1);
  if (!horizontal_sides) {
    perror("Failed to make set");
    free_set(visited, width);
    free(input);
    return EXIT_FAILURE;
  }

  bool **vertical_sides = make_set(width + 1, height);
  if (!vertical_sides) {
    perror("Failed to make set");
    free_set(horizontal_sides, width);
    free_set(visited, width);
    free(input);
  }

  init_set(visited, width, height);
  size_t part1 = 0;
  size_t part2 = 0;
  for (size_t x = 0; x < width; x++) {
    for (size_t y = 0; y < height; y++) {
      if (!visited[x][y]) {
        size_t area = 0;
        size_t perimeter = 0;
        size_t corners = 0;

        init_set(horizontal_sides, width, height + 1);
        init_set(vertical_sides, width + 1, height);

        flood_fill(grid, x, y, grid[x][y], visited, &area, &perimeter, &corners,
                   width, height);

        part1 += area * perimeter;
        part2 += area * corners;
      }
    }
  }

  printf("%zu\n", part1);
  printf("%zu\n", part2);

  free_set(vertical_sides, width + 1);
  free_set(horizontal_sides, width);
  free_set(visited, width);
  free_grid(grid, width);
  free(input);

  return EXIT_SUCCESS;
}
