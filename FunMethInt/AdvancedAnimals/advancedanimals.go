/* -----
This is a coursework, an exercise of interfaces. The program allows to create animals
with name (command 'newanimal <name> <animal>'), and retrieve certain information about
them via command line (command 'query <name> <information>'). The possible animals are
cow, bird and snake which are implemented as different classes, having the interface Animal.
Information can be eat, move or speak.
----- */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	animals := make([]Animal, 0, 10)

	fmt.Println("Hi. Use me for animal information storage.")

	for {
		commands := AskForInput()

		switch commands[0] {
		case "leave":
			fmt.Println("Bye.")
			time.Sleep(8e8)
			os.Exit(0)
		case "newanimal":
			switch commands[2] {
			case "cow":
				animals = append(animals, Cow{name: commands[1]})
			case "bird":
				animals = append(animals, Bird{commands[1]})
			case "snake":
				animals = append(animals, Snake{commands[1]})
			}
		case "query":
			var theAnimal Animal
			for _, v := range animals {
				if v.IsItsName(commands[1]) {
					theAnimal = v
				}
			}

			if theAnimal != nil {
				switch commands[2] {
				case "eat":
					theAnimal.Eat()
				case "move":
					theAnimal.Move()
				case "speak":
					theAnimal.Speak()
				}
			} else {
				fmt.Printf("I have no information about %s.\n", commands[1])
			}
		}
	}
}

type Animal interface {
	Eat()
	Speak()
	Move()
	IsItsName(string) bool
}

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

// Eat prints out the cow's food.
func (c Cow) Eat() {
	fmt.Printf("%s the cow eats grass.\n", c.name)
}

// Move prints out the cow's way of movement.
func (c Cow) Move() {
	fmt.Printf("%s the cow walks.\n", c.name)
}

// Speak prints out the cow's native spoken language.
func (c Cow) Speak() {
	fmt.Printf("%s the cow says moo.\n", c.name)
}

// IsItsName determines if its input is the name of the cow.
func (c Cow) IsItsName(name string) bool {
	return c.name == name
}

// Eat prints out the bird's food.
func (b Bird) Eat() {
	fmt.Printf("%s the bird eats worms.\n", b.name)
}

// Move prints out the bird's way of movement.
func (b Bird) Move() {
	fmt.Printf("%s the bird flies.\n", b.name)
}

// Speak prints out the bird's native spoken language.
func (b Bird) Speak() {
	fmt.Printf("%s the bird says peep.\n", b.name)
}

// IsItsName determines if its input is the name of the bird.
func (b Bird) IsItsName(name string) bool {
	return b.name == name
}

// Eat prints out the snake's food.
func (s Snake) Eat() {
	fmt.Printf("%s the snake eats mice.\n", s.name)
}

// Move prints out the snake's way of movement.
func (s Snake) Move() {
	fmt.Printf("%s the snake slithers.\n", s.name)
}

// Speak prints out the snake's native spoken language.
func (s Snake) Speak() {
	fmt.Printf("%s the snake says hsss.\n", s.name)
}

// IsItsName determines if its input is the name of the snake.
func (s Snake) IsItsName(name string) bool {
	return s.name == name
}

// AskForInput prompts the user for commands. Input loops until a correct command
// is given. AskForInput uses ParseInput to process each input.
func AskForInput() []string {
	command := make([]string, 3)
	scanner := bufio.NewScanner(os.Stdin)
	response := errors.New("Dummy error")

	for response != nil {
		fmt.Print("> ")
		scanner.Scan()

		response = scanner.Err()
		if response != nil {
			fmt.Println("Something went wrong, I couldn't get your command.\n", response)
			continue
		}

		command, response = ParseInput(scanner.Text())
		if response != nil {
			fmt.Println(response)
			continue
		}

		response = nil
	}
	return command
}

// ParseInput takes the raw input and splits it to a slice of command tokens. If the
// input is invalid (too short or contain invalid tokens), the function returns an
// appropriate error. The input is converted to lower case.
func ParseInput(input string) ([]string, error) {
	// acceptedL1Tokens := map[string]bool{"newanimal": true, "query": true, "leave": true}
	acceptedNewTokens := map[string]bool{"cow": true, "bird": true, "snake": true}
	acceptedQueryTokens := map[string]bool{"eat": true, "move": true, "speak": true}
	var returnResponse error = nil

	inputTokens := strings.Fields(input)

	responses := CreateResponses(inputTokens)

	if len(inputTokens) > 3 {
		fmt.Printf("Your command has too many words. I ignore '%s'.\n",
			strings.Join(inputTokens[3:], " "))
		inputTokens = inputTokens[:3]
	}

	switch len(inputTokens) {
	case 0:
		returnResponse = responses["noneGivenErr"]
	case 1:
		switch inputTokens[0] {
		case "leave":
		case "query":
			returnResponse = responses["unspecifiedNameAttrErr"]
		case "newanimal":
			returnResponse = responses["unspecifiedNameAnimalErr"]
		default:
			returnResponse = responses["unknownCommandErr"]
		}
	case 2:
		switch inputTokens[0] {
		case "query":
			returnResponse = responses["unspecifiedAttributeErr"]
		case "newanimal":
			returnResponse = responses["unspecifiedAnimalErr"]
		default:
			returnResponse = responses["unknownCommandErr"]
		}
	case 3:
		if inputTokens[0] == "newanimal" && !acceptedNewTokens[inputTokens[2]] {
			returnResponse = responses["wrongAnimalErr"]
		} else if inputTokens[0] == "query" && !acceptedQueryTokens[inputTokens[2]] {
			returnResponse = responses["wrongAttributeErr"]
		} else if inputTokens[0] != "newanimal" && inputTokens[0] != "query" {
			returnResponse = responses["unknownCommandErr"]
		}
	default:
		returnResponse = responses["unknownCommandErr"]
	}

	return inputTokens, returnResponse
}

// CreateResponses creates a hash of response names --> responses used in input parsing.
func CreateResponses(inputTokens []string) map[string]error {
	responseHash := make(map[string]error)

	responseHash["noneGivenErr"] =
		errors.New("You've given no command... If you want to leave, just type 'leave'.")
	responseHash["unknownCommandErr"] =
		fmt.Errorf("What do you mean by %s?", strings.Join(inputTokens, " "))
	responseHash["unspecifiedAttributeErr"] =
		errors.New("Please specify an attribute.")
	responseHash["unspecifiedNameAttrErr"] =
		errors.New("Please specify a name and an attribute.")
	responseHash["unspecifiedAnimalErr"] =
		errors.New("What animal is it?")
	responseHash["unspecifiedNameAnimalErr"] =
		errors.New("Please specify the name and the animal.")
	responseHash["wrongAnimalErr"] =
		errors.New("I don't know this animal.")
	responseHash["wrongAttributeErr"] =
		errors.New("I don't know this attribute.")

	return responseHash
}
