package polynomial

import "fmt"

// Polynomial is just a list of coefficients. The 0th coefficient is the
// constant term, the 1st coefficient is the linear term, etc.
type Polynomial []complex128

func (p Polynomial) equalOrHigherDegree(other Polynomial) bool {
	return len(p) >= len(other)
}

// Trim removes extra trailing 0 terms from the Polynomial.
func (p Polynomial) Trim() Polynomial {
	for i := range p {
		if p[len(p)-i-1] != 0 {
			return p[0 : len(p)-i]
		}
	}
	// We keep a 0 constant term if the polynomial really is 0.
	return p[0:1]
}

// Degree is the degree of the highest coefficient in the Polynomial.
func (p Polynomial) Degree() int {
	return len(p) - 1
}

func (p Polynomial) String() string {
	polyMax := len(p) - 1
	resultStr := ""
	for i := range p {
		power := polyMax - i
		if power != 0 {
			resultStr += fmt.Sprintf("%gx^%d + ", p[power], power)
		} else {
			resultStr += fmt.Sprintf("%g", p[power])
		}
	}
	return resultStr
}

// Negate takes the negative of every coefficient in the Polynomial.
func (p Polynomial) Negate() Polynomial {
	var result Polynomial
	for _, coefficient := range p {
		result = append(result, -coefficient)
	}
	return result
}

// Add adds two Polynomials together and returns another Polynomial.
func (p Polynomial) Add(other Polynomial) Polynomial {
	var result Polynomial
	var maxIndex int
	var currentCoef complex128
	if len(p) > len(other) {
		maxIndex = len(p)
	} else {
		maxIndex = len(other)
	}
	for i := 0; i < maxIndex; i++ {
		currentCoef = 0
		if i < len(p) {
			currentCoef += p[i]
		}
		if i < len(other) {
			currentCoef += other[i]
		}
		result = append(result, currentCoef)
	}
	return result.Trim()
}

// Sub subtracts the other Polynomial from this Polynomial.
func (p Polynomial) Sub(other Polynomial) Polynomial {
	return p.Add(other.Negate())
}

// MulScalar multiplies each coefficient in the Polynomial by scalar.
func (p Polynomial) MulScalar(scalar complex128) Polynomial {
	var result Polynomial
	for _, coefficient := range p {
		result = append(result, coefficient*scalar)
	}
	return result.Trim()
}

// Mul multiplies two Polynomials together and returns a Polynomial.
func (p Polynomial) Mul(other Polynomial) Polynomial {
	result := make(Polynomial, p.Degree()+other.Degree()+1)
	for i, thisCoef := range p {
		for j, otherCoef := range other {
			result[i+j] += thisCoef * otherCoef
		}
	}
	return result
}

// Div divides this Polynomial by the Other Polynomial. It returns a tuple.
// The first term is the quotient, the second term is the remainder. This is
// done really inefficiently, as it works in terms of the other polynomial
// operations of multiplication and subtraction.
func (p Polynomial) Div(other Polynomial) (Polynomial, Polynomial) {
	mutable := make(Polynomial, len(p))
	copy(mutable, p)
	result := make(Polynomial, p.Degree()-other.Degree()+1)
	for mutable.equalOrHigherDegree(other) {
		power := mutable.Degree() - other.Degree()
		factor := mutable[len(mutable)-1] / other[len(other)-1]
		result[power] = factor
		subtractand := make(Polynomial, power+1)
		subtractand[power] = factor
		subtractand = other.Mul(subtractand)
		mutable = mutable.Sub(subtractand)
	}
	return result, mutable
}
