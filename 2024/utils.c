#include <stdlib.h>
#include <stdio.h>

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
