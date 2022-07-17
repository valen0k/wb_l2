package dev02

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Unpacking(str string) (string, error) {
	if _, err := strconv.Atoi(str); err == nil {
		return "", errors.New("некорректная строка")
	}

	var bufRune rune
	var bufStr string
	var isEscape bool
	buf := bytes.Buffer{}

	for _, c := range str {
		if unicode.IsDigit(c) && !isEscape {
			bufStr = strings.Repeat(string(bufRune), int(c-'0')-1)
			buf.WriteString(bufStr)
		} else {
			if c == '\\' && bufRune != '\\' {
				isEscape = true
			} else {
				isEscape = false
			}
			if !isEscape {
				buf.WriteRune(c)
			}
			bufRune = c
		}
	}
	return buf.String(), nil
}
