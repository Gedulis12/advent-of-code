#include <stdio.h>
#include <stdlib.h>


#define MAX_LINE_LENGTH 10

int calculate_game_choice_points(char me, char opponent)
{
	int points = 0;
	switch (opponent) {
		case 'A':
			if (me == 'X') { // need to lose
				points = 3; // since i need to pick scissors
			}
			else if (me == 'Y') {
				points = 1; // since i need to pick rock
			}
			else if (me == 'Z') {
				points = 2; // since i need to pick paper
			}
			break;
		case 'B':
			if (me == 'X') { // need to draw
				points = 1; // since i need to pick rock
			}
			else if (me == 'Y') {
				points = 2; // since i need to pick paper
			}
			else if (me == 'Z') {
				points = 3; // since i need to pick scissors
			}
			break;
		case 'C':
			if (me == 'X') { // need to win
				points = 2; // since i need to pick paper
			}
			else if (me == 'Y') {
				points = 3; // since i need to pick scissors
			}
			else if (me == 'Z') {
				points = 1; // since i need to pick rock
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
				current_points_choice = calculate_game_choice_points(me, opponent);
				current_points_output = 0;
				break;
			case 'Y':
				current_points_choice = calculate_game_choice_points(me, opponent);
				current_points_output = 3;
				break;
			case 'Z':
				current_points_choice = calculate_game_choice_points(me, opponent);
				current_points_output = 6;
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
