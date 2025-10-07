package parser

import (
	"encoding/json"
	"fmt"

	"github.com/tachRoutine/smartrun/pkg/types"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseExecTags(input string) []string {
	var commands []string
	startTag := "<exec>"
	endTag := "</exec>"

	for {
		startIdx := findIndex(input, startTag)
		if startIdx == -1 {
			break
		}
		endIdx := findIndex(input[startIdx:], endTag)
		if endIdx == -1 {
			break
		}
		endIdx += startIdx

		command := input[startIdx+len(startTag) : endIdx]
		commands = append(commands, command)

		input = input[endIdx+len(endTag):]
	}

	return commands
}

func findIndex(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func (p *Parser) ParseJson(jsonData []byte) (*types.LLMResponse, error) {
	var response types.LLMResponse
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return &response, nil
}

// Parser handles parsing of <exec>...</exec> tags from LLM responses


