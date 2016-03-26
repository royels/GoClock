package constants 


const(
		MIN_ARGS = 2
		MAX_ARGS = 4

		MIN_HR= 0
		MAX_HR= 24
		HR_OFFSET= 23

		MIN_MINUTE= 0
		MAX_MINUTE =60
		MINUTE_OFFSET= 59

		MIN_SEC= 0
		MAX_SEC =60
		SEC_OFFSET= 59

		HR_INDEX= 0
		MIN_INDEX= 1
		SEC_INDEX= 2

		BASE= 10
		MIN_INTERVAL= 1
	MAX_INTERVAL =86399
INTERVAL_OFFSET= 86398
DEF_INTERVAL= 1

MIN_TICKS= 0
MAX_TICKS= 61
DEF_TICKS= 7

TIME_INDEX= 1
INTERVAL_INDEX= 2
TICKS_INDEX= 3

COLUMN_LENGTH= 3
ROW_LENGTH =4
LIST_LENGTH =3

REQ_SEPARATORS =2

TIME_STR_DELIMITER =':'

FIRST_BITMASK= 0x80
SECOND_BITMASK= 0x08


LIT_LED ='O'
UNLIT_LED= '.'


ERR_TIME_FORMAT=  0x1 /* Invalid start time string */
ERR_HR_VALUE   =  0x2 /* Hour value invalid */
ERR_HR_RANGE   =  0x4 /* Hour value out of range */
ERR_MIN_VALUE  =  0x8 /* Min value invalid */
ERR_MIN_RANGE  = 0x10 /* Min value out of range */
ERR_SEC_VALUE  = 0x20 /* Sec value invalid */
ERR_SEC_RANGE  = 0x40 /* Sec value out of range */


)

type CategoryInfo struct {
	MIN_VAL int
	OFFSET int
	ERR_VAL int
	ERR_RANGE int	
}

var (
	HR = &CategoryInfo{MIN_HR, HR_OFFSET, ERR_HR_VALUE, ERR_HR_RANGE}
	MIN = &CategoryInfo{MIN_MINUTE, MINUTE_OFFSET, ERR_MIN_VALUE, ERR_MIN_RANGE}
	SEC = &CategoryInfo{MIN_SEC, SEC_OFFSET, ERR_SEC_VALUE, ERR_SEC_RANGE}
)



