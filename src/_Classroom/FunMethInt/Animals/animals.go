/* -----
This program is a class/object exercise. It maintains pieces of information about certain
animals, which can be queried via the standard input commands. The input is looped until
the user gives the 'leave' command. An animal information request has to be in the
'<animal> <information>' format. Valid animal values: cow, bird, snake. Valid information
values: eat, move, speak.
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
	// Note: pointer type needed, auto ref/deref does not work when object is in a map.
	animals := make(map[string]*Animal)
	var command []string

	animals["cow"] = NewAnimal("grass", "walk", "moo")
	animals["bird"] = NewAnimal("worms", "fly", "peep")
	animals["snake"] = NewAnimal("mice", "slither", "hsss")

	fmt.Println("Hi. Ask me about animals.")

	for {
		command = AskForInput()

		if command[0] == "leave" {
			fmt.Println("Bye.")
			time.Sleep(8e8)
			os.Exit(0)
		}

		switch command[1] {
		case "eat":
			animals[command[0]].Eat()
		case "move":
			animals[command[0]].Move()
		case "speak":
			animals[command[0]].Speak()
		default:
			os.Exit(-1)
		}
	}

}

// Animal is the datatype which stores different attributes of an animal.
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// NewAnimal creates a new instance of Animal with the data specified in the parameters.
func NewAnimal(food, locomotion, noise string) *Animal {
	a := new(Animal)
	a.food = food
	a.locomotion = locomotion
	a.noise = noise
	return a
}

// Eat prints out the animal's food.
func (a *Animal) Eat() {
	fmt.Printf("They eat %s.\n", a.food)
}

// Move prints out the animal's way of movement.
func (a *Animal) Move() {
	fmt.Printf("They %s.\n", a.locomotion)
}

// Speak prints out the animal's native spoken language.
func (a *Animal) Speak() {
	fmt.Printf("They say %s.\n", a.noise)
}

// AskForInput prompts the user for commands. Input loops until a correct command
// is given. AskForInput uses ParseInput to process each input.
func AskForInput() []string {
	command := make([]string, 2)
	scanner := bufio.NewScanner(os.Stdin)
	err := errors.New("Dummy error")

	for err != nil {
		fmt.Print("> ")
		scanner.Scan()

		err = scanner.Err()
		if err != nil {
			fmt.Println("Something went wrong, I couldn't get your command.\n", err)
			continue
		}

		command, err = ParseInput(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = nil
	}
	return command
}

// ParseInput takes the raw input and splits it to a slice of command tokens. If the
// input is invalid (too short or contain invalid tokens), the function returns an
// appropriate error. The input is converted to lower case.
func ParseInput(input string) ([]string, error) {
	acceptedTokens := []map[string]bool{
		map[string]bool{"cow": true, "bird": true, "snake": true, "leave": true},
		map[string]bool{"eat": true, "move": true, "speak": true}}

	inputTokens := strings.Fields(strings.ToLower(input))

	if len(inputTokens) > len(acceptedTokens) {
		fmt.Printf("Your command has too many words. I ignore '%s'.\n",
			strings.Join(inputTokens[len(acceptedTokens):], " "))
		inputTokens = inputTokens[:len(acceptedTokens)]
	}

	for level, token := range inputTokens {
		if acceptedTokens[level][token] == false {
			return inputTokens,
				errors.New("What do you mean by " + token + "?")
		}
	}

	switch {
	case len(inputTokens) == 0:
		return inputTokens,
			errors.New("You've given no command... If you want to leave, just type 'leave'.")
	case len(inputTokens) == 1 && inputTokens[0] != "leave":
		return inputTokens,
			errors.New("Which attribute of " + inputTokens[0] + " are you interested in?")
	}

	return inputTokens, nil
}
