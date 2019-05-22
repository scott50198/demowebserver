package utils

import "time"

func GetNowTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
