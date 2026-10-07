package custom_error

import (
	"fmt"
)

type FileCreationError struct {
	Message string
}

func (e *FileCreationError) Error() string {
	return fmt.Sprintf("Error during file creation: %s", e.Message)
}

type FileReadError struct {
	Message string
}

func (e *FileReadError) Error() string {
	return fmt.Sprintf("Error during file read: %s", e.Message)
}

type FileWriteError struct {
	Message string
}

func (e *FileWriteError) Error() string {
	return fmt.Sprintf("Error during file writing: %s", e.Message)
}
