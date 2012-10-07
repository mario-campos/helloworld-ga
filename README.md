# helloworld-ga

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

### License

*helloworld-ga* is licensed with [Mozilla Public License 2.0][2].

[1]:http://en.wikipedia.org/w/index.php?title=Genetic_algorithm&oldid=516133477
[2]:http://www.mozilla.org/MPL/2.0/