package game

const (
	EmptyVal = 0
)

type empty struct {
	val       int
	pieceRune rune
}

func newEmpty() *empty {
	return &empty{
		val:       EmptyVal,
		pieceRune: '\u25AA',
	}
}

func (e *empty) getAvailableMoves(pos coords, board board) []coords {
	return []coords{}
}

func (e *empty) getPieceVal() int {
	return e.val
}

func (e *empty) getPieceColor() bool {
	return false
}

func (e *empty) getPieceRune() rune {
	return e.pieceRune
}

func (e *empty) String() string {
	return "Empty{}"
}
