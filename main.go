package main

import (
	"github.com/strongishllama/cmfive-cli/pkg/cli"
	"github.com/strongishllama/cmfive-cli/pkg/cmfive"
)

func main() {
	cmfive.TemplatesDir = "cmfive/"
	cli.Execute()
}
