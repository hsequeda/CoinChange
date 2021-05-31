package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

const valueToChange = 60
const iterNum = 10

func main() {
	var population []Individual = initializePopulation()
	for i := 0; i < iterNum; i++ {
		// TODO do something
	}

	fmt.Printf("Final population:\n %#v \n", population)
}

func initializePopulation() []Individual {
	var population = make([]Individual, 0)
	// TODO do something
	return population
}

// func appendBaseSample(n Individual) {
// 	exist := false
// 	for _, val := range baseSample {
// 		if val.list == n.list {
// 			exist = true
// 		}
// 	}

// 	if !exist {
// 		baseSample = append(baseSample, n)
// 	}
// }

// func getCrossing(a, b Individual) Individual {
// 	ns := NewSample([4]int8{crossing(a.list[0], b.list[0], a.v < valueToChange),
// 		crossing(a.list[1], b.list[1], a.v < valueToChange),
// 		crossing(a.list[2], b.list[2], a.v < valueToChange),
// 		crossing(a.list[3], b.list[3], a.v < valueToChange),
// 	})

// 	if rand.Float32() > 0.3 {
// 		if ns.v != valueToChange {
// 			d := valueToChange - ns.v
// 			// fmt.Println(ns.v)
// 			// fmt.Println(d)
// 			switch {
// 			case d >= 25:
// 				return NewSample([4]int8{ns.list[0] + 1, ns.list[1], ns.list[2], ns.list[3]})
// 			case d >= 10:
// 				return NewSample([4]int8{ns.list[0], ns.list[1] + 1, ns.list[2], ns.list[3]})
// 			case d >= 5:
// 				return NewSample([4]int8{ns.list[0], ns.list[1], ns.list[2] + 1, ns.list[3]})
// 			default:
// 				return NewSample([4]int8{ns.list[0], ns.list[1], ns.list[2], ns.list[3] + 1})
// 			}
// 		}
// 	}
// 	return ns
// }

// func crossing(chromosomeFromA, chromosomeFromB int8, toMax bool) int8 {
// 	if toMax {
// 		return int8(math.Max(float64(chromosomeFromA), float64(chromosomeFromB)))
// 	}
// 	return int8(math.Min(float64(chromosomeFromA), float64(chromosomeFromB)))
// }
