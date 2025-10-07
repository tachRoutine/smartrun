# SmartRun Architecture

## Overview
SmartRun is designed to parse LLM responses and execute commands safely.

## Components

### Parser
- Extracts `<exec>` tags from text
- Validates command syntax
- Returns structured command data

### Executor  
- Executes parsed commands
- Implements safety checks
- Returns command output

### LLM Interface
- Connects to various LLM providers
- Handles API communication
- Manages response parsing

### Configuration
- Manages API keys
- Handles user preferences  
- Security settings