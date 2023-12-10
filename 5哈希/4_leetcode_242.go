package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for j := 0; j < len(t); j++ {
		if _, ok := m[t[j]]; ok {
			m[t[j]]--
			if m[t[j]] == 0 {
				delete(m, t[j])
			}
		}
	}

	return len(m) == 0
}
