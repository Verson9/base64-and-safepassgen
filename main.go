package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	capitalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	smallLetters   = "abcdefghijklmnopqrstuvwxyz"
	numbers        = "0123456789"
	specialChars   = "!@#$%^&*()"
	min            = 8
	max            = 16
)

func base64gen(reader *bufio.Reader) {
	fmt.Println("Type text to encode:")
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	encoding := base64.StdEncoding.EncodeToString([]byte(readString[:len(readString)-1]))
	fmt.Println(encoding)
	menu()
}

func safePassGen() {
	passLen := rand.Intn(max-min) + min
	rand.Seed(time.Now().UnixNano())

	chars := []rune(capitalLetters + smallLetters + numbers + specialChars)
	var gen strings.Builder
	for i := 0; i < passLen; i++ {
		gen.WriteRune(chars[rand.Intn(len(chars))])
	}
	validGeneratedPass := makeValid(gen.String())
	fmt.Println(`Generated Password: ` + validGeneratedPass)
	menu()
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(` 
Choose option by typing number:
1. Generate base64 from text
2. Generate safe password`)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	switch readString[0] {
	case '1':
		base64gen(reader)
	case '2':
		safePassGen()
	default:
		fmt.Println("Wrong input...\n")
		menu()
	}

}
func main() {
	menu()
}
