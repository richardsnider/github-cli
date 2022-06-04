package actions

import (
	"github.com/richardsnider/github-cli/cmd/root"
	"github.com/richardsnider/github-cli/pkg/http"
	"github.com/richardsnider/github-cli/pkg/util"
	"github.com/spf13/cobra"
)

var (
	relativeHoursPast int
)

func identiyFailedWorkflowRuns(repoName string, relativeHoursPast int) (runIds []string, err error) {
	// https://docs.github.com/en/rest/actions/workflow-runs#list-workflow-runs-for-a-repository
	response, httpGetErr := http.GetURL(util.GetGithubApiUrl() + "/repos/" + repoName + "/actions/runs")
	if httpGetErr != nil {
		return []string{}, httpGetErr
	}

	return []string{
		response,
	}, nil
}

var workflowSubCommand = &cobra.Command{
	Use:   "workflow",
	Short: "workflow",
	Long:  "workflow",
	Run: func(cmd *cobra.Command, args []string) {
		identiyFailedWorkflowRuns(args[0], relativeHoursPast)
	},
}

func Init() {
	workflowSubCommand.Flags().IntVarP(&relativeHoursPast, "hours", "h", 12, "relative hours past")
	root.AddSubCommand(workflowSubCommand)
}
