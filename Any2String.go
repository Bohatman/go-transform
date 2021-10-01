package go_transform

import "strconv"

func BoolToString(input bool) string {
	return strconv.FormatBool(input)
}
func IntToString(input int) string {
	return strconv.FormatInt(int64(input), 10)
}
func UintToString(input uint) string {
	return strconv.FormatUint(uint64(input), 10)
}
