
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE_LENGTH 80


int main(int argc, char **argv)
{
	char *path;
	char line[MAX_LINE_LENGTH] = {0};
	char duplicate_letter;
	int line_counter = 0;
	int total = 0;
	int group = 0;
	int score = 0;
	int group_item = 0;

	char *one = NULL;
	char *two = NULL;
	char *three = NULL;

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
		score = 0;
		group = ((line_counter / 3) + 1);
		group_item ++;
		line_counter ++;

		switch(group_item)
		{
			case(1):
				one = malloc(sizeof(char) * (strlen(line) + 1));
				one = strcpy(one, line);
				break;
			case(2):
				two = malloc(sizeof(char) * (strlen(line) + 1));
				two = strcpy(two, line);
				break;
			case(3):
				three = malloc(sizeof(char) * (strlen(line) + 1));
				three = strcpy(three, line);
		}

		if (line_counter % 3 == 0)
		{
			int common_found = 0;

			for (int i = 0; i<strlen(one); i++)
			{
				for (int j = 0; j<strlen(two); j++)
				{
					for (int k = 0; k<strlen(three); k++)
					{
						if (one[i] == two[j] && two[j] == three[k])
						{
							printf("%s: %c\n", "common letter is", three[k]);
							common_found = 1;
							break;
						}
					}
					if (common_found)
						break;
				}
				if (common_found)
					break;
			}

			group_item = 0;

			free(one);
			free(two);
			free(three);
			one = NULL;
			two = NULL;
			three = NULL;
		}


	}

	fclose(file);
	return EXIT_SUCCESS;
}
