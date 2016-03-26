package goclock


import "fmt"
import "github.com/royels/constants"

func Usage (programName string) {
  fmt.Printf(constants.STR_USAGE, programName,
  constants.MIN_HR, constants.MAX_HR - 1,
  constants.MIN_MINUTE, constants.MAX_MINUTE - 1,
  constants.MIN_SEC, constants.MAX_SEC - 1,
  constants.MIN_INTERVAL, constants.MAX_INTERVAL, constants.DEF_INTERVAL,
  constants.MIN_TICKS, constants.MAX_TICKS, constants.DEF_TICKS);
}



