package utf8

import (
	"slices"
)

/* UTF8 spec:
Codepoint range         Byte 1      Byte 2      Byte 3      Byte 4
U+0000   – U+007F      0xxxxxxx
U+0080   – U+07FF      110xxxxx    10xxxxxx
U+0800   – U+FFFF      1110xxxx    10xxxxxx    10xxxxxx
U+10000  – U+10FFFF    11110xxx    10xxxxxx    10xxxxxx    10xxxxxx
*/

func DecodeBytes(b []byte) []rune {
	start, end := 0, 1
	runes := []rune{}
	for start < len(b) {
		r, x := DecodeRune(b[start:end])
		if r == -1 {
			end += x - 1
			continue
		}
		runes = append(runes, r)
		start = end
		end = start + 1
	}
	return runes
}

func DecodeRune(b []byte) (rune, int) {
	magic := b[0]
	switch {
	case magic&0x80 == 0:
		return decodeOne(b)
	case magic&0xE0 == 0xC0:
		return decodeTwo(b)
	case magic&0xF0 == 0xE0:
		return decodeThree(b)
	default:
		return decodeFour(b)
	}
}

func decodeOne(b []byte) (rune, int) {
	return rune(b[0]), 1
}

func decodeTwo(b []byte) (rune, int) {
	if len(b) != 2 {
		return -1, 2
	}

	var r rune
	r = rune(b[0]) & 0x1F
	r = r << 6
	r = r | (rune(b[1] & 0x3F))
	return r, 2
}

func decodeThree(b []byte) (rune, int) {
	if len(b) != 3 {
		return -1, 3
	}

	var r rune
	r = rune(b[0]) & 0xF
	r = r << 6
	r = r | (rune(b[1] & 0x3F))
	r = r << 6
	r = r | (rune(b[2] & 0x3F))
	return r, 3
}

func decodeFour(b []byte) (rune, int) {
	if len(b) != 4 {
		return -1, 4
	}

	var r rune
	r = rune(b[0]) & 0x7
	r = r << 6
	r = r | (rune(b[1] & 0x3F))
	r = r << 6
	r = r | (rune(b[2] & 0x3F))
	r = r << 6
	r = r | (rune(b[3] & 0x3F))
	return r, 4
}

func EncodeString(s string) []byte {
	raw := []byte(s)
	result := []byte{}
	start, end := 0, 1
	for start < len(raw) {
		r, x := DecodeRune(raw[start:end])
		if r == -1 {
			end += x - 1
			continue
		}
		encoded, _ := EncodeRune(r)
		result = append(result, encoded...)
		start = end
		end = start + 1
	}
	return result
}

// TODO: keep this implementation so we could improve the perf for our encoder
// and then benchmark it with the for loop
// func EncodeString(s string) []byte {
// 	var finalBytes []byte
// 	for _, r := range s {
// 		bytes, _ := EncodeRune(r)
// 		finalBytes = append(finalBytes, bytes...)
// 	}
// 	return finalBytes
// }

func EncodeRune(r rune) ([]byte, int) {
	switch {
	case r <= 0x7F:
		return encodeOne(r)
	case r <= 0x7FF:
		return encodeTwo(r)
	case r <= 0xFFFF:
		return encodeThree(r)
	default:
		return encodeFour(r)
	}
}

func encodeOne(r rune) ([]byte, int) {
	return []byte{byte(r)}, 1
}

func encodeTwo(r rune) ([]byte, int) {
	bytes := make([]byte, 0, 2)
	// Extract the first byte from right side
	result := r & 0x3F     // 0x3F = 0b00111111 (masking 6 bits of rune r)
	result = 0x80 | result // 0x80 = 0b10000000 (merging the extracted 6 bits as per spec)
	bytes = slices.Insert(bytes, 0, byte(result))
	// Right shift by 6 bits
	r = r >> 6
	// Extract the last byte
	result = r & 0x1F      // 0x3F = 0b00011111
	result = 0xC0 | result // 0xC0 = 0b11000000
	bytes = slices.Insert(bytes, 0, byte(result))
	// Return the spec-compliant bytes
	return bytes, 2
}

func encodeThree(r rune) ([]byte, int) {
	bytes := make([]byte, 0, 3)
	extract := r & 0x3F
	result := 0x80 | extract
	bytes = slices.Insert(bytes, 0, byte(result))
	r = r >> 6
	extract = r & 0x3F
	result = 0x80 | extract
	bytes = slices.Insert(bytes, 0, byte(result))
	r = r >> 6
	extract = r & 0xF       // 0xF 	= 0b00001111
	result = 0xE0 | extract // 0xE0 = 0b11100000
	bytes = slices.Insert(bytes, 0, byte(result))
	return bytes, 3
}

func encodeFour(r rune) ([]byte, int) {
	bytes := make([]byte, 0, 4)
	extract := r & 0x3F
	result := 0x80 | extract
	bytes = slices.Insert(bytes, 0, byte(result))
	r = r >> 6
	extract = r & 0x3F
	result = 0x80 | extract
	bytes = slices.Insert(bytes, 0, byte(result))
	r = r >> 6
	extract = r & 0x3F
	result = 0x80 | extract
	bytes = slices.Insert(bytes, 0, byte(result))
	r = r >> 6
	extract = r & 0x7       // 0x7 	= 0b00000111
	result = 0xF0 | extract // 0xF0 = 0b11110000
	bytes = slices.Insert(bytes, 0, byte(result))
	return bytes, 4
}
