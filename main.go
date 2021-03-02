package main

import (
	"github.com/strongishllama/cmfive-cli/pkg/cmd"
	"github.com/strongishllama/cmfive-cli/pkg/cmfive"
)

func main() {
	cmfive.TemplatesDir = "cmfive/"
	cmd.Execute()
}
