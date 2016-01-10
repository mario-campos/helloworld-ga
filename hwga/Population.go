/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package hwga

import "math/rand"

// A Population is a slice (array) of Chromosome pointers.
type Population []*Chromosome

// Populate assigns a new Initialized Chromosome to each element.
func (p Population) Populate() {
	for i := range p {
		p[i] = new(Chromosome)
		p[i].Initialize()
	}
}

// Crossover Mates each Chromosome with its most-optimal neighbor and
// returns a new Population of the offspring Chromosomes.
func (p Population) Crossover() Population {
	var h, j int
	np := make(Population, len(p))
	np.Populate()

	for i := range p {
		if i == 0 {
			h = len(p) - 1
			j = 1
		} else if i == len(p)-1 {
			h = i - 1
			j = 0
		} else {
			h = i - 1
			j = i + 1
		}

		if p[h].Score > p[j].Score {
			np[i] = p[i].Mate(p[j])
		} else {
			np[i] = p[i].Mate(p[h])
		}

		np[i].CalculateCost()
	}

	return np
}

// Mutate assigns a random ASCII character to a random index of each 
// Chromosome's genetic string.
func (p Population) Mutate() {
	for i := range p {
		pos := rand.Intn(13)
		val := byte(rand.Intn(90) + 32)
		p[i].Code[pos] = val
	}
}

// Cost performs the cost/fitness function on each Chromosome.
func (p Population) Cost() {
	for i := range p {
		p[i].CalculateCost()
	}
}

// Best returns a pointer to the most-optimal Chromosome in the Population.
func (p Population) Best() *Chromosome {
	var c *Chromosome = p[0]

	for i := range p {
		if p[i].Score < c.Score {
			c = p[i]
		}
	}

	return c
}

// IsSolved determines whether or not the Population contains an ideal Chromosome.
func (p Population) IsSolved() bool {
	var retval bool = false

	for i := range p {
		if p[i].Score == 0 {
			retval = true
			break
		}
	}

	return retval
}
