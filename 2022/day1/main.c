#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE_LENGTH 80

struct dyn_array 
{
	int *data;
	int size;
	int capacity;
};

void init_dyn_array(struct dyn_array *array)
{
	array->data = NULL;
	array->size = 0;
	array->capacity = 0;
}

void add_to_dyn_array(struct dyn_array *array, int value)
{
	if (array->size == array->capacity)
	{
		array->capacity = (array->capacity == 0) ? 1 : array->capacity * 2;
		int *new_data = realloc(array->data, array->capacity * sizeof(int));
		if (new_data == NULL)
		{
			printf("Memory allocation failed\n");
			free(array->data);
			exit(1);
		}
		array->data = new_data;
	}
	array->data[array->size++] = value;
}

void print_dyn_array(const struct dyn_array *array)
{
	for (int i = 0; i < array->size; i++)
	{
		printf("%d ", array->data[i]);
	}
	printf("\n");
}

void free_dyn_array(struct dyn_array *array)
{
	free(array->data);
	init_dyn_array(array);
}

int compare_descending(const void *a, const void *b)
{
	int value_a = *((int*)a);
	int value_b = *((int*)b);
	return value_b - value_a;
}

void sort_dyn_array(struct dyn_array *array)
{
	qsort(array->data, array->size, sizeof(int), compare_descending);
}


int main(int argc, char **argv)
{
	struct dyn_array array;
	char *path;
	char line[MAX_LINE_LENGTH] = {0};
	long current_total = 0;
	char *tmp;
	int total = 0;

	init_dyn_array(&array);

	if (argc < 1)
		return EXIT_FAILURE;
	path = argv[1];

	FILE *file = fopen(path, "r");

	if (!file)
	{
		perror(path);
		return EXIT_FAILURE;
	}

	while (fgets(line, MAX_LINE_LENGTH, file))
	{
		if (strcmp(line, "\n") != 0)
		{
			current_total += strtol(line, &tmp, 10);
		}
		else
		{
			add_to_dyn_array(&array, current_total);
			current_total = 0;
		}
	}

	sort_dyn_array(&array);	

	for (int i = 0; i < 3; i++)
	{
		printf("%s%d%s%d\n", "Elf #", i+1, " carries calories: ", array.data[i]);
		total += array.data[i];
	}

	printf("%s%d\n", "In total that makes: ", total);

	
	fclose(file);
	return EXIT_SUCCESS;
}
