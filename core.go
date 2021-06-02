package main

// Chromosome is a entity of the population
type Chromosome interface {
	// GetUtility retuns the total utility (Z) of the individual
	// The utility of the individual is given by the amount of
	// coins of each type that he owns.
	GetUtility() int
	// GetTotalValue returns the quantity of cash of the Chromosome
	GetTotalValue() int
	// GetFenotipe returns the Real (No-Binary) representation of the Chromosome
	GetFenotipe() int
	// GetGen return a gen of the chromosome by its index
	GetGen(index int) bool
	// GetLen return the length of the chromosome chain
	GetLen() int
}

// Population is a composition of all the chromosomes of one generation
type Population interface {
	// AddChromosome add a new chromosome to the population
	AddChromosome(ind Chromosome)
	// GetGeneration returns the generation of the population
	GetGeneration() int
	// GetProbabilityOf return the probability of one chromosome by its index
	GetProbabilityOf(index int) float32
	// GetProbabilityOf return the accumulated probability of one chromosome by its index
	GetAccumulatedProbabilityOf(index int) float32
	// GetProbabilityOf return the total utility of the population
	GetTotalUtility() int
	// GetProbabilityOf return the utility of one chromosome by its index
	GetUtilityOf(index int) int
	// GetProbabilityOf return the total value of one chromosome by its index
	GetTotalValueOf(index int) int
	// GetBestCandidate Return the best candidate (chromosome) of a popution
	GetBestCandidate() Chromosome
	// GetChromosomesLen return the length of chromosomes of the population
	GetChromosomesLen() int
	// GenerateNewPopulation generate the next generation of the population
	GenerateNewPopulation() Population
	// Print write on the stdout the abstract of a generation
	Print()
}
