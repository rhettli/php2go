package php

import "fmt"

func InArray(i interface{}, arr []interface{}) bool {
	var key = i.(string)
	for m := 0; m < len(arr); m++ {
		if fmt.Sprintf("%v", key) == fmt.Sprintf("%v", arr[m]) {
			return true
		}
	}
	return false
}
