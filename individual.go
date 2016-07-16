package gago

import (
	"math"
	"math/rand"
	"sort"
)

// A Genome contains genes
type Genome []interface{}

// An Individual represents a potential solution to a problem. The individual's
// is defined by it's genome, which is a slice containing genes. Every gene is a
// floating point numbers. The fitness is the individual's phenotype and is
// represented by a floating point number.
type Individual struct {
	Genome  Genome
	Fitness float64
	Age     int
	Name    string
}

// Evaluate the fitness of an individual.
func (indi *Individual) evaluate(ff FitnessFunction) {
	// Don't evaluate individuals that have already been evaluated
	if indi.Age == 0 {
		indi.Fitness = ff.apply(indi.Genome)
	}
	indi.Age++
}

// Generate a new individual.
func makeIndividual(nbGenes int, rng *rand.Rand) Individual {
	return Individual{
		Genome:  make([]interface{}, nbGenes),
		Fitness: math.Inf(1),
		Age:     0,
		Name:    randomName(6, rng),
	}
}

// Individuals type is necessary for sorting and selection purposes.
type Individuals []Individual

// Evaluate each individual
func (indis Individuals) evaluate(ff FitnessFunction) {
	for i := range indis {
		indis[i].evaluate(ff)
	}
}

// Generate a slice of new individuals.
func makeIndividuals(nbIndis, nbGenes int, rng *rand.Rand) Individuals {
	var indis = make(Individuals, nbIndis)
	for i := range indis {
		indis[i] = makeIndividual(nbGenes, rng)
	}
	return indis
}

// Sort the individuals of a population in ascending order based on their
// fitness. The convention is that we always want to minimize a function. A
// function f(x) can be function maximized by minimizing -f(x) or 1/f(x).
func (indis Individuals) Len() int           { return len(indis) }
func (indis Individuals) Less(i, j int) bool { return indis[i].Fitness < indis[j].Fitness }
func (indis Individuals) Swap(i, j int)      { indis[i], indis[j] = indis[j], indis[i] }

// Convenience method for calling the Sort method of the sort package
func (indis Individuals) sort() { sort.Sort(indis) }

// Sample k unique individuals from a slice of n individuals.
func (indis Individuals) sample(k int, rng *rand.Rand) ([]int, Individuals) {
	var (
		indexes, _ = randomInts(k, 0, len(indis), rng)
		sample     = make(Individuals, k)
	)
	for i := 0; i < k; i++ {
		sample[i] = indis[indexes[i]]
	}
	return indexes, sample
}

// Extract the fitness of a slice of individuals into a float64 slice.
func (indis Individuals) getFitnesses() []float64 {
	var fitnesses = make([]float64, len(indis))
	for i, indi := range indis {
		fitnesses[i] = indi.Fitness
	}
	return fitnesses
}

// FitnessMean returns the average fitness of a slice of individuals.
func (indis Individuals) FitnessMean() float64 {
	return mean(indis.getFitnesses())
}

// FitnessVar returns the variance of the fitness of a slice of individuals.
func (indis Individuals) FitnessVar() float64 {
	return variance(indis.getFitnesses())
}
