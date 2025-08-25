# go-tutorial

# method
1. Go does not have classes like java. However, we can define methods on structs/types.

A method is a function where we give special receiver the struct as argument. 

The receiver appears in its own argument list between the func keyword and the method name.

# Functions
A function can take zero or more arguments. for arguments the data type comes after the variable name. When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

We mention data type of function like int.

# goroutines (Go’s lightweight threads)

1. Factorial

Description:
Calculates the factorial of a number n.

Concepts Used:
Loops, integer multiplication.

2. Palindrome

What it does:
Checks if a given string is a palindrome by reversing it and comparing with the original.

Concepts used:
Strings, loop, comparison.

3. Hello in Different Languages (with Goroutines)

What it does:
Prints “Hello” in English, Spanish, French, and German concurrently.

Concepts used:

Goroutines (go func())

Concurrency

time.Sleep to allow goroutines to finish before program exits.

Example output (order may vary due to concurrency):

4. Word Count

What it does:
Splits a given text into words and counts how many times each word appears.

Concepts used:

Maps (map[string]int)

strings.Fields to split text

Looping through words

5. Fibonacci

What it does:
Generates the Fibonacci sequence up to n terms.

Concepts used:
Loops
Multiple assignment (first, second = second, first+second)