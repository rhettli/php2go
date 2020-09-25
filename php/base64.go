package php

import "encoding/base64"

// Base64Encode - Encodes data with MIME base64
func Base64Encode(s []byte) string {
	return base64.StdEncoding.EncodeToString(s)
}

// Base64Decode - Decodes data encoded with MIME base64
func Base64Decode(s string) ([]byte, error) {
	bs, err := base64.StdEncoding.DecodeString(s)
	return bs, err
}
