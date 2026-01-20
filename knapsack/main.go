package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"strconv"
	"strings"
)

// knapsack problem implementation using genetic algo
// we have four items with weight and values we have to put them into knapsack with some weight capacity such that maximum value can be put inside knapsack within it's weitght limit.

/*
*
1. define initial population
2. generate generations
3. proceed with those individuals by thier fitness values
*/
type Pair struct {
	weight int
	value  int
}

func main() {
	var s string = solveKnapsack()

	fmt.Println("here's the final string" + s)
}

func solveKnapsack() string {
	// var totalItems int = 4
	var maxKnapsackCapacity = 15

	var itemsData = []Pair{{7, 5}, {2, 4}, {1, 7}, {9, 2}}

	var initialPopulationCount int = 6
	// var initialPopulation []string
	initialPopulation := generateInitialPopulation(initialPopulationCount)
	var avgFitness []string

	i := 0
	for i < 500 {
		avgFitness = append(avgFitness, initialPopulation...)
		initialPopulation = nextGeneration(initialPopulation, itemsData, maxKnapsackCapacity)
		i++
	}

	// populationAfterFitnessCalculation := calculateFitness(itemsData, initialPopulation, maxKnapsackCapacity)
	// sort.Slice(populationAfterFitnessCalculation, func(i, j int) bool {
	// 	return individualFitness(populationAfterFitnessCalculation[i], itemsData, i, maxKnapsackCapacity) > individualFitness(populationAfterFitnessCalculation[j], itemsData, j, maxKnapsackCapacity)
	// })
	// return populationAfterFitnessCalculation[0]

	// populationAfterFitnessCalculation := calculateFitness(itemsData, initialPopulation, maxKnapsackCapacity)
	// parentsPopulation := selectionProcess(initialPopulation, itemsData, maxKnapsackCapacity)

	// newParentsPopulation := crossoverProcess(parentsPopulation)
	// mutationProcess(newParentsPopulation, mutationRate)
}

func nextGeneration(population []string, iData []Pair, maxCap int) []string {
	var mutationRate float32 = 0.02
	var crossoverRate float32 = 0.02
	var reprodutionRate float32 = 0.3

	var nextGen []string

	parents := selectionProcess(population, iData, maxCap)

	for len(nextGen) < len(population) {
		var children []string

		if rand.Float32() < reprodutionRate {
			children = parents
		} else {
			var newParentsPopulation []string
			if rand.Float32() < crossoverRate {
				newParentsPopulation = crossoverProcess(parents)
			}
			if rand.Float32() < mutationRate && len(newParentsPopulation) > 0 {
				mutationProcess(newParentsPopulation, mutationRate)
			}
			children = newParentsPopulation
		}
		nextGen = append(nextGen, children...)
	}
	return nextGen
}

func mutationProcess(npPop []string, mutationRate float32) {
	sz := len(npPop)
	count := 0

	for count < sz {
		res := npPop[count]
		a := 0
		for a < len(res) {
			if mutationRate > rand.Float32() {
				if npPop[count][a] == '0' {
					res := []rune(npPop[count])
					res[a] = '1'
					npPop[count] = string(res)
				}
			}
			a++
		}
		count++
	}

}

func crossoverProcess(fitPop []string) []string {
	n := len(fitPop[0])
	var child1 string = fitPop[0][:n/2] + fitPop[1][n/2:]
	var child2 string = fitPop[1][:n/2] + fitPop[0][n/2:]

	return []string{child1, child2}
}

func selectionProcess(iPop []string, iData []Pair, maxCap int) []string {
	// take first 4 individuals and arrange a tournament
	var parents []string

	rand.Shuffle(len(iPop), func(i, j int) {
		iPop[i], iPop[j] = iPop[j], iPop[i]
	})

	if individualFitness(iPop[1], iData, 1, maxCap) > individualFitness(iPop[0], iData, 0, maxCap) {
		parents = append(parents, iPop[1])
	} else {
		parents = append(parents, iPop[0])
	}

	if individualFitness(iPop[2], iData, 2, maxCap) > individualFitness(iPop[3], iData, 3, maxCap) {
		parents = append(parents, iPop[2])
	} else {
		parents = append(parents, iPop[3])
	}
	return parents
}

func individualFitness(res string, iData []Pair, idx int, maxCap int) int {
	var totalWeight int = 0
	var totalValue int = 0
	count := 0
	for count < len(res) {
		if res[count] != '0' {
			totalWeight += iData[idx].weight
			totalValue += iData[idx].value
		}
	}

	if totalWeight > maxCap {
		return 0
	}
	return totalValue
}

// func calculateFitness(iData []Pair, iPop []string, maxCap int) []string {
// 	var res []string = iPop
// 	count := 0

// 	for count != len(res) {
// 		temp := res[count]
// 		sz := len(temp)
// 		a := 0
// 		var fitness int = 0

// 		for a < sz {
// 			if temp[a] != '0' {
// 				fitness += iData[count].weight
// 			}
// 			a++
// 		}
// 		if fitness > maxCap {
// 			res = append(res[:count], res[count+1:]...)
// 		}
// 		count--
// 	}
// 	return res
// }

func generateInitialPopulation(n int) []string {
	var arr []string
	for len(arr) != n {
		var res strings.Builder
		for range 4 {
			res.WriteString(strconv.Itoa(rand.IntN(2)))
		}
		arr = append(arr, res.String())
	}
	return arr
}
