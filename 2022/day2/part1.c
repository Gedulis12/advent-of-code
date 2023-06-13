#include <stdio.h>
#include <stdlib.h>


#define MAX_LINE_LENGTH 10

int calculate_game_output(char me, char opponent)
{
	int points = 0;
	switch (me) {
		case 'X':
			if (opponent == 'A') {
				points = 3;
			} else if (opponent == 'B') {
				points = 0;
			} else if (opponent == 'C') {
				points = 6;
			}
			break;
		case 'Y':
			if (opponent == 'A') {
				points = 6;
			} else if (opponent == 'B') {
				points = 3;
			} else if (opponent == 'C') {
				points = 0;
			}
			break;
		case 'Z':
			if (opponent == 'A') {
				points = 0;
			} else if (opponent == 'B') {
				points = 6;
			} else if (opponent == 'C') {
				points = 3;
			}
			break;
	}	
	return points;
}

int main(int argc, char **argv)
{
	char *path;
	char line[MAX_LINE_LENGTH] = {0};
	int current_points_choice = 0;
	int current_points_output = 0;
	int total_points = 0;


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
		char opponent = *line;
		char me = *(line + 2);
//		printf("%c-%c\n", opponent, me);
		switch(me)
		{
			case 'X':
				current_points_choice = 1;
				current_points_output = calculate_game_output(me, opponent);
				break;
			case 'Y':
				current_points_choice = 2;
				current_points_output = calculate_game_output(me, opponent);
				break;
			case 'Z':
				current_points_choice = 3;
				current_points_output = calculate_game_output(me, opponent);
				break;
		}

		total_points += current_points_choice;
		total_points += current_points_output;

		current_points_choice = 0;
		current_points_output = 0;
	}
	printf("%s%d\n", "Total points: ", total_points);

	fclose(file);
	return EXIT_SUCCESS;
}
