#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>

#include "multiset.h"
#include "linked_list.h"

int main(int argc, char *argv[]) {
    int fd, i, node_num;
    char *dict, *nl_pos;
    struct stat stat;

    if(argc < 3) {
        fprintf(stderr, "Not enough arguments\n");
        return 1;
    }

    node_num = argc - 2;
    struct node {
        struct multiset ms;
        struct linked_list *anagrams;
    } nodes[node_num];

    fd = open(argv[1], O_RDWR);
    if(fd == -1) {
        fprintf(stderr, "%s: Cannot open file\n", argv[1]);
        return 1;
    }
    
    if(fstat(fd, &stat)) {
        fprintf(stderr, "Some `fstat` error\n");
        return 1;
    }
    
    dict = mmap(0, stat.st_size, PROT_READ | PROT_WRITE, MAP_PRIVATE, fd, 0);
    close(fd);
    if(dict == MAP_FAILED) {
        fprintf(stderr, "Some `mmap` error\n");
        return 1;
    }

    for(i = 0; i < node_num; i++) {
        MULTISET_INIT(&nodes[i].ms);
        multiset_fill(&nodes[i].ms, argv[i + 2]);
        nodes[i].anagrams = NULL;
    }

    while(*dict) {
        nl_pos = strchr(dict, '\n');
        if(!nl_pos)
            nl_pos = dict + strlen(dict);
        *nl_pos = '\0';
        
        for(i = 0; i < node_num; i++)
            if(nodes[i].ms.len == (nl_pos - dict) &&
               !multiset_wrdcmp(&nodes[i].ms, dict))
                nodes[i].anagrams =
                    linked_list_add(nodes[i].anagrams, strdup(dict));

        dict = nl_pos + 1;
    }

    void print_string(void *data) {
        printf("%s\n", (char *)data);
        free(data);
    }

    for(i = 0; i < node_num; i++) {
        printf("%s:\n", argv[i + 2]);
        if(!nodes[i].anagrams)
            continue;
        nodes[i].anagrams = nodes[i].anagrams->head;
        linked_list_exec_and_clean(nodes[i].anagrams, print_string);
    }

    munmap(dict, stat.st_size);    
    return 0;
}
