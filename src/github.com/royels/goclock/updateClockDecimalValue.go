package goclock


import "constants"


func UpdateClockDecimalValue(clockVal *int, updateAmt int, maxVal int) {
	remnant = (*clockVal + updateAmt) % maxVal
	toReturn = (*clockVal + updateAmt) / maxVal
	*clockVal = remnant
	return toReturn

}
