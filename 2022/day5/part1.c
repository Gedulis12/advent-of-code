#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE_LENGTH 80

struct node
{
    char data;
    struct node *next;
};


int main(int argc, char **argv)
{
    char *path;
    char line[MAX_LINE_LENGTH] = {0};
    struct node *top = NULL;
    struct node *new_node = (struct node*)malloc(sizeof(struct node));
    struct node *stack[9];

    if (argc < 1)
        return 1;
    path = argv[1];

    FILE *file = fopen(path, "r");

    if (!file)
    {
        perror(path);
        return EXIT_FAILURE;
    }

    while (fgets(line, MAX_LINE_LENGTH, file))
    {
        // parse the input stacks
        if (line[0] != 'm' || line[1] != '1')
        {
            int stack_id = 0;
            for (int i = 1; i < strlen(line); i = i + 4)
            {
                stack[stack_id] = (struct node*)malloc(sizeof(struct node));
                stack[stack_id]->data = line[i];
                stack[stack_id]->next = NULL;

                printf("%c", line[i]);
                stack_id ++;
            }
            printf("\n");

        for (int i = 0; i < 10; i++)
        {
            printf("data of stack #%d: %c", i, stack[stack_id]->data);
        }
        }
    }
    fclose(file);
    return 0;
}
