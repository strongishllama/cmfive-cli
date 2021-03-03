package main

import (
	"github.com/strongishllama/cmfive-cli/pkg/cli"
	"github.com/strongishllama/cmfive-cli/pkg/gen"
)

func main() {
	gen.TemplatesDir = "gen/"
	cli.Execute()
}
