package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

const (
	VALUE_TO_CHANGE = 63   // VALUE_TO_CHANGE is the target change value
	ITER_NUM        = 200  // ITER_NUM is the number of iteration to do
	P_MUT           = 0.10 // P_MUT is the probability of mutation
	P_CROSS         = 0.8  // P_CROSS is the probability of crossing
	MAX_UTILITY     = 50   // MAX_UTILITY is the maximum value that the utility can have

)

func main() {
	population := InitializePopulation(4)
	population.Print()
	for i := 1; i <= ITER_NUM; i++ {
		population = population.GenerateNewPopulation()
		population.Print()
	}
	fmt.Println("End")
}
