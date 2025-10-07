package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// The idea ihave
// recieve res from an llm then run/execute commands 
// so i what i think is i should wrap the block that need to be executed in a custom tag e.g <exec>...</exec>
// then parse the response to find the tags and execute the commands inside them
// then return the output of the command to the user


func main() {
	url := "http://localhost:8080/"

	payload := map[string]any{
		"input": "Hello, how can I use Docker with Go?",
	}
	payloadBytes, _ := json.Marshal(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))
}
