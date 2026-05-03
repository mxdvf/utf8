package utf8

import (
	"fmt"
	"slices"
	"testing"
)

func TestEncodeCopyrightSymbol(t *testing.T) {
	// U+00A9 is the © symbol
	got, _ := EncodeRune(0x00A9)
	expected := []byte{194, 169}

	if !slices.Equal(got, expected) {
		t.Fatalf("\nbytes mismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestEncodeEuroSymbol(t *testing.T) {
	// U+20AC is the € symbol
	got, _ := EncodeRune(0x20AC)
	expected := []byte{226, 130, 172}

	if !slices.Equal(got, expected) {
		t.Fatalf("\nbytes mismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestEncodeSmileyEmojiSymbol(t *testing.T) {
	// U+1F600 is the 😀 symbol
	got, _ := EncodeRune(0x1F600)
	expected := []byte{240, 159, 152, 128}

	if !slices.Equal(got, expected) {
		t.Fatalf("\nbytes mismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestEncodeFullString(t *testing.T) {
	got := EncodeString("©€😀")
	expected := []byte{194, 169, 226, 130, 172, 240, 159, 152, 128}

	if !slices.Equal(got, expected) {
		t.Fatalf("\nbytes mismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestEncodeCustomGlyph(t *testing.T) {
	got, _ := EncodeRune(0xE001)
	expected := []byte{238, 128, 129}

	fmt.Println("\uE003")

	if !slices.Equal(got, expected) {
		t.Fatalf("\nbytes mismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestDecodeCopyrightSymbol(t *testing.T) {
	got, _ := DecodeRune([]byte{194, 169})
	expected := rune('©')

	if got != rune(expected) {
		t.Fatalf("\nmismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestDecodeEuroSymbol(t *testing.T) {
	got, _ := DecodeRune([]byte{226, 130, 172})
	expected := rune('€')

	if got != rune(expected) {
		t.Fatalf("\nmismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestDecodeSmileyEmojiSymbol(t *testing.T) {
	got, _ := DecodeRune([]byte{240, 159, 152, 128})
	expected := rune('😀')

	if got != rune(expected) {
		t.Fatalf("\nmismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestDecodeFullBytes(t *testing.T) {
	got := DecodeBytes([]byte{194, 169, 226, 130, 172, 240, 159, 152, 128})
	expected := []rune{'©', '€', '😀'}

	if !slices.Equal(got, expected) {
		t.Fatalf("\nmismatch:\n got: %v\n expected: %v\n", got, expected)
	}
}

func TestRoundtrip(t *testing.T) {
	input := "©€😀"

	bytes := EncodeString(input)
	runes := DecodeBytes(bytes)

	output := string(runes)

	if input != output {
		t.Fatalf("\nmismatch:\n got: %v\n expected: %v\n", output, input)
	}
}

func BenchmarkEncodeRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeRune('€')
	}
}

func BenchmarkDecodeRune(b *testing.B) {
	input := []byte{0xE2, 0x82, 0xAC}
	for i := 0; i < b.N; i++ {
		DecodeRune(input)
	}
}
