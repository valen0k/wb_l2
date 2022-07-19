package dev04

import (
	"sort"
	"strings"
)

func VerifyAnagrams(words *[]string) *map[string]*[]string {
	var bufRune RuneSlice
	bufMap := make(map[string][]string)
	//пробегаемся по словам из слайса, делаем нижний регистр, и разделяем на группы
	for _, word := range *words {
		word = strings.ToLower(word)
		bufRune = []rune(word)
		sort.Sort(bufRune)
		bufMap[string(bufRune)] = append(bufMap[string(bufRune)], word)
	}

	result := make(map[string]*[]string)
	//пробегаемся по группам, сортируем слайс и записываем в конечное множество, если там не одно слово
	for _, value := range bufMap {
		bufSlice := newSet(value)
		if len(bufSlice) != 1 {
			sort.Strings(bufSlice)
			result[value[0]] = &bufSlice
		}
	}

	return &result
}

func newSet(words []string) []string {
	buf := make(map[string]struct{})
	set := make([]string, 0)

	for _, word := range words {
		if _, ok := buf[word]; !ok {
			buf[word] = struct{}{}
			set = append(set, word)
		}
	}

	return set
}
