package goclock


import "constants"
import "strings"



func ParseStartTime(clockDecimal [3]int, time string) {
  x:= strings.Split(time, constants.TIME_STR_DELIMITER)
  if len(x) != constants.REQ_SEPARATORS {
    return constants.ERR_TIME_FORMAT
  }
  d := make(chan int)
  go vetTime(&clockDecimal[HR_INDEX], x[HR_INDEX], d)
  go vetTime(&clockDecimal[MIN_INDEX], x[MIN_INDEX], d)
	go vetTime(&clockDecimal[SEC_INDEX], x[SEC_INDEX], d)
	

}

func vetTime(entry *int, parsedVal string, ch chan int) {



}
