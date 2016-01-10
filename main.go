/* Any copyright is dedicated to the Public Domain.
 * http://creativecommons.org/publicdomain/zero/1.0/ */

package main

import (
	"fmt"
	"time"
	"math/rand"

	hwga "github.com/mario-campos/helloworld-ga/hwga"
)

func main() {
	rand.Seed(time.Now().Unix())
	
	p := make(hwga.Population, 1000)
	p.Populate()

	for gen := 1; !p.IsSolved(); gen++ {
		p = p.Crossover()
		p.Mutate()
		p.Cost()
		best := p.Best()
		fmt.Printf("[ %d ]\t%s\t(%d)\n", gen, best.Code, best.Score)
	}
}
