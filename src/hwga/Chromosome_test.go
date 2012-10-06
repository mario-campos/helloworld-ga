/* Any copyright is dedicated to the Public Domain.
 * http://creativecommons.org/publicdomain/zero/1.0/ */

package hwga

import (
	"math/rand"
	"testing"
	"time"
)

func Test_Initialize(t *testing.T) {
	rand.Seed(time.Now().Unix())

	var (
		a   *Chromosome = new(Chromosome)
		sum int         = 0
	)

	a.Initialize()

	for i := range a.Code {
		sum += int(a.Code[i])
	}

	if sum == 0 {
		t.Error("Code:", a.Code, "; expected not 0; got 0")
	}
}

func Test_CalculateCost(t *testing.T) {
	c := new(Chromosome)

	// expected return value: 1129
	if c.CalculateCost(); c.Score != 1129 {
		t.Error("Code:", c.Code, "; expected 1129; got", c.Score)
	}

	// expected return value: 13
	c.Code = [13]byte{73, 102, 109, 109, 110, 45, 33, 88, 112, 113, 109, 99, 32}
	if c.CalculateCost(); c.Score != 13 {
		t.Error("Code:", c.Code, "; expected 13; got", c.Score)
	}

	// expected return value: 0
	c.Code = [13]byte{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	if c.CalculateCost(); c.Score != 0 {
		t.Error("Code:", c.Code, "; expected 0; got", c.Score)
	}
}

func Test_Mate(t *testing.T) {
	m := new(Chromosome)
	f := new(Chromosome)

	// expected return value: 0
	m.Code = [13]byte{72, 0, 108, 0, 111, 0, 32, 0, 111, 0, 108, 0, 33}
	f.Code = [13]byte{0, 101, 0, 108, 0, 44, 0, 87, 0, 114, 0, 100, 0}

	c := m.Mate(f)
	c.CalculateCost()

	if c.Score != 0 {
		t.Error("result:", c.Code, "; expected 0; got", c.Score)
	}
}

func Test_Evaluate(t *testing.T) {
	c := new(Chromosome)

	// expected return value: 1	
	c.Code[0] = 73
	if z := c.Evaluate(0); z != 1 {
		t.Error("passed", string(byte(73)), "; expected 1; got", z)
	}

	c.Code[1] = 100
	if z := c.Evaluate(1); z != 1 {
		t.Error("passed", string(byte(100)), "; expected 1; got", z)
	}

	// expected return value: 0
	c.Code[2] = 108
	if z := c.Evaluate(2); z != 0 {
		t.Error("passed", string(byte(100)), "; expected 0; got", z)
	}

	// expected return value: 147
	c.Code[3] = 255
	if z := c.Evaluate(3); z != 147 {
		t.Error("passed", string(byte(255)), "; expected 147; got", z)
	}
}
