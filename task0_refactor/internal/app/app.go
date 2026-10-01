// Package app coordinates the command-line application's domain packages.
package app

import (
	"errors"
	"fmt"
	"io"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/calculator"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/textanalyzer"
)

// Run executes the sample workflow and writes its output to out.
func Run(out io.Writer) error {
	sum := calculator.Add(2, 3)
	if _, err := fmt.Fprintln(out, "2 + 3 =", sum); err != nil {
		return fmt.Errorf("app: write sum: %w", err)
	}

	quotient, err := calculator.Divide(10, 0)
	if err != nil {
		if errors.Is(err, calculator.ErrDivisionByZero) {
			if _, writeErr := fmt.Fprintln(out, "division by zero was correctly detected:", err); writeErr != nil {
				return fmt.Errorf("app: write division result: %w", writeErr)
			}
		} else {
			return fmt.Errorf("app: divide: %w", err)
		}
	} else if _, err := fmt.Fprintln(out, "10 / 0 =", quotient); err != nil {
		return fmt.Errorf("app: write quotient: %w", err)
	}

	words, err := textanalyzer.WordCount("the quick brown fox")
	if err != nil {
		return fmt.Errorf("app: count words: %w", err)
	}
	if _, err := fmt.Fprintln(out, "word count:", words); err != nil {
		return fmt.Errorf("app: write word count: %w", err)
	}
	return nil
}
