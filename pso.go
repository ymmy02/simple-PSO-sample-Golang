package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	DIM        = 2
	POPULATION = 10
	LOOPMAX    = 10
	SKIP       = 1
	STT        = -5.0
	END        = 5.0
	W          = 0.5
)

type Particle struct {
	Pos       []float64
	Vel       []float64
	Fitness   float64
	Pbest     []float64
	Lbest     []float64
	Neighbors []*Particle
}

func randomPos() float64 {
	var randValue float64
	rand.Seed(time.Now().UnixNano())
	randValue = STT + (END-STT)*rand.Float64()
	return randValue
}

func (p *Particle) Initialize(dim, population int) {
	p.Pos = make([]float64, dim)
	p.Vel = make([]float64, dim)
	p.Pbest = make([]float64, dim)
	p.Lbest = make([]float64, dim)

	for i, _ := range p.Pos {
		tmp := randomPos()
		p.Pos[i] = tmp
	}
	copy(p.Pbest, p.Pos)
}

func evaluate(pos []float64) float64 {
	var result float64 = 0.0
	for _, x := range pos {
		result += x * x
	}
	return result
}

func advance(p Particle) ([]float64, []float64) {
	dim := len(p.Pos)
	rho1 := 0.14 * rand.Float64()
	rho2 := 0.14 * rand.Float64()
	newPos := make([]float64, dim)
	newVel := make([]float64, dim)

	for i := 0; i < dim; i++ {
		newPos[i] = p.Pos[i] + p.Vel[i]
		newVel[i] = W * p.Vel[i]
		newVel[i] += rho1 * (p.Pbest[i] - p.Pos[i])
		newVel[i] += rho2 * (p.Lbest[i] - p.Pos[i])
	}

	return newPos, newVel
}

func main() {
	var swarm []Particle
	var bestParticle *Particle

	/**************/
	/* Initialize */
	/**************/
	swarm = make([]Particle, POPULATION)
	for i := range swarm {
		p := Particle{}
		p.Initialize(DIM, POPULATION)
		p.Fitness = evaluate(p.Pos)
		swarm[i] = p
	}
	// Add Neighbors : Complete Graph
	for i := range swarm {
		for j := range swarm {
			if i != j {
				swarm[i].Neighbors = append(swarm[i].Neighbors, &swarm[j])
			}
		}
	}

	// Pick up the particle which has the best fitness
	bestParticle = &swarm[0]
	for i := range swarm {
		if swarm[i].Fitness < bestParticle.Fitness {
			bestParticle = &swarm[i]
		}
	}
	for i := range swarm {
		for j := range swarm[i].Lbest {
			swarm[i].Lbest[j] = bestParticle.Lbest[j]
		}
	}

	/*************/
	/* Main Loop */
	/*************/
	for n := 0; n < LOOPMAX; n++ {
		// Update Particle Position and Velocity
		for i, p := range swarm {
			swarm[i].Pos, swarm[i].Vel = advance(p)
		}

		// Update Parsonal Best
		for i, p := range swarm {
			fitness := evaluate(p.Pos)
			swarm[i].Fitness = fitness
			pbestFitness := evaluate(p.Pbest)
			if fitness < pbestFitness {
				for j := range swarm[i].Pos {
					swarm[i].Pbest[j] = swarm[i].Pos[j]
				}
			}
		}

		// Update Local Best
		bestParticle = &swarm[0]
		for i := range swarm {
			if swarm[i].Fitness < bestParticle.Fitness {
				bestParticle = &swarm[i]
			}
		}
		for i := range swarm {
			for j := range swarm[i].Lbest {
				swarm[i].Lbest[j] = bestParticle.Lbest[j]
			}
		}

		// Output
		fmt.Println("Best Fitness :", bestParticle.Fitness)
	}

}
