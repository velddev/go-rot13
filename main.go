package main

import (
	"fmt"
	"os"
	"strings"
)

const base = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ !?.,"

func main() {
	if len(os.Args) == 1 {
		usage()
		return
	}

	switch os.Args[1] {
	case "decode":
		decode()
	case "encode":
		encode()
	default:
		usage()
	}
}

func usage() {
	fmt.Print("usage:\n -> go-rot13 encode <text>\n -> go-rot13 decode <text>")
}

func decode() {
	for i := 2; i < len(os.Args); i++ {
		fmt.Print(FromRot13(os.Args[i]))
		fmt.Print(" ")
	}
}

func encode() {
	for i := 2; i < len(os.Args); i++ {
		fmt.Print(ToRot13(os.Args[i]))
		fmt.Print(" ")
	}
}

// FromRot13 decodes rot13 string
func FromRot13(text string) string {
	return ShiftRot(text, -13)
}

// ToRot13 encodes rot13 string
func ToRot13(text string) string {
	return ShiftRot(text, 13)
}

// ShiftRot shifts string characters from base by n
func ShiftRot(text string, n int) string {
	var modified = ""
	for i := 0; i < len(text); i++ {
		nextIdx := strings.Index(base, string(text[i])) + n
		if nextIdx >= len(base) {
			nextIdx -= len(base)
		}
		modified += string(base[nextIdx])
	}
	return modified
}
