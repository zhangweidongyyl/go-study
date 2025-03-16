package compress

func DiffEncode(datas []int) []int {
	if len(datas) == 0 {
		return nil
	}
	res := make([]int, 0)
	res[0] = datas[0]
	for i := 1; i < len(datas); i++ {

	}
	return res
}
