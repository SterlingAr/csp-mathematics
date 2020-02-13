package main

import "fmt"

// Notation: ∈
// Meaning:  is a member of
// Example:	 Napoleon ∈ mankind
func IsAMemberOf(m string) {
	A := []string{"a", "b", "c"}
	for _, x := range A {
		if x == m {
			fmt.Printf("\n∈, %v is a member of %v, result %v", m,A,true)
			return
		}
	}
}

// Notation: ∉
// Meaning:  is not a member of
// Example:  Napoleon ∉ mankind
func IsNotAMemberOf(m string) {
	A := []string{"a", "b", "c"}
	for _, x := range A {
		if x != m {
			fmt.Printf("\n∉, %v is not a member of %v, result %v", m,A,true)
			return
		}
	}
}

// Notation: {x|P(x)}
// Meaning:  the set of all x such that P(x)
// Example:  {E}={x|x % 2 = 0}
func TheSetOfAllxSuchThatP(n []int)  {
	var N []int
	if len(n) > 0 {
		N = n
	} else {
		N = []int{1,2,3,4,5,6,7,8,9}
	}
	var E []int
	for _, x := range N {
		if x % 2 == 0 {
			E = append(E, x)
		}
	}
	fmt.Printf("\n{x|P(x)}, {E}={x|x /2 = 0}, Domain = %v, E = %v", N, E)
}

// Notation: A∪B
// Meaning:  A union B
// Example:  A∪B={x|x∈A ∨ x∈B}
func AUnionB() {
	var A,B,C uint8
	A = 2
	B = 4
	C = A | B // 6
	fmt.Printf("\nA∪B, A∪B={x|x∈A ∨ x∈B}, A=%08b, B=%08b, result = %08b", A,B,C)
}

// Notation: A∩B
// Meaning:  A intersect B
// Example:  A∩B={x|x∈A ∧ x∈B}
func AIntersectB()  {
	var A,B,C uint8
	A = 9
	B = 5
	C = A & B
	fmt.Printf("\nA∩B, A∪B={x|x∈A ∨ x∈B}, A=%08b, B=%08b, result = %08b", A,B,C)
}

// Notation: A−B
// Meaning:  A minus B
// Example:  A−B={x|x∈A ∧ ¬x∈B}
func AMinusB()  {
	var A,B,C uint8
	A = 9
	B = 5
	C = A & (A ^ B)
	//C = A ^ B // 6
	fmt.Printf("\nA-B, A−B={x|x∈A ∧ ¬x∈B}, A=%08b, B=%08b, result = %08b", A,B,C)
}

// Notation: A ⊆ B
// Meaning:  A is contained in B ( A is a subset of B)
// Example:  A ⊆ B ≡ ∀x:A•x∈B
func AIsContainedInB(a, b uint8) uint8 {
	var A,B,C uint8
	var isSubset bool
	A = a
	B = b
	C = A & B

	if C == A {
		isSubset = true
	} else {
		isSubset = false
	}
	fmt.Printf("\nA ⊆ B, A ⊆ B ≡ ∀x:A•x∈B, A=%08b, B=%08b, result = %08b %v", A,B,C, isSubset)
	return C
}

// Notation: A ⊇ B
// Meaning:  A contains B
// Example:  A⊇B ≡ B⊆A
func AContainsB()  {
	var A,B,C uint8
	var contains bool
	A = 255
	B = 5
	C = AIsContainedInB(B, A)

	if C == B {
		contains = true
	} else {
		contains = false
	}
	fmt.Printf("\nA ⊇ B, A⊇B ≡ B⊆A, A=%08b, B=%08b, result = %08b %v", A,B,C, contains)
}

// Notation: {x:A|P(x)}
// Meaning    the set of x in A such that P(x)
func TheSetOfxInASuchThatP()  {
	TheSetOfAllxSuchThatP([]int{2,2,2,4,4,5,5,6,7,7,7,8,8})
}

func main()  {
	IsAMemberOf("a")
	IsNotAMemberOf("ç")
	TheSetOfAllxSuchThatP([]int{})
	AUnionB()
	AIntersectB()
	AMinusB()
	AIsContainedInB(55, 255)
	AIsContainedInB(230, 255)
	AContainsB()
	TheSetOfxInASuchThatP()
}