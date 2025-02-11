package cycle_array

type CycleArray struct {
	arr   []int
	start int
	end   int
	size  int
	count int
}

func NewCycleArray() *CycleArray {
	return &CycleArray{
		arr:   make([]int, 0),
		start: 0,
		end:   0,
		size:  0,
		count: 0,
	}
}
