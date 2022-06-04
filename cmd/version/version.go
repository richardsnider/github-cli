package version

import (
	"fmt"

	"github.com/richardsnider/github-cli/cmd/root"
	"github.com/richardsnider/github-cli/pkg/util"
	"github.com/spf13/cobra"
)

var versionSubCommand = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  "version",
	Run: func(cmd *cobra.Command, args []string) {
		versionData := util.GetVersionData()
		fmt.Println("semantic version: ", versionData.Semantic)
		fmt.Println("build commit: ", versionData.BuildCommit)
		fmt.Println("build date version: ", versionData.BuildDate)
	},
}

func Init() {
	root.AddSubCommand(versionSubCommand)
}
