package problem278

func canConstruct(ransomNote string, magazine string) bool {
	letterMap := make(map[byte]int, 0)
	for i := 0; i < len(magazine); i++ {
		letterMap[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		b := ransomNote[i]
		if n, ok := letterMap[b]; ok {
			if n > 0 {
				letterMap[b]--
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
