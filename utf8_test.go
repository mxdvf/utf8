package utf8

import (
	"testing"
)

func TestEncodeCopyrightSymbol(t *testing.T) {
	// U+00A9 is the © symbol
	got, _ := EncodeRune(0x00A9)
	expected := []byte{194, 169}

	if !isBytesEqual(got, expected) {
		t.Fatalf("\nbytes mismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestEncodeEuroSymbol(t *testing.T) {
	// U+20AC is the € symbol
	got, _ := EncodeRune(0x20AC)
	expected := []byte{226, 130, 172}

	if !isBytesEqual(got, expected) {
		t.Fatalf("\nbytes mismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestEncodeSmileyEmojiSymbol(t *testing.T) {
	// U+1F600 is the 😀 symbol
	got, _ := EncodeRune(0x1F600)
	expected := []byte{240, 159, 152, 128}

	if !isBytesEqual(got, expected) {
		t.Fatalf("\nbytes mismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestEncodeFullString(t *testing.T) {
	got := EncodeString("©€😀")
	expected := []byte{194, 169, 226, 130, 172, 240, 159, 152, 128}

	if !isBytesEqual(got, expected) {
		t.Fatalf("\nbytes mismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestDecodeSimple(t *testing.T) {
	got, _ := DecodeRune([]byte{194, 169})
	expected := rune('©')

	if got != rune(expected) {
		t.Fatalf("\nmismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func isBytesEqual(s1, s2 []byte) bool {
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
