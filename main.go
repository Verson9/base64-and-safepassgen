package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	base64chars    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	capitalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	smallLetters   = "abcdefghijklmnopqrstuvwxyz"
	numbers        = "0123456789"
	specialChars   = "!@#$%^&*()"
	min            = 8
	max            = 16
)

var validGeneratedPass string

type Base64 struct {
	encode    [64]byte
	decodeMap [256]byte
	strict    bool
}

func NewBase64() *Base64 {
	e := new(Base64)
	copy(e.encode[:], base64chars)

	for i := 0; i < len(e.decodeMap); i++ {
		e.decodeMap[i] = 0xFF
	}
	for i := 0; i < len(base64chars); i++ {
		e.decodeMap[base64chars[i]] = byte(i)
	}
	return e
}

func (enc *Base64) encodeToBase64(dst, src []byte) {
	if len(src) == 0 {
		return
	}
	_ = enc.encode
	di, si := 0, 0
	n := (len(src) / 3) * 3
	for si < n {
		val := uint(src[si+0])<<16 | uint(src[si+1])<<8 | uint(src[si+2])

		dst[di+0] = enc.encode[val>>18&0x3F]
		dst[di+1] = enc.encode[val>>12&0x3F]
		dst[di+2] = enc.encode[val>>6&0x3F]
		dst[di+3] = enc.encode[val&0x3F]

		si += 3
		di += 4
	}
	remain := len(src) - si
	if remain == 0 {
		return
	}
	val := uint(src[si+0]) << 16
	if remain == 2 {
		val |= uint(src[si+1]) << 8
	}
	dst[di+0] = enc.encode[val>>18&0x3F]
	dst[di+1] = enc.encode[val>>12&0x3F]
	switch remain {
	case 2:
		dst[di+2] = enc.encode[val>>6&0x3F]
		dst[di+3] = byte('=')
	case 1:
		dst[di+2] = byte('=')
		dst[di+3] = byte('=')
	}
}

func (enc *Base64) encodeString(src []byte) string {
	buf := make([]byte, (len(src)+2)/3*4)
	enc.encodeToBase64(buf, src)
	return string(buf)
}

func base64gen() {
	fmt.Println(`
Now we will encode generated password into base64`)
	fmt.Println("Text to encode:" + validGeneratedPass)
	buf := make([]byte, (len(validGeneratedPass)+2)/3*4)
	NewBase64().encodeToBase64(buf, []byte(validGeneratedPass))
	fmt.Println(string(buf))
}

func safePassGen() {
	fmt.Println("First we will generate password")
	passLen := rand.Intn(max-min) + min
	rand.Seed(time.Now().UnixNano())

	chars := []rune(capitalLetters + smallLetters + numbers + specialChars)
	var gen strings.Builder
	for i := 0; i < passLen; i++ {
		gen.WriteRune(chars[rand.Intn(len(chars))])
	}
	validGeneratedPass = makeValid(gen.String())
	fmt.Println(`Generated Password: ` + validGeneratedPass)
	base64gen()
}

func makeValid(generatedPass string) string {
	var notChanged bool
	switch {
	case !strings.ContainsAny(generatedPass, capitalLetters):
		generatedPass = addSpecificChar(generatedPass, capitalLetters)
		break
	case !strings.ContainsAny(generatedPass, smallLetters):
		generatedPass = addSpecificChar(generatedPass, smallLetters)
		break
	case !strings.ContainsAny(generatedPass, numbers):
		generatedPass = addSpecificChar(generatedPass, numbers)
		break
	case !strings.ContainsAny(generatedPass, specialChars):
		generatedPass = addSpecificChar(generatedPass, specialChars)
		break
	default:
		notChanged = true
	}
	if !notChanged {
		return makeValid(generatedPass)
	}
	return generatedPass
}

func addSpecificChar(generatedPass string, chars string) string {
	intn := rand.Intn(len(generatedPass))
	runes := []rune(generatedPass)
	runes[intn] = rune(chars[rand.Intn(len(chars))])
	return string(runes)
}

func menu() {
	fmt.Println(` 
Sorry but i just become aware of the fact that go playground cannot read from the console online :(`)
	safePassGen()
}

func main() {
	menu()
}
