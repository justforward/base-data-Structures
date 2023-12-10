package main

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	w2p := map[string]byte{}
	p2w := map[byte]string{}

	if len(pattern) != len(words) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		p, w := pattern[i], words[i]
		if a, ok := p2w[p]; ok {
			if a != w {
				return false
			}
		}

		if a, ok := w2p[w]; ok {
			if a != p {
				return false
			}
		}

		w2p[w] = p
		p2w[p] = w
	}
	return true
}
