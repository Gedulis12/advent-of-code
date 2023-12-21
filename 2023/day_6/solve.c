#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#define MAX_LINE_LENGTH 80

void parse_line(char *line, int *array_to_populate);
int get_len(int* arr);
int get_valid_action_count(long int time, long int distance);

int main(int argc, char **argv)
{
    char *path;
    char line[MAX_LINE_LENGTH] = {0};

    if (argc < 1)
    {
        return EXIT_FAILURE;
    }

    path = argv[1];

    FILE *file = fopen(path, "r");
    if (!file)
    {
        perror(path);
        return EXIT_FAILURE;
    }

    // Lines to variables
    char *time_line;
    char *distance_line;

    while (fgets(line, MAX_LINE_LENGTH, file))
    {
        int i = 0;
        while (line[i] != '\0')
        {
            i++;
        }
        i++;

        char *find_time = strstr(line, "Time");
        if (find_time != NULL)
        {
            time_line = malloc(sizeof(char) * i);
            time_line = strncpy(time_line, line, i);
            time_line[i] = '\0';
        }

        char *find_distance = strstr(line, "Distance");
        if (find_distance != NULL)
        {
            distance_line = malloc(sizeof(char) * i);
            distance_line = strncpy(distance_line, line, i);
            distance_line[i] = '\0';
        }
    }


    int times[MAX_LINE_LENGTH] = {0};
    int distances[MAX_LINE_LENGTH] = {0};

    int time_idx = 0;
    int distance_idx = 0;


    parse_line(time_line, times);
    parse_line(distance_line, distances);

    int res = 1;
    for (int i = 0; i < get_len(times); i++)
    {
        int valid_count = get_valid_action_count(times[i], distances[i]);
        res *= valid_count;
    }

    printf("Result for part 1: %i\n", res);

    char p2_time[MAX_LINE_LENGTH] = {0};
    char p2_distance[MAX_LINE_LENGTH] = {0};

    for (int i = 0; i < get_len(times); i++)
    {
        char tmp_buf_time[MAX_LINE_LENGTH] = {0};
        char tmp_buf_dist[MAX_LINE_LENGTH] = {0};

        sprintf(tmp_buf_time, "%d", times[i]);
        sprintf(tmp_buf_dist, "%d", distances[i]);


        strcat(p2_time, tmp_buf_time);
        strcat(p2_distance, tmp_buf_dist);
    }
    long int p2_time_int = strtol(p2_time, NULL, 10);
    long int p2_distance_int = strtol(p2_distance, NULL, 10);

    printf("Result for part 2: %i\n", get_valid_action_count(p2_time_int, p2_distance_int));



    free(time_line);
    free(distance_line);
    return 0;
}

void parse_line(char *line, int *array_to_populate)
{
    int idx = 0;
    char *token = strtok(line, " ");

    while (token != NULL)
    {
        long int curr_int;
        if ((curr_int = strtol(token, NULL, 10)) != 0)
        {
            array_to_populate[idx] = (int)curr_int;
            idx ++;
        }
        token = strtok(NULL, " ");
    }
}

int get_len(int* arr)
{
    int idx = 0;
    while (arr[idx] != 0)
    {
        idx ++;
    }
    return idx;
}

int get_valid_action_count(long int time, long int distance)
{
    long int valid = 0;
    for (long int i = 0; i<= time; i++)
    {
        long int new_distance = i * (time - i);
        if (new_distance > distance)
        {
            valid++;
        }
    }

    return valid;
}
