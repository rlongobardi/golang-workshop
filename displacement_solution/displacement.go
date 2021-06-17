/*Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))*/

/*Test the program by running it twice and
testing it with two different sets of values for acceleration, initial velocity,
initial displacement, and time. Give 3 points if the program works correctly
for one set of values, and give 3 more points if the program works correctly
for the other set of values.

Examine the code. If the code contains a
function called GenDisplaceFn()
which takes three float64 arguments, acceleration a, initial velocity vo,
and initial displacement so and returns a function, then give
another 2 points. If the function returned by GenDisplaceFn() is used to compute the time, give another 2 points.

*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please, enter acceleration a=\n")
	a, err1 := getInputFromCmdLineToFloat64(reader)
	checkError(err1)
	fmt.Printf("Please enter initial velocity\n")
	vt0, err2 := getInputFromCmdLineToFloat64(reader)
	checkError(err2)
	fmt.Printf("Please enter initial displacement d0=\n")
	d0, err3 := getInputFromCmdLineToFloat64(reader)
	checkError(err3)

	fn := GenDisplaceFn(a, vt0, d0)
	//after 3 seconds
	fmt.Println("Displacement after 3 seconds is:", fn(3))
	//after 5 seconds
	fmt.Println("Displacement after 5 seconds is:", fn(5))

}

func getInputFromCmdLineToFloat64(reader *bufio.Reader) (float64, error) {
	line, err := reader.ReadString('\n')
	checkError(err)
	line = strings.Replace(line, "\n", "", -1)
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		log.Fatal("Error, cannot enter empty text")
	}
	val, err1 := strconv.ParseFloat(line, 64)
	return val, err1
}

func checkError(err interface{}) {
	if err != nil {
		log.Fatal("Wrong input, with error:", err)
	}
}

func GenDisplaceFn(a, vot, so float64) func(t0 float64) float64 {
	// (1/2) * a * t^2 + vot + so
	fn := func(t float64) float64 {
		return math.Abs((0.5 * a * (t * t)) + vot + so)
	}
	return fn
}
