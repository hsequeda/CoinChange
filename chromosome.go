package main

type CoinType int

type ChromosomeA struct {
	list [20]bool
}

func GenerateNewRandomChromosome() Chromosome {
	list := [20]bool{}
	boolgen := NewBoolGen()
	for i := 0; i < 20; i++ {
		list[i] = boolgen.Bool()
	}
	return &ChromosomeA{
		list: list,
	}
}

// GetLen return the length of the chromosome chain
func (i *ChromosomeA) GetLen() int {
	return len(i.list)
}

// GetFenotipe returns the Real (No-Binary) representation of the Chromosome
func (ind *ChromosomeA) GetFenotipe() int {
	fenotipe := 0
	for index, isActive := range ind.list {
		if !isActive {
			continue
		}
		switch {
		case isTwentyFiveGuard(index):
			fenotipe += 1000
			break
		case isTenGuard(index):
			fenotipe += 100
			break
		case isFiveGuard(index):
			fenotipe += 10
			break
		case isOneGuard(index):
			fenotipe += 1
			break
		}
	}
	return fenotipe
}

// GetUtility retuns the total utility (Z) of the individual
// The utility of the individual is given by the amount of
// coins of each type that he owns.
func (ind *ChromosomeA) GetUtility() int {
	indUtility := 0
	for index, isActive := range ind.list {
		if !isActive {
			continue
		}
		switch {
		case isTwentyFiveGuard(index):
			indUtility += 1
			break
		case isTenGuard(index):
			indUtility += 2
			break
		case isFiveGuard(index):
			indUtility += 3
			break
		case isOneGuard(index):
			indUtility += 4
			break
		}
	}
	return MAX_UTILITY - indUtility
}

// GetTotalValue returns the quantity of cash of the Chromosome
func (ind *ChromosomeA) GetTotalValue() int {
	indTotalVal := 0
	for index, isActive := range ind.list {
		if !isActive {
			continue
		}
		switch {
		case isTwentyFiveGuard(index):
			indTotalVal += 25
			break
		case isTenGuard(index):
			indTotalVal += 10
			break
		case isFiveGuard(index):
			indTotalVal += 5
			break
		case isOneGuard(index):
			indTotalVal += 1
			break
		}
	}
	return indTotalVal
}

// GetChromosomeLen return the length of the chromosome chain
func (i *ChromosomeA) GetGen(index int) bool {
	return i.list[index]
}

func isTwentyFiveGuard(pos int) bool {
	return pos < 5
}

func isTenGuard(pos int) bool {
	return pos >= 5 && pos < 10
}

func isFiveGuard(pos int) bool {
	return pos >= 10 && pos < 14
}

func isOneGuard(pos int) bool {
	return pos >= 14
}
