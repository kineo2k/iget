package collections

import (
	"fmt"
	"strings"
)

type nothing struct{}

var marking = nothing{}

type Set struct {
	m map[string]nothing
}

func NewSet() *Set {
	return &Set{make(map[string]nothing)}
}

func (s *Set) Add(value string) {
	s.m[value] = marking
}

func (s *Set) Remove(value string) {
	delete(s.m, value)
}

func (s *Set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *Set) Len() int {
	return len(s.m)
}

func (s *Set) Entries() []string {
	entries := make([]string, 0, len(s.m))
	for k := range s.m {
		entries = append(entries, k)
	}

	return entries
}

func (s *Set) String() string {
	return fmt.Sprintf("[%s]", strings.Join(s.Entries(), ", "))
}
