package goclock


import "github.com/royels/constants"

func ConvertToBCD(clockDecimal [] int, clockBCD []int) {
	go craftBCD(clockDecimal[constants.HR_INDEX],  &clockBCD[constants.HR_INDEX])
	go craftBCD(clockDecimal[constants.MIN_INDEX], &clockBCD[constants.MIN_INDEX])
	go craftBCD(clockDecimal[constants.SEC_INDEX], &clockBCD[constants.SEC_INDEX])
}

func craftBCD(src int, dest *int) {
	rem := src % 10
	div := src / 10
	temp1 := div << 4
	temp2 := temp1|rem 
	*dest = temp2
}
