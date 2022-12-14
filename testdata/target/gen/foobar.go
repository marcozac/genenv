// Code generated by genenv. DO NOT EDIT.

package env

import (
	"fmt"
	"os"
)

// FooBar returns the value of the environment variable "FOO_BAR".
func FooBar() (string, error) {
	v := os.Getenv("FOO_BAR")
	
	denied := []string{
		"bar",
	}
	for _, d := range denied {
		if d == v {
			return "", fmt.Errorf("%s not allowed", v)
		}
	}
	
	allowed := []string{
		"foo",
	}
	for _, s := range allowed {
		if s == v {
			return v, nil
		}
	}
	return "", fmt.Errorf("%s not allowed", v)
}
