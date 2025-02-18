package main

func canConstruct(ransomNote string, magazine string) bool {
	countedMagazineLetters := map[rune]int{}
	for i := 0; i < len(magazine); i++ {
		countedMagazineLetters[rune(magazine[i])] += 1
	}
	for i := 0; i < len(ransomNote); i++ {
		if countedMagazineLetters[rune(ransomNote[i])] == 0 {
			return false
		}
		countedMagazineLetters[rune(ransomNote[i])] -= 1
	}
	return true
}
