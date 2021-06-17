package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak.
The user can issue a request to find out one of three things about an animal:
1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks.
The following table contains the three animals and their associated data which should be hard-coded into your program.

+--------+------------+-------------------+--------------+
| Animal | Food eaten | Locomotion method | Spoken sound |
+--------+------------+-------------------+--------------+
| cow    | grass      | walk              | moo          |
| bird   | worms      | fly               | peep         |
| snake  | mice       | slither           | hsss         |

Your program should present the user with a prompt, “>”,to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.
*/

/*This assignment is worth a total of 10 points.

Test the program by running it and testing it by issuing three requests.
Each request should involve a different animal (cow, bird, snake) and a different property of the animal (eat, move, speak).
Give 2 points for each request for which the program provides the correct response.

Examine the code.
If the code contains a type called Animal, which is a struct containing three fields, all of which are strings,
then give another 2 points. If the program contains three methods called Eat(), Move(), and Speak(), and all of their receiver types are Animal,
give another 2 points.
*/

type Animal struct {
	foodEaten        string
	locomotionMethod string
	spokenSound      string
}

func (a *Animal) SetAnimal(food, locomotion, sound string) {
	a.foodEaten = food
	a.locomotionMethod = locomotion
	a.spokenSound = sound
}

func (a *Animal) Move() {
	fmt.Println(a.locomotionMethod)
}

func (a *Animal) Speak() {
	fmt.Println(a.spokenSound)
}

func (a *Animal) Eat() {
	fmt.Println(a.foodEaten)
}

func main() {

	var animal Animal
	reader := bufio.NewReader(os.Stdin)
	for {
		enquireAnimals(reader, animal)
		fmt.Println("To Exit from this loop please use Control+C")
	}

}

func enquireAnimals(reader *bufio.Reader, animal Animal) {
	fmt.Println("Please enter animal name and a command for the animal(eat/speak/move), only two words separate by single space!")
	animalName, command := getAnimalNameAndAnimalCommand(reader)
	name := strings.ToLower(animalName)
	switch {
	case name == "cow":
		animal.SetAnimal("grass", "walk", "moo")
		animalObeyCommand(command, animal)
	case name == "bird":
		animal.SetAnimal("words", "fly", "peep")
		animalObeyCommand(command, animal)
	case name == "snake":
		animal.SetAnimal("mice", "slither", "hss")
		animalObeyCommand(command, animal)
	default:
		fmt.Println("Animal type not found,please check spelling and try again!!")
	}
}

func animalObeyCommand(command string, animal Animal) {
	cmd := strings.ToLower(command)
	if cmd == "move" {
		animal.Move()
	}
	if cmd == "speak" {
		animal.Speak()
	}
	if cmd == "eat" {
		animal.Eat()
	}
}

func getAnimalNameAndAnimalCommand(reader *bufio.Reader) (string, string) {
	line, err := reader.ReadString('\n')
	checkError(err)
	line = strings.Replace(line, "\n", "", -1)
	checkEmpty(line)
	text := strings.Split(line, " ")
	if len(text) != 2 {
		log.Fatal("Error, too many words! Try again with only two animal name and command Move/Eat/Speak")
	}
	name, cmd := text[0], text[1]
	fmt.Printf("animal name=%s and as command=%s \n", name, cmd)

	return name, cmd
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
