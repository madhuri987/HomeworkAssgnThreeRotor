package main

import (
	"fmt"
	"strings"
)

type rotor struct {
	wiring   string
	position int
}

var reflectorB = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	rotorI := rotor{wiring: "EKMFLGDQVZNTOWYHXUSPAIBRCJ"}
	rotorII := rotor{wiring: "AJDKSIRUXBLHWTMCQGZNPYFVOE"}
	rotorIII := rotor{wiring: "BDFHJLCPRTXVZNYEIWGAKMUSQO"}

	plugboard := map[rune]rune{
		'A': 'E', 'B': 'J', 'C': 'M', 'D': 'Z', 'E': 'A',
		// Add more plugboard connections as needed
	}

	plaintext := "HELLO"
	encryptedText := enigmaEncrypt(plaintext, plugboard, rotorI, rotorII, rotorIII)
	fmt.Println("Plaintext: ", plaintext)
	fmt.Println("Encrypted Text: ", encryptedText)

	decryptedText := enigmaDecrypt(encryptedText, plugboard, rotorI, rotorII, rotorIII)
	fmt.Println("Decrypted Text: ", decryptedText)
}

func enigmaEncrypt(plaintext string, plugboard map[rune]rune, rotors ...rotor) string {
	plaintext = strings.ToUpper(plaintext)
	var encrypted strings.Builder

	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			// Apply plugboard transformation
			if plug, ok := plugboard[char]; ok {
				char = plug
			}

			// Pass the character through the rotors from right to left
			char = substitute(char, rotors[2])
			char = substitute(char, rotors[1])
			char = substitute(char, rotors[0])

			// Pass the character through the reflector
			char = reflector(char)

			// Pass the character through the rotors from left to right
			char = substitute(char, rotors[0])
			char = substitute(char, rotors[1])
			char = substitute(char, rotors[2])

			// Apply plugboard transformation
			if plug, ok := plugboard[char]; ok {
				char = plug
			}

			encrypted.WriteRune(char)
		} else {
			// Non-alphabetic characters are not modified
			encrypted.WriteRune(char)
		}
	}

	return encrypted.String()
}

func enigmaDecrypt(plaintext string, plugboard map[rune]rune, rotors ...rotor) string {
	plaintext = strings.ToUpper(plaintext)
	var encrypted strings.Builder

	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			// Apply plugboard transformation
			if plug, ok := plugboard[char]; ok {
				char = plug
			}

			char = decrypt(char, rotors[2])
			char = decrypt(char, rotors[1])
			char = decrypt(char, rotors[0])

			char = reflector(char)

			char = decrypt(char, rotors[0])
			char = decrypt(char, rotors[1])
			char = decrypt(char, rotors[2])

			// Apply plugboard transformation
			if plug, ok := plugboard[char]; ok {
				char = plug
			}

			encrypted.WriteRune(char)
		} else {
			// Non-alphabetic characters are not modified
			encrypted.WriteRune(char)
		}
	}

	return encrypted.String()
}

func substitute(char rune, rotor rotor) rune {
	index := int(char-'A') % 26      //zero-index character position in the A-Z
	return rune(rotor.wiring[index]) //Find that character at the index position on the rotor
}

func decrypt(char rune, rotor rotor) rune {
	index := strings.IndexRune(rotor.wiring, char) % 26 //zero-index character position rotor
	return rune(alphabet[index])                        //Find the character position in the A-Z
}

func reflector(char rune) rune {
	index := strings.IndexRune(reflectorB, char)
	return rune(alphabet[index])
}
