package main

import (
	"fmt"
)

func main() {
	a, b := "аsьвыs", "ыавьss"
	fmt.Println(isAnagram(a, b))
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1b, s2b := []byte(s1), []byte(s2)
	for j := 1; j < len(s1b); j++ {
		key := s1b[j]
		i := j - 1
		for i >= 0 && s1b[i] > key {
			s1b[i+1] = s1b[i]
			i--
		}
		s1b[i+1] = key
	}
	for j := 1; j < len(s2b); j++ {
		key := s2b[j]
		i := j - 1
		for i >= 0 && s2b[i] > key {
			s2b[i+1] = s2b[i]
			i--
		}
		s2b[i+1] = key
	}
	return string(s1b) == string(s2b)
}
