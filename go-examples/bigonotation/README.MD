# Big O Notation

The Big O notation is a way of describing how much a given algorithm will consume computational resources as its inputs increase. As a computational resource we can cite processing time and memory consumption.

This annotation is the best known within asymptotic analysis, representing the worst case, but there are two other types known as Big Θ (theta) and Big Ω (Omega).

## Calculating time complexity

To define the complexity of an algorithm it is possible to follow the following steps:

Let's use this function as an example

```go
func Sum(array []int) int {
	var total int
	for i := 0; i < len(array) ; i++ {
	    total += array[i]	
    }
}
```

**Step 1 - Define the complexity of each operation**

```go
func Sum(array []int) int {
	var total int // O(1)
	for i := 0; i < len(array) ; i++ {
	    total += array[i]  // n * O(1), given that this line will be executed n times, where n = len(array) - 1
    }
	return total // O(1) 
}
```

**Step 2 - Describe the sets of complexities in function form**

Based on the previous analysis, adding all the terms and transforming them into a T function, we have:
```
T(n) = O(1) + n * O(1) + O(1)
```

For a clearer view, let's transform O(1) into letters, representing constants, then:

```
T(n) = a*n + b + c
```

**Step 3 - Find the term with the highest growth factor**

Among the three terms of our function, we have 2 constants (b and c) and a value that grows based on `n` (a*n), which is then our term with the highest growth factor.

**Step 4 - Disregard the coefficients/constants**

Since `a*n` is our most relevant term, disregarding the coefficient `a`, we can define this function with complexity `O(n)`.

## Common notations

O(1) - Constant <br />
O(log n) - Logarithmic  <br />
O(n) - Linear  <br />
O(n log n) - Linearithmic  <br />
O(n^2) - Quadratic  <br />
O(n^3) - Cubic  <br />
O(2^n) - Exponential  <br />
O(n!) - Factorial  <br />
O(n^n) - Polynomial  <br />