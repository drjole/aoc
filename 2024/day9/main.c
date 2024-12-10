#include "../utils.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void debug(int *disk, size_t disk_size) {
  for (size_t i = 0; i < disk_size; i++) {
    if (disk[i] == -1) {
      printf(".");
    } else {
      printf("%d", disk[i]);
    }
  }
  printf("\n");
}

int main(void) {
  // Read input
  char *input = read_file("input.txt");
  if (!input) {
    perror("Failed to read input.txt");
    return EXIT_FAILURE;
  }

  // Allocate disk map
  size_t *disk_map = malloc((strlen(input) - 1) * sizeof(size_t));
  if (!disk_map) {
    perror("Failed to allocate memory");
    free(input);
    return EXIT_FAILURE;
  }

  // Populate disk map
  char *p = input;
  size_t disk_map_size = 0;
  while (*p != '\n') {
    disk_map[disk_map_size++] = *p - '0';
    p++;
  }

  // Calculate disk size
  size_t disk_size = 0;
  for (size_t i = 0; i < disk_map_size; i++) {
    disk_size += disk_map[i];
  }

  // Allocate disk
  int *disk = malloc(disk_size * sizeof(int));
  if (!disk) {
    perror("Failed to allocate memory");
    free(disk_map);
    free(input);
    return EXIT_FAILURE;
  }

  // Populate disk
  bool file = true;
  size_t disk_index = 0;
  size_t disk_map_index = 0;
  size_t file_id = 0;
  while (disk_index < disk_size) {
    for (size_t i = 0; i < disk_map[disk_map_index]; i++) {
      if (file) {
        disk[disk_index++] = file_id;
      } else {
        disk[disk_index++] = -1;
      }
    }
    if (file) {
      file_id++;
    }
    file = !file;
    disk_map_index++;
  }

  // Allocate disk copy for use in part 2
  int *disk_copy = malloc(disk_size * sizeof(int));
  if (!disk_copy) {
    perror("Failed to allocate memory");
    free(disk);
    free(disk_map);
    free(input);
  }

  // Populate disk copy for use in part 2
  memcpy(disk_copy, disk, disk_size * sizeof(int));

  // Defragment disk for part 1
  file = true;
  disk_index = 0;
  disk_map_index = 0;
  size_t disk_index_reverse = disk_size - 1;
  while (disk_index < disk_index_reverse) {
    for (size_t i = 0; i < disk_map[disk_map_index]; i++) {
      if (!file) {
        disk[disk_index] = disk[disk_index_reverse];
        disk[disk_index_reverse] = -1;
        while (disk[disk_index_reverse] == -1) {
          disk_index_reverse--;
        }
      }
      disk_index++;
    }
    file = !file;
    disk_map_index++;
  }

  // Checksum for part 1
  size_t part1 = 0;
  for (size_t i = 0; i < disk_size; i++) {
    if (disk[i] != -1) {
      part1 += i * disk[i];
    }
  }

  // Defragment disk for part 2
  size_t last_file_index = disk_map_size - 1;
  disk_index_reverse = disk_size - disk_map[last_file_index];
  for (size_t i = last_file_index;; i -= 2) {
    size_t required_space = disk_map[i];
    size_t free_space = 0;
    disk_index = 0;
    for (size_t j = 0; j < disk_index_reverse; j++) {
      if (disk_copy[j] == -1) {
        free_space++;
      } else {
        free_space = 0;
      }
      disk_index++;
      if (free_space == required_space) {
        break;
      }
    }

    if (free_space == required_space) {
      for (size_t k = 0; k < required_space; k++) {
        disk_copy[--disk_index] = disk_copy[disk_index_reverse + k];
        disk_copy[disk_index_reverse + k] = -1;
      }
    }
    disk_index_reverse -= disk_map[i - 1] + disk_map[i - 2];

    if (i == 0) {
      break;
    }
  }

  // Checksum for part 2
  size_t part2 = 0;
  for (size_t i = 0; i < disk_size; i++) {
    if (disk_copy[i] != -1) {
      part2 += i * disk_copy[i];
    }
  }

  // Print results
  printf("%zu\n", part1);
  printf("%zu\n", part2);

  // Free memory
  free(disk_copy);
  free(disk);
  free(disk_map);
  free(input);
}
