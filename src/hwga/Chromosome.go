/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

// Package hwga provides a Genetic Algorithm to resolve random strings 
// to 'Hello, World!'.
package hwga

import "math/rand"

// A Chromosome, the building block of a GA, stores a genetic 
// code (string) and its fitness/cost.
type Chromosome struct {
	code  [13]byte
	score int
}

// Initialize assigns a random string to the Chromosome.
func (c *Chromosome) Initialize() {
	for i := range c.code {
		c.code[i] = byte(rand.Intn(90) + 32)
	}

	c.CalculateCost()
}

// CalculateCost performs the cost/fitness of the Chromosome's string
// and stores the result in Chromosome.score.
func (c *Chromosome) CalculateCost() {
	c.score = 0

	for i := range c.code {
		c.score += c.Evaluate(i)
	}
}

// Mate, given a pointer to a Chromosome, optimally merges the two 
// Chromosomes and returns a pointer to the offspring Chromosome.
func (a *Chromosome) Mate(b *Chromosome) *Chromosome {
	c := new(Chromosome)

	for i := 0; i < 13; i++ {
		if a.Evaluate(i) < b.Evaluate(i) {
			c.code[i] = a.code[i]
		} else {
			c.code[i] = b.code[i]
		}
	}

	return c
}

// Evaluate, given an index, returns an absolute numerical difference
// of the indexed ASCII character and the ideal ASCII character.
func (c *Chromosome) Evaluate(i int) int {
	var x int

	switch i {
	case 0:
		x = 72 // H
	case 1:
		x = 101 // e
	case 2:
		x = 108 // l
	case 3:
		x = 108 // l
	case 4:
		x = 111 // o
	case 5:
		x = 44 // ,
	case 6:
		x = 32 // <space>
	case 7:
		x = 87 // W
	case 8:
		x = 111 // o
	case 9:
		x = 114 // r
	case 10:
		x = 108 // l
	case 11:
		x = 100 // d
	case 12:
		x = 33 // !
	}

	eval_val := x - int(c.code[i])

	if eval_val < 0 {
		eval_val = -eval_val
	}

	return eval_val
}
