package main

type desk struct {
	x, y    int         // ячейки на доске
	figures chessPieces // общее допустимое колличество фигур на доске
}
type chessPieces struct {
	king, queen, bishop, knight, rook, pawn int
}

func NewDesk() *desk {
	figures := chessPieces{
		king:   1,
		queen:  1,
		bishop: 2,
		knight: 2,
		rook:   2,
		pawn:   8,
	}

	return &desk{
		x:       8,
		y:       8,
		figures: figures,
	}
}
