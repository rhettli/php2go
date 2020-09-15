package php

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// Md5 - Calculate the md5 hash of a string
func Md5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// Md5File - Calculates the md5 hash of a given file
func Md5File(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic("open file for md5 fail:==" + filename + err.Error())
		return ""
	}
	defer file.Close()

	md5_ := md5.New()

	//_, err := io.Copy(md5_, file)

	if _, err := io.Copy(md5_, file); err == nil {
		MD5Str := hex.EncodeToString(md5_.Sum(nil))
		//log.Println("md5_:", MD5Str)
		return MD5Str
	} else {
		panic("get md5_ err:==" + filename + err.Error())
	}
	return ""
}
