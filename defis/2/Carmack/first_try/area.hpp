#ifndef _AREA_HPP
#define _AREA_HPP

struct point {
    int x, y;
};

struct area {
    int x1, y1, x2, y2;
    area();
    area(int x1, int y1, int x2, int y2);
};

bool contains(area &a, area &b);

bool pcontains(area &a, point &p);

void intersect(area &a, area &b, area &out);

int surface(area &a);

#endif
