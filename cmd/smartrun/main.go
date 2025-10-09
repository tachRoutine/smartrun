package main

import (
	"fmt"
	"os"

	"github.com/tachRoutine/smartrun/internal/executor"
	"github.com/tachRoutine/smartrun/internal/parser"
	"github.com/tachRoutine/smartrun/pkg/types"
)

// The idea ihave
// recieve res from an llm then run/execute commands
// so i what i think is i should wrap the block that need to be executed in a custom tag e.g <exec>...</exec>
// then parse the response to find the tags and execute the commands inside them
// then return the output of the command to the user

func main() {
	testResponse, err := os.ReadFile("./testdata/sample_llm_response.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	parser := parser.NewParser()
	execCommands := parser.ParseExecTags(string(testResponse))
	fmt.Println("Extracted commands between <exec> tags:")
	for _, cmd := range execCommands {
		fmt.Println(cmd)
	}

	jsonObjects := parser.ExtractJson(string(testResponse))
	
	for _, jsonObj := range jsonObjects {
		fmt.Println("\nExtracted JSON object:" , jsonObj)
		var parsed *types.LLMResponse
		parsed, err = parser.ParseJson([]byte(jsonObj))
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}
		fmt.Println("PARSED JSON", parsed)
		executor.NewExecutor().Execute(*parsed)
	}

	fmt.Println("==================================================")

}
