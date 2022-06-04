package util

import (
	"context"
	"os"

	"github.com/richardsnider/github-cli/pkg/http"
	log "github.com/sirupsen/logrus"
)

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	http.InitClient(context.Background(), os.Getenv("GITHUB_TOKEN"))

	if os.Getenv("GITHUB_API_BASE_URL") != "" {
		githubApiUrl = os.Getenv("GITHUB_API_BASE_URL")
	}
}
