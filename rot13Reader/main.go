package main

import (
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(n []byte) (int, error) {
	buffer := make([]byte, len(n))
	nBytesRead, err := r13.r.Read(buffer)
	if err != nil {
		return 0, err
	}

	for i := 0; i < nBytesRead; i++ {
		decryptedChar := buffer[i] - 13
		if !isValid13Rot(buffer[i], decryptedChar) {
			decryptedChar = buffer[i] + 13
		}
		if !isValid13Rot(buffer[i], decryptedChar) {
			decryptedChar = buffer[i]
		}
		n[i] = decryptedChar
	}

	return len(n), nil
}

func isValid13Rot(origin byte, encrypted byte) bool {
	runeOrigin, runeEncrypted := rune(origin), rune(encrypted)
	if !unicode.IsLetter(runeOrigin) || !unicode.IsLetter(runeEncrypted) {
		return false
	}

	if unicode.IsUpper(runeOrigin) {
		return unicode.IsUpper(runeEncrypted)
	}

	return unicode.IsLower(runeEncrypted)
}

func main() {
	// TODO: Look for other solutions for this
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	if _, err := io.Copy(os.Stdout, &r); err != nil {
		log.Fatal(err)
	}
}
