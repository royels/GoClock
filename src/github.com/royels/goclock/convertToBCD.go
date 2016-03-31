package goclock


import "github.com/royels/constants"
//import "fmt"
func ConvertToBCD(clockDecimal [3] int, clockBCD *[3]int) {
	ch := make(chan int)
	go craftBCD(clockDecimal[constants.HR_INDEX],  ch)
	go craftBCD(clockDecimal[constants.MIN_INDEX],  ch)
	go craftBCD(clockDecimal[constants.SEC_INDEX],  ch)
	a, b, c := <-ch, <-ch, <-ch
	clockBCD[constants.HR_INDEX]  = a
	clockBCD[constants.MIN_INDEX] = b
	clockBCD[constants.SEC_INDEX] = c
	

}

func craftBCD(src int, ch chan int) {
	rem := src % 10
	div := src / 10
	temp1 := div << 4
	temp2 := temp1|rem 
	ch <- temp2
}
