package goclock


import "github.com/royels/constants"


func UpdateClockDecimal(clockDecimal []int, interval int) {
	a := UpdateClockDecimalValue(&clockDecimal[constants.SEC_INDEX], interval, constants.MAX_SEC)
	b := UpdateClockDecimalValue(&clockDecimal[constants.MIN_INDEX], a, constants.MAX_MINUTE)
	_ = UpdateClockDecimalValue(&clockDecimal[constants.HR_INDEX], b, constants.MAX_HR) 

}
