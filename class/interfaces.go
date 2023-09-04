package class

type Class interface {
	NarcissisticNumber(num int64) bool
	ParityOutlier(arr []int64) (int64, bool)
	NeedleInAHaystack(arr []string, keyword string) int
	TheBlueOceanReverse(arr1, arr2 []int64) []int64
}
