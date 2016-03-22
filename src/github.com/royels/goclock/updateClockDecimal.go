package goclock


import "constants"


func updateClockDecimal(clockDecimal []int, interval int) {
	const a, b, c int
	a = updateClockDecimalValue(&clockDecimal[constants.SEC_INDEX], interval, constants.MAX_SEC)
	b = updateClockDecimalValue(&clockDecimal[constants.MIN_INDEX], a, constants.MAX_MINUTE)
	c = updateClockDecimalValue(&clockDecimal[constants.HR_INDEX], b, constants.MAX_HR) 
}
