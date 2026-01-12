package main

import (
	"github.com/iihsannlearn/corscan/internal/options"
	"github.com/iihsannlearn/corscan/internal/runner"
)

func main() {
	opts := options.ParseOptions()
	runner.Run(opts)
}
