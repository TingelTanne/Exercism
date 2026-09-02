package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	occupiedCells := 0

	for _, rank := range cb[file] {
		if rank {
			occupiedCells++
		}
	}
	return occupiedCells
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	occupiedCells := 0

	if rank > 0 && rank < 9 {
		for _, file := range cb {
			if file[rank-1] {
				occupiedCells++
			}
		}
	}

	return occupiedCells
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	totalCells := 0

	for _, file := range cb {
		totalCells += len(file)
	}
	return totalCells
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupiedCells := 0
	for _, file := range cb {
		for _, rank := range file {
			if rank {
				occupiedCells++
			}
		}
	}

	return occupiedCells
}
