package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type PopulationA struct {
	generation  int
	chromosomes []Chromosome
}

// InitializePopulation generate a base population with random chromosomes
func InitializePopulation(ChromosomeNumber int) Population {
	population := &PopulationA{
		generation:  0,
		chromosomes: make([]Chromosome, 0),
	}
	for len(population.chromosomes) < ChromosomeNumber {
		individual := GenerateNewRandomChromosome()
		if individual.GetTotalValue() == VALUE_TO_CHANGE && !population.contains(individual) {
			population.AddChromosome(individual)
		}
	}

	return population
}

// InitializePopulation a population without chromosomes
func NewEmptyPopulation(generation int) Population {
	return &PopulationA{
		generation:  generation,
		chromosomes: make([]Chromosome, 0),
	}
}

// GetChromosomesLen returns the length of chromosomes on the population
func (p *PopulationA) GetChromosomesLen() int {
	return len(p.chromosomes)
}

func (p *PopulationA) GetBestCandidate() Chromosome {
	return p.chromosomes[0]
}

func (p *PopulationA) contains(ind Chromosome) bool {
	for e := range p.chromosomes {
		if p.chromosomes[e].GetFenotipe() == ind.GetFenotipe() {
			return true
		}
	}
	return false
}

func (p *PopulationA) getCrossCandidate() (Chromosome, int) {
	// Generate candidate probability
	selectionProbability := rand.Float32()
	for e := range p.chromosomes {
		if p.GetAccumulatedProbabilityOf(e) >= selectionProbability {
			return p.chromosomes[e], e
		}
	}
	// Return last
	return p.chromosomes[len(p.chromosomes)-1], len(p.chromosomes) - 1
}

func (p *PopulationA) GenerateNewPopulation() Population {
	newPopulation := NewEmptyPopulation(p.generation + 1)
	for newPopulation.GetChromosomesLen() < p.GetChromosomesLen() {
		// Select first candidate
		candidate1, selectedIndex := p.getCrossCandidate()

		var candidate2 Chromosome
		var selectedIndex2 int = selectedIndex
		for selectedIndex2 == selectedIndex {
			// Select second candidate
			candidate2, selectedIndex2 = p.getCrossCandidate()
		}
		if canBeCrossed() {
			child1, child2 := crossCandidates(candidate1, candidate2)
			if isValidCandidate(child1) && newPopulation.GetChromosomesLen() < p.GetChromosomesLen() {
				newPopulation.AddChromosome(child1)
			}
			if isValidCandidate(child2) && newPopulation.GetChromosomesLen() < p.GetChromosomesLen() {
				newPopulation.AddChromosome(child2)
			}
		}
	}
	return newPopulation
}

func (p *PopulationA) Print() {
	fmt.Printf("Population %d: totalUtility %d , len %d \n", p.GetGeneration(), p.GetTotalUtility(), p.GetChromosomesLen())
	bestCandidate := p.GetBestCandidate()
	fmt.Printf("Best candidate: fenotipe: %d, utility: %d ,probability: %f, accumulatedProv: %f \n",
		bestCandidate.GetFenotipe(),
		bestCandidate.GetUtility(), p.GetProbabilityOf(0), p.GetAccumulatedProbabilityOf(0))
	fmt.Printf("1 candidate: fenotipe: %d, utility: %d ,probability: %f, accumulatedProv: %f \n",
		p.chromosomes[1].GetFenotipe(),
		p.chromosomes[1].GetUtility(), p.GetProbabilityOf(1), p.GetAccumulatedProbabilityOf(1))
	fmt.Printf("2 candidate: fenotipe: %d, utility: %d ,probability: %f, accumulatedProv: %f \n",
		p.chromosomes[2].GetFenotipe(),
		p.chromosomes[2].GetUtility(), p.GetProbabilityOf(2), p.GetAccumulatedProbabilityOf(2))
	fmt.Printf("3 candidate: fenotipe: %d, utility: %d ,probability: %f, accumulatedProv: %f \n",
		p.chromosomes[3].GetFenotipe(),
		p.chromosomes[3].GetUtility(), p.GetProbabilityOf(3), p.GetAccumulatedProbabilityOf(3))
}

func (p *PopulationA) AddChromosome(ind Chromosome) {
	p.chromosomes = append(p.chromosomes, ind)
	sort.Slice(p.chromosomes, func(i, j int) bool {
		return p.chromosomes[i].GetUtility() > p.chromosomes[j].GetUtility()
	})
}

func (p *PopulationA) GetTotalUtility() int {
	totalUtility := 0
	for e := range p.chromosomes {
		totalUtility += p.chromosomes[e].GetUtility()
	}
	return totalUtility
}

func (p *PopulationA) GetUtilityOf(index int) int {
	if index < 0 || index >= len(p.chromosomes) {
		panic("error")
	}

	return p.chromosomes[index].GetUtility()
}

func (p *PopulationA) GetTotalValueOf(index int) int {
	if index < 0 || index >= len(p.chromosomes) {
		panic("error")
	}

	return p.chromosomes[index].GetTotalValue()
}

func (p *PopulationA) GetGeneration() int {
	return p.generation
}

func (p *PopulationA) GetProbabilityOf(index int) float32 {
	return float32(p.GetUtilityOf(index)) / float32(p.GetTotalUtility())
}

func (p *PopulationA) GetAccumulatedProbabilityOf(index int) (acumulatedProbability float32) {
	acumulatedProbability = 0
	for i := 0; i <= index; i++ {
		acumulatedProbability += p.GetProbabilityOf(i)
	}
	return acumulatedProbability
}

func crossCandidates(parent1, parent2 Chromosome) (Chromosome, Chromosome) {
	list1 := [20]bool{}
	list2 := [20]bool{}
	crossPoint := int(float32(parent1.GetLen()) * rand.Float32())
	for i := 0; i < 20; i++ {
		if i < crossPoint {
			list1[i] = parent1.GetGen(i)
			list2[i] = parent2.GetGen(i)
		} else {
			list1[i] = parent2.GetGen(i)
			list2[i] = parent1.GetGen(i)
		}
		// Can child1 chain be muted
		if canBeMutated() {
			list1[i] = !list1[i]
		}
		// Can child2 chain be muted
		if canBeMutated() {
			list1[i] = !list1[i]
		}
	}
	return &ChromosomeA{
			list: list1,
		}, &ChromosomeA{
			list: list2,
		}
}

func canBeCrossed() bool {
	return rand.Float32() < P_CROSS
}

func canBeMutated() bool {
	return rand.Float32() < P_MUT
}

func isValidCandidate(ind Chromosome) bool {
	return ind.GetTotalValue() == VALUE_TO_CHANGE
}
