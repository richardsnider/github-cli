package help

import (
	"github.com/richardsnider/github-cli/cmd/root"
	"github.com/spf13/cobra"
)

var helpSubCommand = &cobra.Command{
	Use:   "help",
	Short: "help",
	Long:  "help",
	Run: func(cmd *cobra.Command, args []string) {
		root.PrintHelp()
	},
}

func Init() {
	root.AddSubCommand(helpSubCommand)
}
