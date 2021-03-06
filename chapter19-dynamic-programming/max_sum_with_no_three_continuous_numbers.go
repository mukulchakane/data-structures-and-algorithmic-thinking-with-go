// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

package main

import (
	"fmt"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxSumWithNoThreeContinuousNumbers(A []int) int {
	n := len(A)
	M := make([]int, n+1)
	M[0] = A[0]
	if n >= 1 {
		M[0] = A[0]
	}
	if n >= 2 {
		M[1] = A[0] + A[1]
	}

	if n > 2 {
		M[2] = max(M[1], max(A[1]+A[2], A[0]+A[2]))
	}

	for i := 3; i < n; i++ {
		M[i] = max(max(M[i-1], M[i-2]+A[i]), A[i]+A[i-1]+M[i-3])
	}
	return M[n-1]
}

func main() {
	fmt.Println(maxSumWithNoThreeContinuousNumbers([]int{1, 2, 9, 4, 5, 0, 4, 11, 6}))
	fmt.Println(maxSumWithNoThreeContinuousNumbers([]int{100, 1000, 100, 1000, 1}))
}

