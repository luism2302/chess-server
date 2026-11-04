package game

const (
	EmptyVal = 0
)

type empty struct {
	color     color
	val       int
	pieceRune rune
}

func newEmpty() *empty {
	return &empty{
		color:     noColor,
		val:       EmptyVal,
		pieceRune: '\u25AA',
	}
}

func (e *empty) getAvailableMoves(pos coords, game *Game) []coords {
	return []coords{}
}

func (e *empty) getPieceVal() int {
	return e.val
}

func (e *empty) getPieceColor() color {
	return e.color
}

func (e *empty) getPieceRune() rune {
	return e.pieceRune
}

func (e *empty) String() string {
	return "Empty{}"
}
