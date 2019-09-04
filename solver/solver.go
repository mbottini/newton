package solver

import (
	"errors"
	"fmt"
	"math/cmplx"

	"github.com/mbottini/newton/polynomial"
)

// FixedPoint takes a function f, a guess, and an epsilon and repeatedly applies
// f until either the function converges within eps or runs for more than
// maxIters iterations.
func fixedPoint(f func(complex128) complex128,
	guess complex128,
	eps float64,
	maxIters int) (complex128, error) {
	newGuess := f(guess)
	currentIter := 0
	for cmplx.Abs(newGuess-guess) > eps && currentIter < maxIters {
		guess = newGuess
		newGuess = f(guess)
		currentIter++
	}
	if cmplx.Abs(newGuess-guess) < eps {
		return newGuess, nil
	}
	return 0, errors.New("solver: FixedPoint did not converge")
}

// NewtonFunc takes a Polynomial and returns a function that applies Newton's
// Method: x_1 = x_0 - multiplicity * f(x) / f'(x).
// Multiplicity is used when the polynomial has multiple instances of the same
// root.
func newtonFunc(p polynomial.Polynomial, multiplicity int) func(complex128) complex128 {
	return func(x complex128) complex128 {
		f := p.Eval()
		g := p.Derivative().Eval()
		return x - complex(float64(multiplicity), 0)*f(x)/g(x)
	}
}

// SolvePolynomial takes a Polynomial, an initial guess, an epsilon, and a
// maximum number of iterations before the function returns an error. It
// performs Newton's Method on the polynomial and guess to find a single root.
func SolvePolynomial(p polynomial.Polynomial,
	guess complex128,
	eps float64,
	maxIters int) (complex128, error) {
	// Stupid edge case - what if the guess is right on the solution, but the
	// solution's derivative is 0?
	deriv := p.Derivative().Eval()
	if deriv(guess) == 0 {
		if p.Eval()(guess) == 0 {
			return guess, nil
		}
	}
	// Even iterating through
	possibleMultiplicity := p.Degree()
	fmt.Printf("degree = %d\n", p.Degree())
	for i := 1; i <= possibleMultiplicity; i++ {
		f := newtonFunc(p, i)
		result, err := fixedPoint(f, guess, eps, maxIters)
		if err == nil {
			return result, err
		}
	}
	return 0, errors.New("solver: polynomial did not converge")
}
