package parser

import (
	"reflect"
	"testing"

	"github.com/tachRoutine/smartrun/pkg/types"
)

// Package parser tests
func TestNewParser(t *testing.T) {
	parser := NewParser()
	if parser == nil {
		t.Error("NewParser() should return a non-nil parser")
	}
}

func TestExtractJson(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "single json object",
			input:    `{"key": "value"}`,
			expected: []string{`{"key": "value"}`},
		},
		{
			name:     "multiple json objects",
			input:    `{"first": "obj"} some text {"second": "obj"}`,
			expected: []string{`{"first": "obj"}`, `{"second": "obj"}`},
		},
		{
			name:     "nested json",
			input:    `{"outer": {"inner": "value"}}`,
			expected: []string{`{"outer": {"inner": "value"}}`},
		},
		{
			name:     "json with text around",
			input:    `Here is some text {"json": "data"} and more text`,
			expected: []string{`{"json": "data"}`},
		},
		{
			name:     "no json",
			input:    `This is just plain text without any json`,
			expected: []string{},
		},
		{
			name:     "empty string",
			input:    ``,
			expected: []string{},
		},
		{
			name:     "malformed json - missing closing brace",
			input:    `{"incomplete": "json"`,
			expected: []string{},
		},
		{
			name:     "json array",
			input:    `[{"item": 1}, {"item": 2}]`,
			expected: []string{`{"item": 1}`, `{"item": 2}`},
		},
	}

	parser := NewParser()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.ExtractJson(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ExtractJson() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestParseExecTags(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "single exec tag",
			input:    `<exec>ls -la</exec>`,
			expected: []string{"ls -la"},
		},
		{
			name:     "multiple exec tags",
			input:    `<exec>pwd</exec> some text <exec>whoami</exec>`,
			expected: []string{"pwd", "whoami"},
		},
		{
			name:     "exec tag with surrounding text",
			input:    `Here's a command: <exec>echo "hello"</exec> that prints hello`,
			expected: []string{`echo "hello"`},
		},
		{
			name:     "no exec tags",
			input:    `This text has no execution tags`,
			expected: []string{},
		},
		{
			name:     "empty exec tag",
			input:    `<exec></exec>`,
			expected: []string{""},
		},
		{
			name:     "malformed exec tag - missing closing",
			input:    `<exec>incomplete command`,
			expected: []string{},
		},
		{
			name:     "malformed exec tag - missing opening",
			input:    `incomplete command</exec>`,
			expected: []string{},
		},
		{
			name:     "nested-like tags",
			input:    `<exec>echo "<exec>inner</exec>"</exec>`,
			expected: []string{`echo "<exec>inner</exec>"`},
		},
		{
			name:     "multiline command",
			input:    "<exec>echo 'line1'\necho 'line2'</exec>",
			expected: []string{"echo 'line1'\necho 'line2'"},
		},
	}

	parser := NewParser()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.ParseExecTags(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseExecTags() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		substr   string
		expected int
	}{
		{
			name:     "substring at beginning",
			s:        "hello world",
			substr:   "hello",
			expected: 0,
		},
		{
			name:     "substring in middle",
			s:        "hello world",
			substr:   "lo wo",
			expected: 3,
		},
		{
			name:     "substring at end",
			s:        "hello world",
			substr:   "world",
			expected: 6,
		},
		{
			name:     "substring not found",
			s:        "hello world",
			substr:   "xyz",
			expected: -1,
		},
		{
			name:     "empty substring",
			s:        "hello world",
			substr:   "",
			expected: 0,
		},
		{
			name:     "empty string",
			s:        "",
			substr:   "test",
			expected: -1,
		},
		{
			name:     "substring longer than string",
			s:        "hi",
			substr:   "hello",
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findIndex(tt.s, tt.substr)
			if result != tt.expected {
				t.Errorf("findIndex(%q, %q) = %d, want %d", tt.s, tt.substr, result, tt.expected)
			}
		})
	}
}

