/* Any copyright is dedicated to the Public Domain.
 * http://creativecommons.org/publicdomain/zero/1.0/ */

package hwga

import "testing"

func Test_Populate(t *testing.T) {
	p := make(Population, 10)
	p.Populate()

	// expected: no nil values
	for i, v := range p {
		if v == nil {
			t.Error("position", i, "is nil")
			break
		}
	}
}

func Test_Crossover(t *testing.T) {
	test := make(Population, 3)
	test.Populate()

	control := make(Population, 3)
	var a *Chromosome = test[0]
	var b *Chromosome = test[1]
	var c *Chromosome = test[2]

	if c.Score < b.Score {
		control[0] = a.Mate(c)
	} else {
		control[0] = a.Mate(b)
	}

	if a.Score < c.Score {
		control[1] = b.Mate(a)
	} else {
		control[1] = b.Mate(c)
	}

	if b.Score < a.Score {
		control[2] = c.Mate(b)
	} else {
		control[2] = c.Mate(a)
	}

	test = test.Crossover()

	for i := range test {
		control[i].CalculateCost()
		test[i].CalculateCost()

		if test[i].Score != control[i].Score {
			t.Error("expected", control[i].Code, "; got", test[i].Code)
		}
	}
}

func Test_Mutate(t *testing.T) {
	p := make(Population, 1)
	p.Populate()

	var sum_before, sum_after int
	for _, val := range p[0].Code {
		sum_before += int(val)
	}

	p.Mutate()

	for _, val := range p[0].Code {
		sum_after += int(val)
	}

	if sum_before == sum_after {
		t.Error("before", sum_before, "; after", sum_after)
	}
}

func Test_Cost(t *testing.T) {
	p := make(Population, 1)
	p.Populate()

	for i := range p[0].Code {
		p[0].Code[i] = 0
	}
	p[0].Score = 0

	// expected return value: 1129
	if p.Cost(); p[0].Score != 1129 {
		t.Error("expected 1129; got", p[0].Score)
	}
}

func Test_Best(t *testing.T) {
	p := make(Population, 5)
	p.Populate()

	var best *Chromosome = p.Best()

	for _, c := range p {
		if c.Score < best.Score {
			t.Error(c.Score, "<", best.Score)
		}
	}
}

func Test_IsSolved(t *testing.T) {
	p := make(Population, 3)
	p.Populate()

	// test case: not solved (return value = false)
	for i := range p {
		p[i].Code = [13]byte{0}
		p[i].Score = 1129
	}

	if p.IsSolved() == true {
		t.Error("returned 'true' on [13]byte{0}")
	}

	// test case: solved (return value = true)
	p[2].Code = [13]byte{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	p[2].Score = 0

	if p.IsSolved() == false {
		t.Error("returned 'false' on 'Hello, World!'")
	}
}
