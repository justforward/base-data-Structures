package main

func isAnagram(s string, t string) bool {
	m := [26]int{}
	for _, ss := range s {
		m[ss-'a']++
	}

	for _, tt := range t {
		m[tt-'a']--
	}

	return m == [26]int{} // go语言中，数组的比较是逐个元素进行的。即每个数组的对应元素都会被依次比较
}

func main() {
	println(isAnagram("sbc", "cbd"))
}
