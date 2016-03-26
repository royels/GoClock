package goclock

import "fmt"
import "github.com/royels/constants"




func displayBCDClock(clockBCD []int) {
	temp1 := constants.FIRST_BITMASK
	temp2 := constants.SECOND_BITMASK
	for i := 0; i < constants.ROW_LENGTH; i++ {
		for j := 0; j < constants.COLUMN_LENGTH; j++ {
			val = temp1&clockBCD[j]
			if val == 0 {
				PrintChar(constants.UNLIT_LED)
			}	else {
				PrintChar(constants.LIT_LED)
			}
			val2 = temp2&clockBCD[j]
			if val2 == 0 {
				PrintChar(constants.UNLIT_LED)
			} else {
				PrintChar(constants.LIT_LED)
			}
			PrintChar(' ')
	}
	temp1 = temp1 >> 1
	temp2 = temp2 >> 1
	PrintChar('\n')

}
for k := 0; k < constants.COLUMN_LENGTH; k++ {
	fmt.Printf(constants.STR_BCD_TIME, clockBCD[k])
}
PrintChar('\n')
PrintChar('\n')
}
