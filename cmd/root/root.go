package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	verbose bool
)

func PrintHelp() {
	fmt.Println("HELP")
}

var rootCommand = &cobra.Command{
	Use:   "github-client",
	Short: "github-client",
	Long:  "github-client",
	Run: func(cmd *cobra.Command, args []string) {
		PrintHelp()
	},
}

func Init() {
	rootCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

func AddSubCommand(subcommand *cobra.Command) {
	rootCommand.AddCommand(subcommand)
}

func Execute() error {
	return rootCommand.Execute()
}
