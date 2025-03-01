package main

import (
	"io"
)

// FilteringPipe filters out numbers from the input
type FilteringPipe struct {
	writer io.Writer
}

// NewFilteringPipe creates a new FilteringPipe with the given writer
func NewFilteringPipe(w io.Writer) FilteringPipe {
	return FilteringPipe{writer: w}
}

// Write writes the bytes to the writer after filtering out any numbers
func (f *FilteringPipe) Write(p []byte) (int, error) {
	var filtered []byte
	for _, b := range p {
		if b < '0' || b > '9' {
			filtered = append(filtered, b)
		}
	}

	if _, err := f.writer.Write(filtered); err != nil {
		return 0, err
	}

	return len(p), nil
}
