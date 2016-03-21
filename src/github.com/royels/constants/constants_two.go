package constants

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
