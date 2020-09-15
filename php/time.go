package php

import "time"

// Time - Return current Unix timestamp
func Time(micro ...bool) int64 {
	if len(micro) > 0 && micro[0] {
		return time.Now().UnixNano() / 1e6
	}
	return time.Now().Unix()
}
