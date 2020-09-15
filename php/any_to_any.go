package php

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// any to ten or ten to any
func AnyToAny(numi interface{}, n int, tenToAny bool) (interface{}, error) {
	if n > 64 {
		return "", errors.New("max to small than 64")
	}
	var tenToAnyMap map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z",
		36: "A", 37: "B", 38: "C", 39: "D", 40: "E", 41: "F", 42: "G", 43: "H", 44: "I", 45: "J", 46: "K", 47: "L", 48: "M", 49: "N", 50: "O", 51: "P",
		52: "Q", 53: "R", 54: "S", 55: "T", 56: "U", 57: "V", 58: "W", 59: "X", 60: "Y", 61: "Z", 62: "*", 63: "/", 64: "="}
	if tenToAny {
		newNumStr := ""
		var remainder int
		var remainderString string
		num := numi.(int)
		for num != 0 {
			remainder = num % n
			if 76 > remainder && remainder > 9 {
				remainderString = tenToAnyMap[remainder]
			} else {
				remainderString = strconv.Itoa(remainder)
			}
			newNumStr = remainderString + newNumStr
			num = num / n
		}
		return newNumStr, nil
	} else {
		// map根据value找key
		findKey := func(in string) int {
			result := -1
			for k, v := range tenToAnyMap {
				if in == v {
					result = k
				}
			}
			return result
		}

		// 任意进制转10进制
		num := numi.(string)
		var newNum float64
		newNum = 0.0
		nNum := len(strings.Split(num, "")) - 1
		for _, value := range strings.Split(num, "") {
			tmp := float64(findKey(value))
			if tmp != -1 {
				newNum = newNum + tmp*math.Pow(float64(n), float64(nNum))
				nNum = nNum - 1
			} else {
				break
			}
		}
		return int(newNum), nil
	}

}
