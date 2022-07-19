package dev04

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVerifyAnagrams(t *testing.T) {
	testCases := []string{"пятка", "пятак", "тяПка", "листок", "слиток", "столик", "столик", "тоР"}
	result := map[string][]string{
		"пятка":  []string{"пятак", "пятка", "тяпка"},
		"листок": []string{"листок", "слиток", "столик"},
	}

	anagrams := VerifyAnagrams(&testCases)
	for key, value := range *anagrams {
		_, ok := result[key]
		if !ok {
			t.Errorf("key: %s = %v\n", key, ok)
		}
		if !reflect.DeepEqual(*value, result[key]) {
			fmt.Printf("value: %v != %v\n", *value, result[key])
		}
	}
}
