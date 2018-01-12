#ifndef _MULTISET_H
#define _MULTISET_H

#include <string.h>

struct multiset {
    int occur[27];
    int len;
};

#define MULTISET_INIT(ms) (memset((ms), 0, sizeof *(ms)))

void multiset_fill(struct multiset *ms, const char *str);

int multiset_wrdcmp(const struct multiset *ms, const char *str);

#endif
