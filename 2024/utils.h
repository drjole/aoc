#ifndef IO_H
#define IO_H

#include <stdbool.h>
#include <stdio.h>

char *read_file(char *path);

size_t grid_width(char *s);
size_t grid_height(char *s);
char **make_grid(char *s);
void print_grid(char **grid, size_t width, size_t height);
void free_grid(char **grid, size_t width);
bool in_bounds(int x, int y, int width, int height);
bool is_neighbor(int ax, int ay, int bx, int by);

#endif
