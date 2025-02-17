package linklist

func MinString(source, includeString string) string {
	needMap := make(map[byte]int, 0)
	for i := 0; i < len(includeString); i++ {
		needMap[includeString[i]]++
	}

	windowMap := make(map[byte]int, 0)
	left, right := 0, 0
	// 判断 字符出现的次数，如果在两个map出现一样多则 ++
	valid := 0
	for right < len(source) {
		c := source[right]
		// c 是要判断的字符
		if _, ok := needMap[c]; ok {
			windowMap[c]++
			if windowMap[c] == needMap[c] {
				valid++
			}
		}
		// 此时 windowMap里出现了 needMap里的所有字符，且都是一次
		for valid == len(needMap) {

		}
		right++
	}
}
