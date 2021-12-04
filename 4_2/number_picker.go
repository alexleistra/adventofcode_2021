package main

type NumberPicker struct {
	boards           []*BingoBoard
	numbers          []int
	currentNumberIdx int
	currentNumber    int
}

func NewNumberPicker(numbers []int) *NumberPicker {
	return &NumberPicker{
		numbers: numbers,
	}
}

func (np *NumberPicker) pickNextNumber() {
	np.currentNumber = np.numbers[np.currentNumberIdx]
	np.currentNumberIdx = np.currentNumberIdx + 1
	np.notifyAll()
}

func (np *NumberPicker) hasNextNumber() bool {
	return len(np.numbers) > np.currentNumberIdx && len(np.boards) > 0
}

func (np *NumberPicker) register(o *BingoBoard) {
	np.boards = append(np.boards, o)
}

func (np *NumberPicker) deregister(o *BingoBoard) {
	for i, observer := range np.boards {
		if o.getID() == observer.getID() {
			np.boards = append(np.boards[:i], np.boards[i+1:]...)
			break
		}
	}
}

func (np *NumberPicker) notifyAll() {
	for _, board := range np.boards {
		board.update(np.currentNumber)
	}
}

func (np *NumberPicker) findWinningBoards() (bool, []int) {
	var winningScores []int
	winnerFound := false
	for i := len(np.boards) - 1; i >= 0; i-- {
		if np.boards[i].winner {
			winnerFound = true

			// winning score
			winningScores = append(winningScores, np.boards[i].unmarkedSum*np.currentNumber)

			// remove this board
			np.boards = append(np.boards[:i], np.boards[i+1:]...)
		}
	}

	return winnerFound, winningScores
}
