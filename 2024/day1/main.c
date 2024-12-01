#include <stdio.h>
#include <stdlib.h>

#define N 1000

int compare(const void *a, const void *b) {
  int int_a = *((int *)a);
  int int_b = *((int *)b);

  if (int_a == int_b) {
    return 0;
  } else if (int_a < int_b) {
    return -1;
  } else {
    return 1;
  }
}

void part1(int *left, int *right) {
  qsort(left, N, sizeof(int), compare);
  qsort(right, N, sizeof(int), compare);

  int sum = 0;
  for (int i = 0; i < N; i++) {
    sum += abs(right[i] - left[i]);
  }

  printf("%d\n", sum);
}

void part2(int *left, int *right) {
  int sum = 0;
  for (int i = 0; i < N; i++) {
    int count = 0;
    for (int j = 0; j < N; j++) {
      if (right[j] == left[i]) {
        count++;
      }
    }
    sum += count * left[i];
  }

  printf("%d\n", sum);
}

int main(void) {
  FILE *file;
  file = fopen("input.txt", "r");

  int left[N];
  int right[N];
  int index = 0;
  while (fscanf(file, "%d   %d\n", &left[index], &right[index]) != EOF) {
    index++;
  }

  fclose(file);

  part1(left, right);
  part2(left, right);

  return 0;
}
