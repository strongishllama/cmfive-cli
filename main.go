package main

import (
	"github.com/strongishllama/cmfive-cli/cmd"
	"github.com/strongishllama/cmfive-cli/pkg/cmfive"
)

func main() {
	cmfive.TemplatesDir = "cmfive/"
	cmd.Execute()
}
