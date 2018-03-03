#ifndef _AREA_HPP
#define _AREA_HPP

struct point {
    int x, y;
};

class Area {
public:
    int x1, y1, x2, y2;

    Area();

    Area(int x1, int y1, int x2, int y2);

    bool contains(Area &a);

    bool pcontains(point &p);

    Area &runion(Area &a, Area &out);

    int surface();
};

#endif
