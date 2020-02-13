package main

import "fmt"

// Notation: P∧Q
// Meaning:  P and Q(both true)
// Example:  x≤x+1∧x≠x+1
func PAndQ() {
	var x int
	P := func(x int) bool { return x <= x+1 }
	Q := func(x int) bool { return x != x+1 }
	x = 5
	r := P(x) && Q(x)
	fmt.Printf("\nP∧Q, x≤x+1 ∧ x≠x+1, x=%v, result %v", x, r)
}

// Notation: P∨Q
// Meaning:  P or Q(one or both true)
// Example:  x≤y∨y≤x
func POrQ() {
	var x, y int
	P := func(x, y int) bool { return x <= y }
	Q := func(x, y int) bool { return y <= x }
	x = 5
	y = 6
	r := P(x, y) || Q(x, y)
	fmt.Printf("\nP∨Q, x≤y ∨ y≤x, x = %v, y = %v, result %v", x, y, r)
}

// Notation: ¬P
// Meaning:  not P(P is not true)
// Example:  ¬3≥5
func NotP() {
	var x, y int
	P := func(x, y int) bool { return x >= y }
	x = 3
	y = 5
	r := !P(x, y)
	fmt.Printf("\n¬P, ¬ x≥y, x = %v, y = %v, result %v", x, y, r)
}

// Notation: P ⇒ Q
// Meaning:  if P then Q
// Example:  x<y ⇒ x≤y
func IfPThenQ() {
	var x, y int
	P := func(x, y int) bool { return x < y }
	Q := func(x, y int) bool { return x <= y }
	x = 3
	y = 5
	if P(x, y) {
		r := Q(x, y)
		fmt.Printf("\nP ⇒ Q, x<y ⇒ x≤y, x = %v, y = %v, result %v", x, y, r)
	}
}

// P ≡ Q
// Meaning: P if and only if Q
// Example: x<y ≡ y>x
func PIfAndOnlyIfQ() {
	var x, y int
	P := func(x, y int) bool { return x < y }
	Q := func(x, y int) bool { return y > x }
	x = 3
	y = 5

	r := Q(x, y) && P(x, y)

	if r {
	}

	fmt.Printf("\nP ≡ Q, x<y ≡ y > x, x = %v, y = %v, result %v", x, y, r)
}

// Notation: ∃x•P
// Meaning:  there exists an x such that P
// Example:  ∃x•x>y
func ThereExistsAnxSuchThatP(domain []int) {
	var D []int
	if len(domain) != 0 {
		D = domain
	} else {
		D = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	var y int

	P := func(x, y int) bool { return x > y }

	y = 8

	for _, x := range D {
		r := P(x, y)
		if r {
			fmt.Printf("\n∃x•P, ∃x•x>y, Domain %v,  x = %v, y = %v, result %v", D, x, y, r)
			return
		}
	}
}

// Notation: ∀x•P
// Meaning:  for all x,P
// Example:  ∀x•x<x+1
func ForAllxP(domain []int) {
	var D []int

	if len(domain) != 0 {
		D = domain
	} else {
		D = []int{1, 2, 3}
	}

	P := func(x int) bool { return x < x+1 }

	for _, x := range D {
		r := P(x)
		if r {
			fmt.Printf("\n∀x•P, ∀x•x<x+1,  Domain %v,  x = %v,  result %v", D, x, r)
		}
	}
}

// Notation: ∃x:A•P
// Meaning:  there exists an x in set A such that P
func ThereExistsAnxInSetASuchThatP() {
	A := []int{1, 2, 3, 4, 5, 6, 7}
	ThereExistsAnxSuchThatP(A)
}

// Notation: ∀x:A•P
// Meaning:  for all x in set A,P
func ForAllxInSetAP() {
	A := []int{1, 2, 3, 4, 5, 6, 7}
	ForAllxP(A)
}

func main() {
	PAndQ()
	POrQ()
	NotP()
	IfPThenQ()
	PIfAndOnlyIfQ()
	ThereExistsAnxSuchThatP([]int{})
	ForAllxP([]int{})
	ThereExistsAnxInSetASuchThatP()
	ForAllxInSetAP()
}
