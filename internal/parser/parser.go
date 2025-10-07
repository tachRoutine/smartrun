package parser

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tachRoutine/smartrun/pkg/types"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

// ExtractJson extracts JSON objects from a string, handling nested braces properly
func (p *Parser) ExtractJson(input string) []string {
	var jsonStrings []string

	for i := 0; i < len(input); i++ {
		if input[i] == '{' {
			// Found start of JSON object, now find matching closing brace
			braceCount := 1
			start := i
			j := i + 1

			for j < len(input) && braceCount > 0 {
				switch input[j] {
				case '{':
					braceCount++
				case '}':
					braceCount--
				}
				j++
			}

			if braceCount == 0 {
				// Found matching closing brace
				jsonString := input[start:j]
				jsonStrings = append(jsonStrings, jsonString)
				i = j - 1 // Continue from after this JSON object
			}
		}
	}

	return jsonStrings
}

// ParseExecTags extracts content between <exec> and </exec> tags
func (p *Parser) ParseExecTags(input string) []string {
	var commands []string
	startTag := "<exec>"
	endTag := "</exec>"

	for {
		startIdx := strings.Index(input, startTag)
		if startIdx == -1 {
			break
		}

		endIdx := strings.Index(input[startIdx+len(startTag):], endTag)
		if endIdx == -1 {
			break
		}

		// Adjust endIdx to be relative to original input
		endIdx += startIdx + len(startTag)

		command := input[startIdx+len(startTag) : endIdx]
		commands = append(commands, command)

		// Continue searching from after the end tag
		input = input[endIdx+len(endTag):]
	}

	return commands
}

func findIndex(s string, substr string, start int) int {
	index := strings.Index(s[start:], substr)
	if index == -1 {
		return -1
	}
	return start + index
}

// ParseJson returns types.LLMResponse from a json bytedata
func (p *Parser) ParseJson(jsonData []byte) (*types.LLMResponse, error) {
	var response types.LLMResponse
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return &response, nil
}
