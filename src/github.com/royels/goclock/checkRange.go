package goclock

func CheckRange(minRange, offset, value int) int {
	if(IsNegative(offset) == 1) {
		return -1
	}
	sum :=  minRange + offset
	if value < minRange {
		return 0
	}
	if value > sum {
		return 0
	} else {
		return 1
	}
}





