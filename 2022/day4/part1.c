#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE_LENGTH 80

int main(int argc, char **argv)
{
    char *path;
    char line[MAX_LINE_LENGTH] = {0};
    int final_count = 0;


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
        int ints[4];
        int inner_pair_index;
        int outer_pair_index;
        memset(&ints[0], 0, sizeof ints);

        char *pair_1 = strtok(line, ",");
        char *pair_2 = strtok(NULL, ",");

        ints[0] = atoi(strtok(pair_1, "-"));
        ints[1] = atoi(strtok(NULL, "-"));

        ints[2] = atoi(strtok(pair_2, "-"));
        ints[3] = atoi(strtok(NULL, "-"));


        if ((ints[0] <= ints[2] && ints[1] >= ints[3]) || (ints[2] <= ints[0] && ints[3] >= ints[1]))
        {
            final_count += 1;
        }
    }
    
    printf("%s: %d\n", "final count", final_count);

    fclose(file);
    return EXIT_SUCCESS;
}
