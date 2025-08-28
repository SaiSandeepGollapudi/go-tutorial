# go-tutorial

# method
1. Go does not have classes like java. However, we can define methods on structs/types.

A method is a function where we give special receiver the struct as argument. 

The receiver appears in its own argument list between the func keyword and the method name.

# Functions
A function can take zero or more arguments. for arguments the data type comes after the variable name. When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

We mention data type of function like int.

# goroutines (Go’s lightweight threads)
go f(x, y, z) starts a new goroutine running

The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.

Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives.

# Channels
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:
ch := make(chan int)

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

# Buffered Channels
Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

ch := make(chan int, 100)

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

# WaitGroups
To wait for multiple goroutines to finish, we can use a wait group.

It functions as a counter that is incremented when a goroutine is launched and decremented when a goroutine completes its task. The Wait() method blocks the execution of the calling goroutine until the internal counter of the WaitGroup reaches zero, indicating that all registered goroutines have finished.

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