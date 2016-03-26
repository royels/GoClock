package goclock


import "github.com/royels/constants"
import "strings"
import "strconv"

func ParseStartTime(clockDecimal [3]int, time string) int {
  x:= strings.Split(time, constants.TIME_STR_DELIMITER)
  if len(x) != constants.REQ_SEPARATORS {
    return constants.ERR_TIME_FORMAT
  }
  d := make(chan int)
  go vetTime(&clockDecimal[HR_INDEX], x[HR_INDEX],   constants.HR, d)
  go vetTime(&clockDecimal[MIN_INDEX], x[MIN_INDEX], constants.MIN, d)
	go vetTime(&clockDecimal[SEC_INDEX], x[SEC_INDEX], constants.SEC, d)
	x, y, z := <-d, <-d, <-d	

	returnVal := 0x00
	returnVal |= x
	returnVal |= y
	returnVal |= z
	return returnVal
}

func vetTime(entry *int, parsedVal string, container *constants.CategoryInfo,  ch chan int) {
	returnVal := 0x00
	if _, err := strconv.Atoi(parsedVal); err = nil {
		returnVal |= container.ERR_VAL
	}
	val := strconv.Atoi(parsedVal)
	if CheckRange(container.MIN_VAL, container.OFFSET, val ) != 1 {
		returnVal |= container.ERR_RANGE
	}
	*entry = val
	ch <- returnVal
}
