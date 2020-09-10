package main

var str = "abchdbasbcaheaeijdaibcasd"


func frequencyByMap() rune {
	fm := make(map[rune]int)
	max := 0
	var mr rune

	for _, r := range str {
		fm[r]++
		if i, _ := fm[r]; i > max {
			max = i
			mr = r
		}
	}

	return mr
}

func frequencyByArray() rune {
	var fm [26]int
	max := 0
	var mr rune

	for _, r := range str {
		fm[r - 'a']++
		if i := fm[r - 'a']; i > max {
			max = i
			mr = r
		}
	}

	return mr
}
