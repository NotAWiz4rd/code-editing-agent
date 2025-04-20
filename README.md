# Code Editing Agent

*Note: This README was created by me - the very same Claude-powered code editing agent that this repository implements. Meta, right?*

## Overview

This repository contains a Go implementation of an AI-powered code editing agent that uses the Anthropic Claude API. The agent can help you manage files in your project through a chat interface, providing a seamless way to perform file operations without leaving your terminal.

## Features

The agent supports the following file operations:

- **Read Files**: View the contents of files in your project
- **List Files**: See all files and directories in a specified path
- **Edit Files**: Make changes to text files by replacing text
- **Delete Files**: Remove files (with safeguards to prevent accidental deletions)

## How It Works

This tool creates a chat interface that connects to Claude via the Anthropic API. The agent uses function calling to execute file operations requested during the conversation. When you ask to perform a file operation, the agent:

1. Interprets your request
2. Calls the appropriate function
3. Returns the results to you in a conversational format

## Requirements

- Go 1.24 or higher
- Anthropic API key (set as an environment variable)

## Installation

1. Clone this repository:

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Set your Anthropic API key:
   ```
   export ANTHROPIC_API_KEY=your_api_key_here
   ```

## Usage

Run the agent:

```
go run .
```

This will start the chat interface where you can interact with the agent. For example:

- `List all files in the current directory`
- `Show me the content of main.go`
- `Create a new file called example.txt with "Hello World" in it`
- `Replace "Hello World" with "Hello, AI!" in example.txt`

## Architecture

The project consists of several components:

- `main.go`: Entry point that initializes the chat agent
- `ReadFileTool.go`: Logic for reading file contents
- `ListFilesTool.go`: Logic for listing files and directories
- `EditFileTool.go`: Logic for modifying or creating files
- `DeleteFileTool.go`: Logic for safely deleting files


---

*I find it a bit surreal to be documenting myself, but I hope this README helps you understand the capabilities of this code editing agent. If you have any questions, just ask me in the chat!*