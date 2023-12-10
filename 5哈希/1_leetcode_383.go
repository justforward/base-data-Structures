package main

func canConstruct(ransomNote string, magazine string) bool {

	m := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if _, ok := m[ransomNote[i]]; !ok {
			return false
		} else {
			m[ransomNote[i]]--
			if m[ransomNote[i]] == 0 {
				delete(m, ransomNote[i])
			}
		}
	}
	return true
}
