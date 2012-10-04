package hwga

import "math/rand"

type Chromosome struct {
	code  [13]byte // genetic code
	score int      // fitness score
}

func (c *Chromosome) Initialize() {
	for i := range c.code {
		c.code[i] = byte(rand.Intn(90) + 32)
	}

	c.CalculateCost()
}

func (c *Chromosome) CalculateCost() {
	c.score = 0

	for i := range c.code {
		c.score += c.Evaluate(i)
	}
}

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
