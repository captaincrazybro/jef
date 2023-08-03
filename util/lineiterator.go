package util

import lu "github.com/captaincrazybro/literalutil"

// LineIterator a struct to iterator through an array of string lines
type LineIterator struct {
	lines []lu.String
	i     int
}

// New initializes the current struct with a list of string lines
func (lI *LineIterator) New(newLines []lu.String) {
	lI.lines = newLines
	lI.i = -1
}

// Next goes to the next index if possible
// If it is not possible, it will return false
func (lI *LineIterator) Next() bool {
	// Checks to see if the index has exceeded the array length
	if lI.i >= len(lI.lines)+1 {
		return false
	} else {
		// Advances the index of the iterator
		lI.i++
		return true
	}
}

// Current gets the current value of the iterator
func (lI *LineIterator) Current() lu.String {
	if lI.i >= len(lI.lines) || lI.i < 0 {
		return ""
	} else {
		return TrimWhitespaces(lI.lines[lI.i])
	}
}

// Index gets the current index of the iterator
func (lI *LineIterator) Index() int {
	return lI.i
}

// Lines gets the array of lines
func (lI *LineIterator) Lines() []lu.String {
	return lI.lines
}
