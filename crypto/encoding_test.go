package crypto

import "testing"

func TestBase64Encoding(t *testing.T) {
	encoded := Base64Encoding("Hello World!")
	expected := "SGVsbG8gV29ybGQh"

	if encoded != expected {
		t.Errorf("Expected %v but got %v", expected, encoded)
	}
}

func TestBase64Decoding(t *testing.T) {
	decoded := Base64Decoding("SGVsbG8gV29ybGQh")
	expected := "Hello World!"

	if decoded != expected {
		t.Errorf("Expected %v but got %v", expected, decoded)
	}
}