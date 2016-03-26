package main

import "fmt"
import "github.com/royels/constants"
import "github.com/royels/goclock"
import "os"
import "strconv"

func main() {
	numTicks := constants.DEF_TICKS	
	intervalSeconds := constants.DEF_INTERVALS
	
	if len(os.Args) < constants.MIN_ARGS || len(os.Args) > constants.MAX_ARGS {
		goclock.Usage(os.Args[0])
		os.Exit(1)
	}
	var returnVal int
	var clockDecimal [constants.LIST_LENGTH]int
	returnVal = goclock.ParseStartTime(clockDecimal, os.Args[constants.TIME_INDEX])
	if (returnVal&constants.ERR_TIME_FORMAT) != 0 {
		fmt.Fprintf(os.stderr, constants.STR_ERR_TIME_FORMAT)
	}
	// implement error handling. Too bored at this stage in the night to do it....
	// intervalSeconds, err := strconv.Atoi(os.Args[constants.INTERVAL_INDEX])
	// numTicks, err := strconv.Atoi(os.Args[constants.TICKS_INDEX])
	var clockBCD [constants.LIST_LENGTH]int
	goclock.ConvertToBCD(clockDecimal, clockBCD)
	goclock.DisplayBCDClock(clockBCD)
	for i := 0; i < numTicks; i++ {
		goclock.UpdateClockDecimal(clockDecimal, intervalSeconds)
		goclock.ConvertToBCD(clockDecimal, clockBCD)
		goclock.DisplayBCDClock(clockBCD)
	}

}