func TestParseJson(t *testing.T) {
	tests := []struct {
		name        string
		jsonData    string
		expected    *types.LLMResponse
		expectError bool
	}{
		{
			name: "valid json response",
			jsonData: `{
                "instructions": "Execute these commands",
                "commands": [
                    {
                        "command": "ls -la",
                        "description": "List files",
                        "platform": "all",
                        "dangerous": false
                    }
                ]
            }`,
			expected: &types.LLMResponse{
				Instructions: "Execute these commands",
				Commands: []types.Command{
					{
						Command:     "ls -la",
						Description: "List files",
						Platform:    "all",
						Dangerous:   false,
					},
				},
			},
			expectError: false,
		},
		{
			name: "multiple commands",
			jsonData: `{
                "instructions": "Multiple commands to execute",
                "commands": [
                    {
                        "command": "pwd",
                        "description": "Show directory",
                        "platform": "linux",
                        "dangerous": false
                    },
                    {
                        "command": "rm -rf /",
                        "description": "Dangerous command",
                        "platform": "linux",
                        "dangerous": true
                    }
                ]
            }`,
			expected: &types.LLMResponse{
				Instructions: "Multiple commands to execute",
				Commands: []types.Command{
					{
						Command:     "pwd",
						Description: "Show directory",
						Platform:    "linux",
						Dangerous:   false,
					},
					{
						Command:     "rm -rf /",
						Description: "Dangerous command",
						Platform:    "linux",
						Dangerous:   true,
					},
				},
			},
			expectError: false,
		},
		{
			name: "empty commands array",
			jsonData: `{
                "instructions": "No commands",
                "commands": []
            }`,
			expected: &types.LLMResponse{
				Instructions: "No commands",
				Commands:     []types.Command{},
			},
			expectError: false,
		},
		{
			name:        "invalid json",
			jsonData:    `{"invalid": json}`,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "empty json",
			jsonData:    ``,
			expected:    nil,
			expectError: true,
		},
		{
			name: "missing required fields",
			jsonData: `{
                "instructions": "Test"
            }`,
			expected: &types.LLMResponse{
				Instructions: "Test",
				Commands:     nil,
			},
			expectError: false,
		},
	}

	parser := NewParser()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.ParseJson([]byte(tt.jsonData))

			if tt.expectError {
				if err == nil {
					t.Error("ParseJson() expected error, but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("ParseJson() unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseJson() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkExtractJson(b *testing.B) {
	parser := NewParser()
	input := `{"first": "value"} some text {"second": "value"} more text {"third": "value"}`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parser.ExtractJson(input)
	}
}

func BenchmarkParseExecTags(b *testing.B) {
	parser := NewParser()
	input := `<exec>command1</exec> text <exec>command2</exec> more text <exec>command3</exec>`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parser.ParseExecTags(input)
	}
}

func BenchmarkParseJson(b *testing.B) {
	parser := NewParser()
	jsonData := []byte(`{
        "instructions": "Test instructions",
        "commands": [
            {
                "command": "ls -la",
                "description": "List files",
                "platform": "all",
                "dangerous": false
            }
        ]
    }`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parser.ParseJson(jsonData)
	}
}

// Test with actual sample file
func TestParseJsonWithSampleFile(t *testing.T) {
	// This would be the content from your sample file
	sampleJsonData := `{
        "instructions": "Sure, I can help you with that! Let me check what files are in the current directory, show you the current working directory, and display system information.",
        "commands": [
            {
                "command": "ls -la",
                "description": "List all files in the current directory with detailed information",
                "platform": "all",
                "dangerous": false
            },
            {
                "command": "pwd",
                "description": "Show the current working directory path",
                "platform": "all", 
                "dangerous": false
            },
            {
                "command": "uname -a",
                "description": "Display detailed system information including kernel and architecture",
                "platform": "all",
                "dangerous": false
            }
        ]
    }`

	parser := NewParser()
	result, err := parser.ParseJson([]byte(sampleJsonData))

	if err != nil {
		t.Fatalf("ParseJson() failed with sample data: %v", err)
	}

	if result.Instructions == "" {
		t.Error("Instructions should not be empty")
	}

	if len(result.Commands) != 3 {
		t.Errorf("Expected 3 commands, got %d", len(result.Commands))
	}

	// Verify first command
	firstCmd := result.Commands[0]
	if firstCmd.Command != "ls -la" {
		t.Errorf("Expected 'ls -la', got '%s'", firstCmd.Command)
	}

	if firstCmd.Dangerous != false {
		t.Errorf("Expected dangerous=false, got %v", firstCmd.Dangerous)
	}

	if firstCmd.Platform != "all" {
		t.Errorf("Expected platform='all', got '%s'", firstCmd.Platform)
	}
}
