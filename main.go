package main

import (
	"github.com/richardsnider/github-cli/cmd"
	"github.com/richardsnider/github-cli/pkg/util"
)

func main() {
	util.Init()
	cmd.Execute()
}
