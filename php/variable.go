package php

import (
	"fmt"
	"reflect"
	"strconv"
)

// Boolval - Get the boolean value of a variable
func Boolval(val interface{}) bool {

	switch v := val.(type) {
	case int, int8, int16, int32, int64:
		if v != 0 {
			return true
		}
		return false
	case uint, uint8, uint16, uint32, uint64:
		if v != 0 {
			return true
		}
		return false
	case bool:
		return v
	case complex64, complex128:
		if v != complex128(0) {
			return true
		}
		return false
	case float32, float64:
		if v != float64(0) {
			return true
		}
	default:
		return v == nil
	}

	return false
}

// Intval - Get the integer value of a variable
func Intval(v interface{}) (int, error) {

	return strconv.Atoi(fmt.Sprintf("%v", v))
}

// Strval - Get string value of a variable
func Strval(val interface{}) string {

	return fmt.Sprintf("%v", val)
}

// Gettype - Get the type of a variable
func Gettype(v interface{}) string {

	t := reflect.TypeOf(v)

	return t.String()
}

// IsBool - Finds out whether a variable is a boolean
func IsBool(v interface{}) bool {

	//_, ok := v.(bool)
	//
	//return ok
	return v == true || v == false
}