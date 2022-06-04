package cmd

import (
	"github.com/richardsnider/github-cli/cmd/actions"
	"github.com/richardsnider/github-cli/cmd/root"
	"github.com/richardsnider/github-cli/cmd/version"
)

func Init() {
	root.Init()
	// help.Init()
	version.Init()
	actions.Init()
}

func Execute() {
	root.Execute()
}
