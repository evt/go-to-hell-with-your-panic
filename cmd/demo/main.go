package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/evt/go-to-hell-with-your-panic/pkg"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	err := pkg.DoSomething()
	if err != nil {
		if errors.Is(err, pkg.ErrDoSomethingElse) {
			return fmt.Errorf("intercepted: %w", err)
		}
		return fmt.Errorf("doSomething failed: %w", err)
	}
	return nil
}
