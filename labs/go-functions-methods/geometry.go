// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 156.

// Package geometry defines simple types for plane geometry.
//!+point
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Point ...
type Point struct{ x, y, angle float64 }

// X ...
func (p Point) X() float64 {
	return p.x
}

// Y ...
func (p Point) Y() float64 {
	return p.y
}

// Distance traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X()-p.X(), q.Y()-p.Y())
}

// Distance same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X()-p.X(), q.Y()-p.Y())
}

//!-point

//!+path

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	sum += path[len(path)-1].Distance(path[0])
	return sum
}

//!-path

// SortByAngle ...
func (path Path) SortByAngle(minPoint Point) Path {
	angles := Path{}
	for _, dot := range path {
		co := dot.Y() - minPoint.Y()
		ca := dot.X() - minPoint.X()
		dot.angle = math.Atan2(co, ca)
		angles = append(angles, dot)
	}
	for i := 0; i < len(angles); i++ {
		for j := i; j < len(angles); j++ {
			if angles[j].angle < angles[i].angle {
				aux := angles[j]
				angles[j] = angles[i]
				angles[i] = aux
			}
		}
	}
	return angles
}

// RandomGenerator ...
func RandomGenerator(sides int64) (Path, Point) {
	planeRange := [2]float64{-100, 100}
	vertex := Path{}
	minY := planeRange[1] + 1
	minPoint := Point{}
	rand.Seed(time.Now().UTC().UnixNano())

	for i := int64(0); i < sides; i++ {
		x := rand.Float64()*(planeRange[1]-planeRange[0]) + planeRange[0]
		y := rand.Float64()*(planeRange[1]-planeRange[0]) + planeRange[0]
		point := Point{x, y, 0}
		vertex = append(vertex, point)
		if y < minY {
			minY = y
			minPoint = point
		}
	}
	return vertex, minPoint
}

func main() {
	sides, err := strconv.ParseInt(os.Args[1:2][0], 10, 64)
	if err != nil {
		return
	}

	vertex, minPoint := RandomGenerator(sides)

	vertex = vertex.SortByAngle(minPoint)

	fmt.Println(vertex)
	fmt.Println(vertex.Distance())
}
