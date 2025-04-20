package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var DeleteFileDefinition = ToolDefinition{
	Name:        "delete_file",
	Description: "Delete a file at the given path. This operation cannot be undone, so use with caution.",
	InputSchema: DeleteFileInputSchema,
	Function:    DeleteFile,
}

type DeleteFileInput struct {
	Path string `json:"path" jsonschema_description:"The relative path of the file to delete."`
}

var DeleteFileInputSchema = GenerateSchema[DeleteFileInput]()

func DeleteFile(input json.RawMessage) (string, error) {
	deleteFileInput := DeleteFileInput{}
	err := json.Unmarshal(input, &deleteFileInput)
	if err != nil {
		return "", err
	}

	if deleteFileInput.Path == "" {
		return "", fmt.Errorf("path must be provided")
	}

	// Check if file exists
	fileInfo, err := os.Stat(deleteFileInput.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file does not exist: %s", deleteFileInput.Path)
		}
		return "", err
	}

	// Don't allow deletion of directories for safety
	if fileInfo.IsDir() {
		return "", fmt.Errorf("cannot delete directories, '%s' is a directory", deleteFileInput.Path)
	}

	// Delete the file
	err = os.Remove(deleteFileInput.Path)
	if err != nil {
		return "", fmt.Errorf("failed to delete file: %w", err)
	}

	return fmt.Sprintf("Successfully deleted file: %s", deleteFileInput.Path), nil
}
