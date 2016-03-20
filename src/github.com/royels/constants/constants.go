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

const(

STR_USAGE             = "Usage: %s startTime [intervalSeconds [numTicks]]\n    (required startTime = HH:MM:SS)\n    (HH = [%d-%d]; MM = [%d-%d]; SS = [%d-%d])\n    (optional intervalSeconds = [%d-%d], default = %d)\n    (optional numTicks = [%d-%d], default = %d)\n\n"
STR_BCD_TIME          = "%02x "
STR_ERR_TICK_RANGE    = "\n\t# of clock ticks must be in the range [%d-%d]\n\n"

STR_ERR_TICK_VALUE    = "\n\tError parsing clock ticks\n\n"

STR_ERR_INTERVAL_RANGE= "\n\tNumber of interval seconds must be in the range [%d-%d]\n\n"

STR_ERR_INTERVAL_VALUE= "\n\tError parsing interval seconds\n\n"

STR_ERR_HR_RANGE      = "\n\tHours value must be in the range [%d-%d]\n\n"

STR_ERR_HR_VALUE      = "\n\tError parsing hours\n\n"

STR_ERR_MIN_RANGE     = "\n\tMinutes value must be in the range [%d-%d]\n\n"

STR_ERR_MIN_VALUE     = "\n\tError parsing minutes\n\n"

STR_ERR_SEC_RANGE     = "\n\tSeconds value must be in the range [%d-%d]\n\n"

STR_ERR_SEC_VALUE     = "\n\tError parsing seconds\n\n"

STR_ERR_TIME_FORMAT   = "\n\tStarting time format incorrect\n\n"
) 
