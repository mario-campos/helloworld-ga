package main;

import (
	"fmt"
	"math/rand"
	"time"
)

type Population []*Chromosome;

type Chromosome struct {
	code  [13]byte // genetic code
	score int      // fitness score
}

func (c *Chromosome) initialize() {
	for i := range c.code {
		c.code[i] = byte(rand.Intn(90) + 32)
	}

	c.calculate()
}

func (c *Chromosome) calculate() {
	c.score = 0

	for i := range c.code {
		c.score += c.eval(i)
	}
}

func (a *Chromosome) mate(b *Chromosome) *Chromosome {
	c := new(Chromosome)

	for i := 0; i < 13; i++ {
		if a.eval(i) < b.eval(i) {
			c.code[i] = a.code[i]
		} else {
			c.code[i] = b.code[i]
		}
	}

	return c
}

func (c *Chromosome) eval(i int) int {
	var x int

	switch i {
	case 0:  x = 72  // H
        case 1:  x = 101 // e
        case 2:  x = 108 // l
        case 3:  x = 108 // l
        case 4:  x = 111 // o
        case 5:  x = 44  // ,
        case 6:  x = 32  // <space>
        case 7:  x = 87  // W
        case 8:  x = 111 // o
        case 9:  x = 114 // r
        case 10: x = 108 // l
        case 11: x = 100 // d
        case 12: x = 33  // !
	}

	eval_val := x - int(c.code[i])

	if eval_val < 0 {
		eval_val = -eval_val
	}

	return eval_val
}

func populate(size int) Population {
	p := make(Population, size)

	for i := range p {
		p[i] = new(Chromosome)
		p[i].initialize()
	}
	
	return p
}

func crossover(p Population) Population {
	var h,j int
	var np Population = populate(len(p))

	for i := range p {
		if i == 0 {
			h = len(p)-1
			j = 1
		} else if i == len(p)-1 {
			h = i-1
			j = 0
		} else {
			h = i-1
			j = i+1
		}
				
		if p[h].score > p[j].score {
			np[i] = p[i].mate(p[j])
		} else {
			np[i] = p[i].mate(p[h])
		}
	}

	return np
}

func mutate(p Population) {
	for i := range p {
		pos := rand.Intn(13)
		val := byte(rand.Intn(90) + 32)

		p[i].code[pos] = val
	}
}

func fitness(p Population) {
	for i := range p {
		p[i].calculate()
	}
}

func salute(p Population, gen int) {
	var c *Chromosome = p[0]

	for i := range p {
		if p[i].score < c.score {
			c = p[i]
		}
	}

	fmt.Printf("[ %d ]\t%s\t(%d)\n", gen, c.code, c.score)
}

func solved(p Population) bool {
	var retval bool = false

	for i := range p {
		if p[i].score == 0 {
			retval = true
			break
		}
	}

	return retval	
}

func main() {
	rand.Seed(time.Now().Unix())
	pop := populate(1000)

	for i := 1; !solved(pop); i++ {
		pop = crossover(pop)
		mutate(pop)
		fitness(pop)
		salute(pop, i)
	}
}