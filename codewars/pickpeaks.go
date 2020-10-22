package codewars

type PosPeaks struct {
	Pos 	[]int
	Peaks 	[]int
}

func PickPeaks(array []int) PosPeaks {
	pos, peaks := []int{}, []int{}
	tmp := 0
	for i := 1; i < len(array); i++ {
		if array[i-1] < array[i] {
			tmp = i
		} else if array[i-1] > array[i] && tmp > 0 {
			pos = append(pos, tmp)
			peaks = append(peaks, array[tmp])
			tmp = 0
		}
	}
	return PosPeaks{pos, peaks}
}