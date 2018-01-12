#include "multiset.h"

#define CHECK_ID(id) \
    (((id) < 0 || (id) > 26) ? 26 : (id))
#define LETTER_ID(l) \
    CHECK_ID(((l) >= 97) ? ((l) - 97) : ((l) - 65))

void multiset_fill(struct multiset *ms, const char *str) {
    if(!*str)
        return;
    
    do {
        ms->occur[LETTER_ID(*str)]++;
        ms->len++;
    }
    while(*(++str));
}

int multiset_wrdcmp(const struct multiset *ms, const char *str) {
    int l_id;
    struct multiset inc_ms;

    if(!*str)
        return 42;
    
    MULTISET_INIT(&inc_ms);
    
    do {
        l_id = LETTER_ID(*str);
        if((++inc_ms.occur[l_id]) > ms->occur[l_id])
            return 42;
        inc_ms.len++;
    } while(*(++str));

    return memcmp(&inc_ms, ms, sizeof *ms);
}
