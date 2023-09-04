package class

import (
	"fmt"
	"math"
	"strconv"
)

type SpeSkillTest struct {}

func (s *SpeSkillTest) NarcissisticNumber(num int64) bool {
	strNum := fmt.Sprintf("%v", num)
	digit := len(strNum)

	var result float64 = 0
	for _, v := range strNum {
		tempNum, err := strconv.ParseFloat(string(v), 10)
		if err != nil {
			panic(err)
		}
		result = result + math.Pow(tempNum, float64(digit))
	}

	if int64(result) == num {
		return true
	}
	return false
}

func (s *SpeSkillTest) ParityOutlier(arr []int64) (int64,bool) {
	oddArr := make([]int64, 0)
	evenArr := make([]int64, 0)

	for _, val := range arr {
		if val % 2 == 0 {
			evenArr = append(evenArr, val)
			continue
		}
		oddArr = append(oddArr, val)
	}
	if len(oddArr) == 0|| len(evenArr) == 0 {
		return 0, false
	}

	if len(oddArr) == 1 {
		return oddArr[0], true
	}
	return evenArr[0], true
}

func (s *SpeSkillTest) NeedleInAHaystack(arr []string, keyword string) int {
	for idx, val := range arr {
		if val == keyword {
			return idx
		}
	}
	return -1
}

func (s *SpeSkillTest) TheBlueOceanReverse(arr1, arr2 []int64) []int64 {
	arrRes := make([]int64,0)
	mapNumProhibit := make(map[int64]bool, 0)
	for _, val := range arr2 {
		mapNumProhibit[val] = true
	}

	for _, val := range arr1 {
		exist, _ := mapNumProhibit[val]
		if exist {
			continue
		}
		arrRes = append(arrRes, val)
	}

	return arrRes
}

func InitSpeSkillTest() Class {
	return &SpeSkillTest{}
}
