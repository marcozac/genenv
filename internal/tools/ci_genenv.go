//go:build ignore
// +build ignore

// Running after tests, checks the repository files status, reporting
// an error if there are any changes.
package main

import (
	"log"
	"os/exec"
)

func main() {
	o, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		log.Fatal(err)
	}
	if len(o) > 0 {
		log.Fatalf("genenv ci failed:\n %s", string(o))
	}
}
