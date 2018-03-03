#ifndef _AREA_HPP
#define _AREA_HPP

struct point {
    int x, y;
};

template<typename coord_type>
class Area {
public:
    coord_type x1, y1, x2, y2;

    Area() { }
    Area(coord_type x1, coord_type y1, coord_type x2, coord_type y2) :
        x1(x1), y1(y1), x2(x2), y2(y2) { }

    template<typename snd_coord_type>
    Area &operator=(const Area<snd_coord_type> &a) {
        x1 = a.x1;
        x2 = a.x2;
        y1 = a.y1;
        y2 = a.y2;
    }

    template<typename snd_coord_type>
    bool contains(Area<snd_coord_type> &a) {
        return a.x1 >= x1 && a.x2 <= x2 &&
               a.y1 >= y1 && a.y2 <= y2;
    }

    bool pcontains(point &p) {
        return p.x >= x1 && p.y >= y1 &&
               p.x <= x2 && p.y <= y2;
    }

    template<typename snd_coord_type, typename trd_coord_type>
    Area<trd_coord_type> &runion(Area<snd_coord_type> &a,
                                 Area<trd_coord_type> &out) {
        out.x1 = (x1 < a.x1) ? x1 : a.x1;
        out.y1 = (y1 < a.y1) ? y1 : a.y1;
        out.x2 = (x2 > a.x2) ? x2 : a.x2;
        out.y2 = (y2 > a.y2) ? y2 : a.y2;
        return out;
    }

    int surface() {
        return (x2 - x1) * (y2 - y1);
    }
};

#endif
