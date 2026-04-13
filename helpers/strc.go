package helpers

import(
	"strconv"
)

func StringToUint(val string) (uint,error){
	res,err := strconv.ParseUint(val,10,64)
	return uint(res),err
}