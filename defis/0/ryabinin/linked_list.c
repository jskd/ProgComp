#include <stdlib.h>

#include "linked_list.h"

struct linked_list *linked_list_head(void *data) {
    struct linked_list *head = malloc(sizeof(struct linked_list));
    if(!head)
        return NULL;
    
    head->data = data;
    head->next = NULL;
    head->head = head;

    return head;
}

struct linked_list *linked_list_add(struct linked_list *ll, void *data) {
    struct linked_list *next;

    if(!ll)
        return linked_list_head(data);

    next = malloc(sizeof(struct linked_list));
    if(!next)
        return NULL;

    ll->next = next;
    next->data = data;
    next->next = NULL;
    next->head = ll->head;

    return next;
}

void linked_list_exec_and_clean(struct linked_list *head, void (*f)(void *)) {
    struct linked_list *prev_head = NULL;
    
    while(head) {
        if(prev_head)
            free(prev_head);
        f(head->data);
        prev_head = head;
        head = head->next;
    }

    if(prev_head)
        free(prev_head);
}
