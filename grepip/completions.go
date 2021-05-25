package main

import (
	"github.com/ipinfo/cli/lib"
	"github.com/ipinfo/cli/lib/complete"
)

var completions = lib.CompletionsGrepIP

func handleCompletions() {
	completions.Complete(progBase)
}
