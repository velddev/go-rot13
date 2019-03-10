package main;

import "fmt"
import "strings"

const base = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	fmt.Println(rot13("abc"));
}

func rot13(text string) string {
	var modified = ""
	for i := 0; i < len(text); i++ {
		nextIdx := strings.Index(base, string(text[i])) + 13;
		if nextIdx >= len(base) {
			nextIdx -= len(base);
		}
		modified += string(base[nextIdx]);
	}
	return modified;
}