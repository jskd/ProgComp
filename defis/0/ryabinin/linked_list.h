#ifndef _LINKED_LIST_H
#define _LINKED_LIST_H

struct linked_list {
    void *data;
    struct linked_list *next;
    struct linked_list *head;
};

struct linked_list *linked_list_head(void *data);

struct linked_list *linked_list_add(struct linked_list *ll, void *data);

void linked_list_exec_and_clean(struct linked_list *head, void (*f)(void *));

#endif
