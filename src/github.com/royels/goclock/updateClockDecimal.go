package goclock


import "github.com/royels/constants"


func UpdateClockDecimal(clockDecimal []int, interval int) {
	a := updateClockDecimalValue(&clockDecimal[constants.SEC_INDEX], interval, constants.MAX_SEC)
	b := updateClockDecimalValue(&clockDecimal[constants.MIN_INDEX], a, constants.MAX_MINUTE)
	c := updateClockDecimalValue(&clockDecimal[constants.HR_INDEX], b, constants.MAX_HR) 
}
