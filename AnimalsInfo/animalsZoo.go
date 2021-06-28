package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*Write a program which allows the user to create a set of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake. With each command,
the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user.
Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
The following table contains the three types of animals and their associated data.
+--------+------------+-------------------+--------------+
| Animal | Food eaten | Locomotion method | Spoken sound |
+--------+------------+-------------------+--------------+
| cow    | grass      | walk              | moo          |
| bird   | worms      | fly               | peep         |
| snake  | mice       | slither           | hsss         |

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
Your program should continue in this loop forever.
Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal.
The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type.
Your program should call the appropriate method when the user issues a query command.
*/

/*This assignment is worth 10 points.

Test the program by running it and issuing three newanimal commands followed by three query commands.
Each new animal should involve a different animal type (cow, bird, snake), each with a different name.
Each query should involve a different animal and a different property of the animal (eat, move, speak).
Give 2 points for each query for which the program provides the correct response.

Examine the code. If the code contains an interface type called Animal,
which is a struct containing three fields, all of which are strings, then give another 2 points.
If the program contains three types – Cow, Bird, and Snake – which all satisfy the Animal interface, give another 2 points.
*/

type Cow struct {
	name string
}

func (c Cow) getName() string {
	return c.name
}

type Snake struct {
	name string
}

func (s Snake) getName() string {
	return s.name
}

type Bird struct {
	name string
}

func (b Bird) getName() string {
	return b.name
}

type Animal interface {
	eat()
	move()
	speak()
	getName() string
}

func (s Snake) eat() {
	fmt.Printf("%s eats mice\n", s.name)
}

func (b Bird) eat() {
	fmt.Printf("%s eats worms\n", b.name)
}

func (c Cow) eat() {
	fmt.Printf("%s eats grass\n", c.name)
}

func (s Snake) move() {
	fmt.Printf("%s slither\n", s.name)
}

func (b Bird) move() {
	fmt.Printf("%s flies\n", b.name)
}

func (c Cow) move() {
	fmt.Printf("%s walks\n", c.name)
}

func (s Snake) speak() {
	fmt.Printf("%s hsss\n", s.name)
}

func (b Bird) speak() {
	fmt.Printf("%s peep\n", b.name)
}

func (c Cow) speak() {
	fmt.Printf("%s moo\n", c.name)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	var sliceAnimals []Animal
	for {
		fmt.Println("Welcome to our ZOO where you can create and query animals")
		fmt.Println("You can create a new animal with following Syntax: newanimal (name of animal) (type of animal)")
		fmt.Println("Other option, you can query the animal(s) with following Syntax: query (name of animal) (action of animal)")
		fmt.Print(">")
		choice, n, p := readAnimalsCommandLine(reader)
		switch choice {
		case "newanimal":
			sliceAnimals = addNewAnimal(p, sliceAnimals, n)
		case "query":
			fmt.Printf("%+v \n", sliceAnimals)
			found := findAnimalIndexIfPresent(sliceAnimals, n)
			if found != -1 {
				switch p {
				case "speak":
					sliceAnimals[found].speak()
				case "eat":
					sliceAnimals[found].eat()
				case "move":
					sliceAnimals[found].move()
				default:
					fmt.Printf("%s This query command has not been recognised. Please check the spelling and try again!\n!")
				}
			} else {
				fmt.Printf("The animal with name %s was not found, please check spell or create a new one\n")
			}
		default:
			fmt.Println("Action command not recognised please check the spelling and try again ('newanimal' or 'query')!")
		}
		fmt.Println("To Exit from this loop please use Control+C")
	}

}

func findAnimalIndexIfPresent(sliceAnimals []Animal, n string) int {
	for index := range sliceAnimals {
		if sliceAnimals[index].getName() == n {
			return index
		}
	}
	return -1
}

func addNewAnimal(p string, sliceAnimals []Animal, n string) []Animal {
	switch p {
	case "cow":
		sliceAnimals = append(sliceAnimals, Cow{name: n})
		fmt.Println("Created it!")
	case "snake":
		sliceAnimals = append(sliceAnimals, Snake{name: n})
		fmt.Println("Created it!")
	case "bird":
		sliceAnimals = append(sliceAnimals, Bird{name: n})
		fmt.Println("Created it!")
	default:
		fmt.Println("Cannot be created an animal for this type, is not accepted!")
	}
	return sliceAnimals
}

func readAnimalsCommandLine(reader *bufio.Reader) (string, string, string) {
	line, err := reader.ReadString('\n')
	checkError(err)
	line = strings.Replace(line, "\n", "", -1)
	checkEmpty(line)
	text := strings.Split(line, " ")
	if len(text) < 3 || len(text) > 3 {
		log.Fatal("Error: you must enter 3 words, Please try again!")
	}
	action, name, param := strings.ToLower(text[0]), strings.ToLower(text[1]), strings.ToLower(text[2])
	log.Printf("action: %s and as name=%s and param=%s \n", action, name, param)

	return action, name, param
}

func checkEmpty(text string) {
	if len(text) == 0 {
		log.Fatal("Error, cannot enter empty text")
	}
}

func checkError(err interface{}) {
	if err != nil {
		log.Fatal("Wrong input, with error:", err)
	}
}
