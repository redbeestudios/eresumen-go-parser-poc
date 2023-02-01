package utils

import (
	"strings"
)

type LineReader struct {
	line  string
	index int
}

func NewLineReader(line string) *LineReader {
	return &LineReader{
		line:  line,
		index: 0,
	}
}

func (r *LineReader) Read(characterCount int, charactersToSkip ...int) string {
	line := r.line[r.index : r.index+characterCount]
	r.index = r.index + characterCount
	if len(charactersToSkip) > 0 {
		r.Skip(charactersToSkip[0])
	}
	return strings.Trim(line, " ")
}

func (r *LineReader) Skip(characterCount int) {
	r.index = r.index + characterCount
}
