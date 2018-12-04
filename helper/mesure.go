package helper

import (
	"fmt"
	"knapsack/algorithms"
	"time"
)

func relativeError(optimal int, apx int) float64 {
	return float64((optimal - apx))/float64(optimal)
}

func MesureInstance() {

	var instanceTimeBF float64 = 0
	var instanceTimeWeight float64 = 0
	var instanceTimeValue float64 = 0
	var instanceTimeCoef float64 = 0

	var (
		sumErrorWeight float64 = 0
		sumErrorValue float64 = 0
		sumErrorCoef float64 = 0
	)

	for {
		_, capacity, items, error := ReadInstance()
		if error != nil {
			fmt.Printf("%v\n", instanceTimeBF)
			fmt.Printf("%v\n", (instanceTimeWeight + instanceTimeValue + instanceTimeCoef)/3)
			break
		}

		start := time.Now()
		_, valueBF, _ := algorithms.KnapsackBruteForce(capacity, items, []int{}, 0, 0, 0)
		elapsed := time.Since(start).Seconds()
		instanceTimeBF += elapsed

		start = time.Now()
		valueWeight, _ := algorithms.KnapsackHeuristic(capacity, items, algorithms.WeightSorter(items))
		elapsed = time.Since(start).Seconds()
		instanceTimeWeight += elapsed

		start = time.Now()
		valueValue, _ := algorithms.KnapsackHeuristic(capacity, items, algorithms.ValueSorter(items))
		elapsed = time.Since(start).Seconds()
		instanceTimeValue += elapsed

		start = time.Now()
		valueCoef, _ := algorithms.KnapsackHeuristic(capacity, items, algorithms.CoefSorter(items))
		elapsed = time.Since(start).Seconds()
		instanceTimeCoef += elapsed

		sumErrorWeight += relativeError(valueBF, valueWeight)
		sumErrorValue += relativeError(valueBF, valueValue)
		sumErrorCoef += relativeError(valueBF, valueCoef)

		//fmt.Printf("BF %v , W %v, P %v, C %v\n", valueBF, valueWeight, valueValue, valueCoef)
		//fmt.Printf("E Weight: %v\n", errorWeight)
		//fmt.Printf("E Value: %v\n", errorValue)
		//fmt.Printf("E Coef: %v\n", errorCoef)
	}

	fmt.Printf("%v\n", sumErrorWeight/50)
	fmt.Printf("%v\n", sumErrorValue/50)
	fmt.Printf("%v\n", sumErrorCoef/50)

}

func MesureFPTASAcc() {

	accuracy := []float32{0.001, 0.01, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	accuracyError := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for {
		_, capacity, items, error := ReadInstance()
		if error != nil {
			fmt.Printf("\n")
			for i := 0; i < len(accuracy); i++ {
				fmt.Printf("%v\n", accuracyError[i]/50)
			}
			break
		}

		precalSumValue := make([]int, len(items) + 1)
		for i := 0; i < len(items); i++ {
			precalSumValue[i] = 0
			for j := i; j < len(items); j++ {
				precalSumValue[i] += items[j].Value
			}
		}
		precalSumValue[len(items)] = 0

		for i := 0; i < len(accuracy); i++ {
			valueBB, _, _ := algorithms.KnapsackBranchAndBound(capacity, items, 0, 0, 0, make([]int, 0), precalSumValue)
			valueFPTAS, _ := algorithms.KnapsackFPTAS(capacity, items, accuracy[i])

			accuracyError[i] += relativeError(valueBB, valueFPTAS)
		}

	}
}
