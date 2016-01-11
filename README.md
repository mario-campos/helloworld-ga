# helloworld-ga [![Build Status](https://travis-ci.org/mario-campos/helloworld-ga.svg)](https://travis-ci.org/mario-campos/helloworld-ga)

A simple "Hello, World!" program, using a genetic algorithm.

### What is a Genetic Algorithm?

Courtesy of [Wikipedia][1]:
>a genetic algorithm (GA) is a search heuristic that mimics the process of natural evolution.

In other words, you can simulate evolution in a program to solve difficult optimization-class problems.

### How It Works

The algorithm operates on two "objects"

* a `Population` -- an unordered list of `Chromosome`s
* a `Chromosome` -- an individual evolvable building block

...at any given state:

* Initialization
* Crossover
* Mutation
* Fitness

The Initialization phase is the starting point. During the Initialization phase, the Population is set to a pre-determined size of random Chromosomes. Next, the Population undergoes Crossover -- the Chromosomes are optimally mated with each other to produce the next-generation Population. The new-born Population is mutated to introduce variety and evolution. Finally, the Population undergoes fitness (or conversely, cost). The fitness/cost determines the optimality of each Chromosome in the Population. Ideally, each generation is  geneticaly superior to its predecessor. This process repeats until any candidate is optimal enough, or "solved."

### What I've Learned

* Crossover covers the 90%

The solution is only as good as the algorithm's Crossover function. If Crossover can't reasonably close to a solution, then Mutation won't help. In other words, most of the evolutionary progress occurs in the Crossover phase.

* Mutation covers the 10%

Altering the frequency of Mutation adjusts the algorithm's ability to converge to a solution. Mutate too little, and the algorithm will continue asymptotically, failing to converge to the most optimal solution. Mutate too much, and the algorithm will either not converge (due to its inability to maintain consistent progressive evolution) or require more time converging to the most optimal solution.

* Mutation is more or less important, depending on the size of the population

The importantance of Mutation is to introduce new variety, ultimately to continue evolution when Crossover can't.This is more likely to happen in a smaller population as there is less chance for variety than in a larger population. Therefore, Mutation plays a larger role in smaller populations, since Crossover will likely result in inbreeding.

### Documentation

[Documentation][3]

### References

Burak Kanber's [Machine Learning: Genetic Algorithms Tutorial][4]

### License

*helloworld-ga* is licensed with [Mozilla Public License 2.0][2].

### INSTALL

```
go get github.com/mario-campos/helloworld-ga
```

### USE

```
$GOPATH/bin/helloworld-ga
[ 1 ]   Ilvqw/%hix\g'   (93)
[ 2 ]   Pfmcv20Ooqlo!   (68)
[ 3 ]   Egktm, UpslP!   (40)
[ 4 ]   Egktm, WpslP!   (38)
[ 5 ]   Hhlhi-!Xogif    (33)
[ 6 ]   Felmm.$Qlqlc#   (24)
[ 7 ]   Felmq+$Qlqlc#   (23)
[ 8 ]   Gdomp+$Ynsk_!   (22)
[ 9 ]   Gfhkq) Wkrnd    (19)
[ 10 ]  Egkim, Wpslh    (18)
[ 11 ]  Eglqm, Wpslf!   (16)
[ 12 ]  Demln*"Woqld!   (11)
[ 13 ]  Hglln- Uosld#   (9)
[ 14 ]  Hgmln+ Uoqld#   (10)
[ 15 ]  Ehlmp, Woslc!   (10)
[ 16 ]  Hdllp+ Xoskf!   (8)
[ 17 ]  Igmlp. Znskc!   (14)
[ 18 ]  Hdmlp- Xpskc!   (9)
[ 19 ]  Ielmn+ Worjb!   (8)
[ 20 ]  Jejmn, World#   (8)
[ 21 ]  Hdlln- Vorjd"   (7)
[ 22 ]  Hdlmo. Uormd!   (7)
[ 23 ]  Hdmlp+ Xoslc!   (7)
[ 24 ]  Hdllp, Xoskd!   (5)
[ 25 ]  Ifllo, Voqnd!   (6)
[ 26 ]  Gello- Xoqmd!   (5)
[ 27 ]  Eelmp- Wosld!   (7)
[ 28 ]  Gfllo, Xoqnd!   (6)
[ 29 ]  Hello,!Xormd!   (3)
[ 30 ]  Heilp,!World!   (5)
[ 31 ]  Helmo-!Woskd    (6)
[ 32 ]  Hello- Workd"   (3)
[ 33 ]  Eello, Wosld!   (4)
[ 34 ]  Hcllo, Workd"   (4)
[ 35 ]  Hellp, World!   (1)
[ 36 ]  Helmo.!Wormd!   (5)
[ 37 ]  Hekln, Workd!   (3)
[ 38 ]  Hdlmo, Wormd!   (3)
[ 39 ]  Hellp, Vorld!   (2)
[ 40 ]  Hellp, World!   (1)
[ 41 ]  Hellp, World!   (1)
[ 42 ]  Hfllo, World!   (1)
[ 43 ]  Hello, Xorld!   (1)
[ 44 ]  Helmo, World!   (1)
[ 45 ]  Hello, Vprld!   (2)
[ 46 ]  Hello, World!   (0)
```

[1]:http://en.wikipedia.org/w/index.php?title=Genetic_algorithm&oldid=516133477
[2]:http://www.mozilla.org/MPL/2.0/
[3]:http://go.pkgdoc.org/github.com/iamrekcah/helloworld-ga/src/hwga
[4]:http://burakkanber.com/blog/machine-learning-genetic-algorithms-part-1-javascript/
