package util

var (
	semanticVersionLinkerFlag  string
	buildCommitLinkerFlag      string
	buildDateVersionLinkerFlag string

	githubApiUrl = "https://api.github.com"
)

type VersionData struct {
	Semantic    string
	BuildCommit string
	BuildDate   string
}

func GetVersionData() VersionData {
	return VersionData{
		Semantic:    semanticVersionLinkerFlag,
		BuildCommit: buildCommitLinkerFlag,
		BuildDate:   buildDateVersionLinkerFlag,
	}
}

func GetGithubApiUrl() string {
	return githubApiUrl
}
