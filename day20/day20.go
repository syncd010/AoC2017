package main //package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	re := regexp.MustCompile("(-?[[:digit:]])+")
	for i, line := range input {
		matches := myAtoi(re.FindAllString(line, -1)...)
		if len(matches) != 9 {
			return errors.New("Malformed input at line " + strconv.Itoa(i+1))
		}
	}
	return nil
}

const (
	X = iota
	Y
	Z
)

// Vec3 saves a value in 3D
type Vec3 [3]int

// Manhattan distance for this coordinate
func (c Vec3) norm() int {
	return helpers.Abs(c[X]) + helpers.Abs(c[Y]) + helpers.Abs(c[Z])
}

func (c1 Vec3) equals(c2 Vec3) bool {
	return (c1[X] == c2[X]) && (c1[Y] == c2[Y]) && (c1[Z] == c2[Z])
}

// Particle represents a particle position, velocity and accel
type Particle struct {
	id                        int
	position, velocity, accel Vec3
}

func (p *Particle) move() {
	p.velocity[X] += p.accel[X]
	p.velocity[Y] += p.accel[Y]
	p.velocity[Z] += p.accel[Z]
	p.position[X] += p.velocity[X]
	p.position[Y] += p.velocity[Y]
	p.position[Z] += p.velocity[Z]
}

func myAtoi(str ...string) []int {
	res := make([]int, len(str))
	for i, s := range str {
		res[i], _ = strconv.Atoi(s)
	}
	return res
}

// Converts to an appropriate format
func convert(input []string) []*Particle {
	particles := make([]*Particle, len(input))

	re := regexp.MustCompile("(-?[[:digit:]])+")
	for i, line := range input {
		matches := myAtoi(re.FindAllString(line, -1)...)

		particles[i] = &Particle{
			id:       i,
			position: Vec3{X: matches[0], Y: matches[1], Z: matches[2]},
			velocity: Vec3{X: matches[3], Y: matches[4], Z: matches[5]},
			accel:    Vec3{X: matches[6], Y: matches[7], Z: matches[8]},
		}
	}
	return particles
}

func updateCollisions(particles []*Particle) (numCollisions int) {
	for i := 0; i < len(particles); i++ {
		p1 := particles[i]
		if p1 == nil {
			continue
		}
		collided := false
		for j := i + 1; j < len(particles); j++ {
			p2 := particles[j]
			if p2 == nil {
				continue
			}
			if p1.position.equals(p2.position) {
				particles[j] = nil
				numCollisions++
				collided = true
			}
		}
		if collided {
			particles[i] = nil
		}
	}
	return numCollisions
}

// Check if we should end
// All particles are moving in their final direction
// The farthest particles have velocity and accel bigger than the closer particles
func isSimulationOver(particles []*Particle) bool {
	for _, p := range particles {
		if p == nil {
			continue
		}
		if p.accel[X]*p.velocity[X] < 0 ||
			p.accel[Y]*p.velocity[Y] < 0 ||
			p.accel[Z]*p.velocity[Z] < 0 ||
			p.velocity[X]*p.position[X] < 0 ||
			p.velocity[Y]*p.position[Y] < 0 ||
			p.velocity[Z]*p.position[Z] < 0 {
			return false
		}
	}

	// All particles moving in their final direction
	sort.Slice(particles,
		func(i, j int) bool {
			if particles[i] == nil {
				return true
			}
			if particles[j] == nil {
				return false
			}
			return particles[i].position.norm() < particles[j].position.norm()
		})

	for i, j := 0, 1; j < len(particles); j++ {
		if particles[i] == nil {
			i++
			continue
		}
		if particles[j] == nil {
			continue
		}

		if particles[i].velocity.norm() > particles[j].velocity.norm() ||
			particles[i].accel.norm() > particles[j].accel.norm() {
			return false
		}
		i = j
	}
	return true
}

func simulate(input []*Particle, withColisions bool) []*Particle {
	particles := make([]*Particle, len(input))
	for i := range input {
		newPart := (*input[i])
		particles[i] = &newPart
	}

	for iter := 0; ; iter++ {
		for _, p := range particles {
			if p != nil {
				p.move()
			}
		}

		if withColisions {
			updateCollisions(particles)
		}

		if iter%100 == 0 && isSimulationOver(particles) {
			break
		}
	}
	return particles
}

func solvePart1Wrong(input []*Particle) int {
	// Get the particle with the minimum acceleration. If more than one, just choose the first one...
	minAccel := helpers.MaxInt
	var neareast int
	for i, p := range input {
		if p.accel.norm() < minAccel {
			minAccel = p.accel.norm()
			neareast = i
		}
	}
	return neareast
}

func solvePart1(input []*Particle) int {
	particles := simulate(input, false)
	minPos := helpers.MaxInt
	var neareast int
	for i, p := range particles {
		if p != nil && p.position.norm() < minPos {
			minPos = p.position.norm()
			neareast = i
		}
	}
	return particles[neareast].id
}

func solvePart2(input []*Particle) int {
	particles := simulate(input, true)
	n := 0
	for _, p := range particles {
		if p != nil {
			n++
		}
	}
	return n
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	convertedInput := convert(input)
	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(convertedInput))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(convertedInput))

}
