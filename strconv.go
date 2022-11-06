package genenv

import (
	"strings"
	"sync"
	"unicode"
)

type StrConv struct {
	mu sync.Mutex
	B  *strings.Builder
	l  int
}

func NewStrConv() *StrConv {
	return &StrConv{
		B: new(strings.Builder),
	}
}

// ToPascal returns a pascal-cased copy of str.
// Any space, hyphen, underscore, comma or dot is removed and the next
// character is capitalized. Any other character is lowercase.
//
//	sc := NewStrConv()
//	s := sc.ToPascal("FOO_-._BAR")
//	// FooBar
func (s *StrConv) ToPascal(str string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.reset(str)

	u := true
	for _, r := range str {
		switch r {
		case ' ', '-', '_', ',', '.':
			u = true
			continue
		default:
			if !u {
				r = unicode.ToLower(r)
				break // switch
			}
			r = unicode.ToUpper(r)
			u = false
		}
		s.B.WriteRune(r)
	}

	return s.B.String()
}

func (s *StrConv) ToLower(str string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.reset(str)

	for _, r := range str {
		s.B.WriteRune(unicode.ToLower(r))
	}
	return s.B.String()
}

func (s *StrConv) reset(str string) {
	s.B.Reset()
	if l := len(str) - s.l; l > 0 {
		s.B.Grow(l)
		s.l += l
	}
}
