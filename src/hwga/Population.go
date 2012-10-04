/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package hwga

import "math/rand"

type Population []*Chromosome

func (p Population) Populate() {
	for i := range p {
		p[i] = new(Chromosome)
		p[i].Initialize()
	}
}

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

		if p[h].score > p[j].score {
			np[i] = p[i].Mate(p[j])
		} else {
			np[i] = p[i].Mate(p[h])
		}
	}

	return np
}

func (p Population) Mutate() {
	for i := range p {
		pos := rand.Intn(13)
		val := byte(rand.Intn(90) + 32)
		p[i].code[pos] = val
	}
}

func (p Population) Cost() {
	for i := range p {
		p[i].CalculateCost()
	}
}

func (p Population) Best() *Chromosome {
	var c *Chromosome = p[0]

	for i := range p {
		if p[i].score < c.score {
			c = p[i]
		}
	}

	return c
}

func (p Population) IsSolved() bool {
	var retval bool = false

	for i := range p {
		if p[i].score == 0 {
			retval = true
			break
		}
	}

	return retval
}
