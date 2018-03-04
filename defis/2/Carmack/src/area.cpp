#include "area.hpp"

Area::Area() { }

Area::Area(int x1, int y1, int x2, int y2) :
    x1(x1), y1(y1), x2(x2), y2(y2) { }

bool Area::contains(Area &a) {
    return a.x1 >= x1 && a.x2 <= x2 &&
        a.y1 >= y1 && a.y2 <= y2;
}

bool Area::pcontains(point &p) {
    return p.x >= x1 && p.y >= y1 &&
        p.x <= x2 && p.y <= y2;
}

Area &Area::runion(Area &a, Area &out) {
    out.x1 = (x1 < a.x1) ? x1 : a.x1;
    out.y1 = (y1 < a.y1) ? y1 : a.y1;
    out.x2 = (x2 > a.x2) ? x2 : a.x2;
    out.y2 = (y2 > a.y2) ? y2 : a.y2;
    return out;
}

int Area::surface() {
    return (x2 - x1) * (y2 - y1);
}
