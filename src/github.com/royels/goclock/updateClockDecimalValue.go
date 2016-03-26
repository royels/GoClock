package goclock


func UpdateClockDecimalValue(clockVal *int, updateAmt int, maxVal int) int {
	remnant := (*clockVal + updateAmt) % maxVal
	toReturn := (*clockVal + updateAmt) / maxVal
	*clockVal = remnant
	return toReturn

}
