#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE_LENGTH 80

int find_duplicate_letter(int length, char first_half[length], char second_half[length], char *duplicate_letter);
int get_priority_score(char leter);

int main(int argc, char **argv)
{
	char *path;
	char line[MAX_LINE_LENGTH] = {0};
	char duplicate_letter;
	int total = 0;

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
		int half = (strlen(line) - 1) / 2;
		char *first_half = (char *)malloc((half +1) * sizeof(char));
		char *second_half = (char *)malloc((half +1) * sizeof(char));
		int score = 0;
		strncpy(first_half, line, half);
		strncpy(second_half, &line[half], half);
		first_half[half] = '\0';
		second_half[half] = '\0';
		find_duplicate_letter(half, first_half, second_half, &duplicate_letter);
		score = get_priority_score(duplicate_letter);
		printf("%s: %s%s: %s\n%s: %s\n%s: %c\n%s: %d\n\n", "Line", line, "First half", first_half, "Second half", second_half, "Common letter", duplicate_letter, "Score", score);
 		total += score;

	}

	printf("%s: %d\n", "Total score is", total);
	fclose(file);
	return EXIT_SUCCESS;
}

int find_duplicate_letter(int length, char *first_half, char *second_half, char *duplicate_letter)
{
	int i;
	int j;
	for (i = 0; i<length; i++)
	{
		for (j = 0; j<length; j++)
		{
			if (first_half[i] == second_half[j])
			{
				*duplicate_letter = first_half[i];
				return 0;
			}
		}
	}
	return 0;
}

int get_priority_score(char letter)
{
	int score = 0;
	if (letter >= 97)
	{
		score = letter - 96;
	}
	else
	{
		score = letter - 64 + 26;
	}
	return score;
}
