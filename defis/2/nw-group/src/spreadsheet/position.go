package spreadsheet

type position struct {
	row uint32
	col uint32
}

func (pos1 position) before(pos2 position) bool {
	return pos1.row <= pos2.row && pos1.col <= pos2.col
}

func (pos1 position) after(pos2 position) bool {
	return pos1.row >= pos2.row && pos1.col >= pos2.col
}

type area struct {
	src position
	dst position
}

func (area area) contains(pos position) bool {
	return pos.after(area.src) && pos.before(area.dst)
}
