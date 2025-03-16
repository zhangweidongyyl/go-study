package compress

// 适用于数组中大部分数据都是默认数据的情况，实现简单，存储效率较高

type SpareMatrix struct {
	DefaultValue         int         // 数组中大部分的默认值
	OtherDefaultValueMap map[int]int // 非默认值的值以及索引
	Length               int
}
type OriginSpareMatrixNums []int

func (datas OriginSpareMatrixNums) Compress(defaultValue int) SpareMatrix {
	spareMatrix := SpareMatrix{
		DefaultValue:         defaultValue,
		OtherDefaultValueMap: nil,
		Length:               len(datas),
	}
	otherDefaultValueMap := make(map[int]int, 0)
	for i, data := range datas {
		if data != defaultValue {
			otherDefaultValueMap[i] = data
		}
	}
	spareMatrix.OtherDefaultValueMap = otherDefaultValueMap
	return spareMatrix
}
func (spareMatrix SpareMatrix) Decompress() []int {
	datas := make([]int, spareMatrix.Length)
	for i := 0; i < spareMatrix.Length; i++ {
		if data, ok := spareMatrix.OtherDefaultValueMap[i]; ok {
			datas[i] = data
		} else {
			datas[i] = spareMatrix.DefaultValue
		}

	}
	return datas
}
