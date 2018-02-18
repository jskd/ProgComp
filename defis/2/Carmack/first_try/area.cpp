#include "area.hpp"

area::area() { }

area::area(int x1, int y1, int x2, int y2) {
    this->x1 = x1;
    this->y1 = y1;
    this->x2 = x2;
    this->y2 = y2;
}

bool contains(area &a,  area &b) {
    return b.x1 >= a.x1 && b.x2 <= a.x2 &&
           b.y1 >= a.y1 && b.y2 <= a.y2;
}

bool pcontains(area &a, point &p) {
    return p.x >= a.x1 && p.y >= a.y1 &&
           p.x <= a.x2 && p.y <= a.y2;
}

void intersect(area &a, area &b, area &out) {
    out.x1 = (a.x1 < b.x1) ? a.x1 : b.x1;
    out.y1 = (a.y1 < b.y1) ? a.y1 : b.y1;
    out.x2 = (a.x2 > b.x2) ? a.x2 : b.x2;
    out.y2 = (a.y2 > b.y2) ? a.y2 : b.y2;
}

int surface(area &a) {
    return (a.x2 - a.x1) * (a.y2 - a.y1);
}
