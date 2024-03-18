package utils

import "strconv"

func StringToInt64(str string) int64 {
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return id
}

func StringToUint64(str string) uint64 {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return id
}
