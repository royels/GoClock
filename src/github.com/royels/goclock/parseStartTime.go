package goclock


import "github.com/royels/constants"
import "strings"
import "strconv"
import "fmt"

func ParseStartTime(clockDecimal [3]int, time string) int {
  x:= strings.Split(time, constants.TIME_STR_DELIMITER)
  if len(x) != constants.REQ_SEPARATORS {
    return constants.ERR_TIME_FORMAT
  }
	fmt.Printf("%q\n", x) 
  d := make(chan int)
  go vetTime(&clockDecimal[constants.HR_INDEX], x[constants.HR_INDEX],   constants.HR, d)
  go vetTime(&clockDecimal[constants.MIN_INDEX], x[constants.MIN_INDEX], constants.MIN, d)
	go vetTime(&clockDecimal[constants.SEC_INDEX], x[constants.SEC_INDEX], constants.SEC, d)
	a, b, c := <-d, <-d, <-d	

	returnVal := 0x00
	returnVal |= a
	returnVal |= b
	returnVal |= c
	return returnVal
}

func vetTime(entry *int, parsedVal string, container *constants.CategoryInfo,  ch chan int) {
	returnVal := 0x00
	val, err := strconv.Atoi(parsedVal);
	if err != nil {
		returnVal |= container.ERR_VAL
	}
	if CheckRange(container.MIN_VAL, container.OFFSET, val ) != 1 {
		returnVal |= container.ERR_RANGE
	}
	*entry = val
	ch <- returnVal
}
