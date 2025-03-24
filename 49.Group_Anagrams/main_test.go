package main

import (
	"reflect"
	"slices"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	words := make(map[string][]string)
	for _, str := range strs {
		byteRepresentation := []byte(str)
		slices.Sort(byteRepresentation)
		words[string(byteRepresentation)] = append(words[string(byteRepresentation)], str)
	}
	answer := make([][]string, 0)
	for _, value := range words {
		answer = append(answer, value)
	}
	return answer
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			"base test",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			"one empty string test",
			[]string{"a"},
			[][]string{{"a"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := groupAnagrams(test.input)
			if reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test failed, expected %v, got %v", test.expected, result)
			}
		})
	}
}
