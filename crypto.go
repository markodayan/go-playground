package main

import "encoding/base64"

func base64Encoding(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func base64Decoding(encrypted string) string {
	decoded, err := base64.RawStdEncoding.DecodeString(encrypted)

	if err != nil {
		panic(err)
	}

	return string(decoded)
}