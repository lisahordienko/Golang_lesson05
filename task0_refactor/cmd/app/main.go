// Command app wires the application and its process-level error handling.
package main

import (
	"fmt"
	"os"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/app"
)

func main() {
	if err := app.Run(os.Stdout); err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}
}
