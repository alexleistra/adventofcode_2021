package main

type BingoBoard struct {
	id          int
	board       [][]int
	rowStamps   []int
	colStamps   []int
	winner      bool
	unmarkedSum int
}

func NewBingoBoard(id int, board [][]int) *BingoBoard {
	var sum int
	for _, nr := range board {
		for _, nc := range nr {
			sum += nc
		}
	}

	return &BingoBoard{
		id:          id,
		board:       board,
		rowStamps:   make([]int, 5),
		colStamps:   make([]int, 5),
		winner:      false,
		unmarkedSum: sum,
	}
}

func (b *BingoBoard) update(number int) {
	for i, nr := range b.board {
		for k, nc := range nr {
			if nc == number {
				b.unmarkedSum = b.unmarkedSum - nc
				b.rowStamps[i] = b.rowStamps[i] + 1
				if b.rowStamps[i] == 5 {
					b.winner = true
				}

				b.colStamps[k] = b.colStamps[k] + 1
				if b.colStamps[k] == 5 {
					b.winner = true
				}
			}
		}
	}
}

func (b *BingoBoard) getID() int {
	return b.id
}
