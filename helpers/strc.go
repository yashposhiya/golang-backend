package helpers

import(
	"strconv"
)

func StringToUint(val string) (uint,error){
	res,err := strconv.ParseUint(val,10,64)
	return uint(res),err
}

func StringToInt(val string) (int, error){
	res, err := strconv.ParseInt(val,10,64);
	return int(res),err
}