package main

import (
	"fmt"

	"github.com/mbottini/newton/polynomial"
	"github.com/mbottini/newton/solver"
)

const eps = 0.000001
const maxIters = 25

func main() {
	var p polynomial.Polynomial
	var guess complex128
	var results []complex128
	p = append(p, 1, 0, 0, 0, 0, 0, 1)

	fmt.Println(p)
	fmt.Println(p.Derivative())

	for p.Degree() > 0 {
		var err error
		var result complex128
		for p.Degree() > 0 {
			fmt.Printf("Input a guess. ")
			fmt.Scanf("%g", &guess)
			for err == nil && p.Degree() > 0 {
				result, err = solver.SolvePolynomial(p, guess, eps, maxIters)
				if err == nil {
					fmt.Printf("Found root at %v\n", result)
					results = append(results, result)
					p, _ = p.DivideByTerm(result)
				} else {
					fmt.Println("Polynomial did not converge with that guess.")
				}
			}
		}
	}

	fmt.Println(results)
}
