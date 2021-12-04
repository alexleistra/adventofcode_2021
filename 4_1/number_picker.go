package main

type NumberPicker struct {
	observerList     []*BingoBoard
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
	return len(np.numbers) > np.currentNumberIdx
}

func (np *NumberPicker) register(o *BingoBoard) {
	np.observerList = append(np.observerList, o)
}

func (np *NumberPicker) deregister(o *BingoBoard) {
	for i, observer := range np.observerList {
		if o.getID() == observer.getID() {
			np.observerList = append(np.observerList[:i], np.observerList[i+1:]...)
			break
		}
	}
}

func (np *NumberPicker) notifyAll() {
	for _, observer := range np.observerList {
		observer.update(np.currentNumber)
	}
}

func (np *NumberPicker) findWinningBoard() (bool, int, int) {
	for _, board := range np.observerList {
		if board.winner {
			return true, board.getID(), board.unmarkedSum * np.currentNumber
		}
	}

	return false, 0, 0
}
