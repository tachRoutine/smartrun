package main

import (
	"fmt"
)

// The idea ihave
// recieve res from an llm then run/execute commands
// so i what i think is i should wrap the block that need to be executed in a custom tag e.g <exec>...</exec>
// then parse the response to find the tags and execute the commands inside them
// then return the output of the command to the user

func main() {
	fmt.Println("Hello, World!")
}
