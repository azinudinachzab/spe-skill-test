package main

import (
	"fmt"

	spe "github.com/azinudinachzab/spe-skill-test/class"
)

func main() {
	speClass := spe.InitSpeSkillTest()
	fmt.Println("NarcissisticNumber Sample 1 (153): ", speClass.NarcissisticNumber(153))
	fmt.Println("NarcissisticNumber Sample 2 (111): ", speClass.NarcissisticNumber(111))

	res1, ok1 := speClass.ParityOutlier([]int64{
		2,4,0,100,4,11,2602,36,
	})
	res2, ok2 := speClass.ParityOutlier([]int64{
		11,13,15,19,9,13,-21,
	})
	fmt.Println("ParityOutlier Sample 1 (2,4,0,100,4,11,2602,36): ", res1, ok1)
	fmt.Println("ParityOutlier Sample 1 (11,13,15,19,9,13,-21): ", res2, ok2)

	fmt.Println("NeedleInAHaystack Sample 1 (red, blue, yellow, black, grey): ", speClass.NeedleInAHaystack([]string{
		"red","blue","yellow","black","grey",
	}, "black"))

	fmt.Println("TheBlueOceanReverse Sample 1 (1, 2, 3, 4, 6, 10), (1): ", speClass.TheBlueOceanReverse([]int64{
		1,2,3,4,6,10}, []int64{1}))

	fmt.Println("TheBlueOceanReverse Sample 2 (1, 5, 5, 5, 5, 3), (5): ", speClass.TheBlueOceanReverse([]int64{
		1,5,5,5,5,3}, []int64{5}))
}
