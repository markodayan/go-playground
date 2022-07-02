package crypto

import "encoding/base64"

func Base64Encoding(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func Base64Decoding(encrypted string) string {
	decoded, err := base64.RawStdEncoding.DecodeString(encrypted)

	if err != nil {
		panic(err)
	}

	return string(decoded)
}