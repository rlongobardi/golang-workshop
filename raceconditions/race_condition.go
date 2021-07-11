package main

import (
	"fmt"
	"time"
)

/* TASK
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.
Students will receive full credit
if their goroutines produce a race condition
and they explain what the race condition is and why it occurs.
*/

//ANSWER
/*
By definition, a race condition is a condition of a program where its behavior depends on relative timing
or interleaving of multiple threads or processes. One or more possible outcomes may be undesirable, resulting in a bug.
We refer to this kind of behavior as nondeterministic.

The most common type of race condition is caused by the pattern check-then-act.
It’s defined by a program flow where a potentially stale observation is used to decide what to do next.

Another type of race condition is closely related to the data race.A data race occurs when two threads access the same variable concurrently, and at least one of the accesses is a write.
The data race concept is more specific to memory access in a particular concurrency model and, thus, varies across platforms.
Thread-safe is the term we use to describe a program, code, or data structure free of race conditions when accessed
by multiple threads.*/
//EXPLANATION : we will see closely this second case applied in the below example.
//Here a function that check if a number (from 1 to 20) is even or odd Then print and wait
//There is a shared counter for who performe faster get out the cycle
//Now the counter can be accessed by both goroutine/threads but only on goroutine/thread
//will be able to execute all the other will be left out to complete.

var counter = 0

func main() {
	fmt.Println("start")

	go MyRoutine("GoRoutine_2")
	go MyRoutine("GoRoutine_1")

	time.Sleep(time.Second) //otherwise the above goroutine won't be executed main thread ends .
	fmt.Println("done")
}

func MyRoutine(from string) {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(from, "print even number:", i)
			time.Sleep(20 * time.Millisecond)

		} else {
			fmt.Println(from, "print odd number:", i)
			time.Sleep(40 * time.Millisecond)
		}
		counter++

		if counter == 19 {
			break
		}
	}

}
