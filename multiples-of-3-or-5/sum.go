package main

func SumMultiplesOf3And5(number int) int {

	largestNFor3 := getLargestNForADivisor(number-1, 3)
	largestNFor5 := getLargestNForADivisor(number-1, 5)
	largestNFor15 := getLargestNForADivisor(number-1, 15)

	nthTermFor3 := getNthTermForAnInteger(largestNFor3, 3)
	nthTermFor5 := getNthTermForAnInteger(largestNFor5, 5)
	nthTermFor15 := getNthTermForAnInteger(largestNFor15, 15)

	sumOfNTermsFor3 := getSumOfNTerms(3, nthTermFor3, largestNFor3)
	sumOfNTermsFor5 := getSumOfNTerms(5, nthTermFor5, largestNFor5)
	sumOfNTermsFor15 := getSumOfNTerms(15, nthTermFor15, largestNFor15)

	return sumOfNTermsFor3 + sumOfNTermsFor5 - sumOfNTermsFor15
}

func getLargestNForADivisor(number int, divisor int) int {
	// a = first term
	// n = number
	// d = difference between numbers
	// a + (n - 1)d = x
	// n = (x - a) / d + 1

	if number < divisor {
		return 0
	}

	n := (number-divisor)/divisor + 1

	return n
}

func getNthTermForAnInteger(n int, integer int) int {
	// l = a + (n - 1)d
	// l = nth term
	// a = first term
	// d = difference between terms

	nthTerm := integer + (n-1)*integer

	return nthTerm
}

func getSumOfNTerms(firstTerm int, nthTerm int, n int) int {
	// s = (n / 2) * (a + l)
	// s = sum of N terms
	// a = first term
	// l = nth term

	sumOfNTerms := int((float64(n) / 2) * (float64(firstTerm) + float64(nthTerm)))

	return sumOfNTerms
}
