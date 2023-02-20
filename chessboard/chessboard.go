package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (quantityOfFilledPlaces int) {
	quantityOfFilledPlaces = 0

	for key, value := range cb {
		if file == key {
			for _, fileValue := range value {
				if fileValue {
					quantityOfFilledPlaces++
				}
			}
		}
	}

	return quantityOfFilledPlaces
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (quantityOfFilledPlaces int) {
	quantityOfFilledPlaces = 0

	for _, value := range cb {
		for index, fileValue := range value {
			if fileValue && index+1 == rank {
				quantityOfFilledPlaces++
			}
		}
	}

	return quantityOfFilledPlaces
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (quantityOfPlaces int) {
	quantityOfPlaces = 0

	for _, value := range cb {
		for range value {
			quantityOfPlaces++
		}
	}

	return quantityOfPlaces
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (quantityOfPlaces int) {
	quantityOfPlaces = 0

	for _, value := range cb {
		for _, file := range value {
			if file {
				quantityOfPlaces++
			}
		}
	}

	return quantityOfPlaces
}
