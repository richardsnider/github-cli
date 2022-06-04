package http

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"golang.org/x/oauth2"
)

var httpClient *http.Client

func InitClient(ctx context.Context, oauth2AccessToken string) {
	// https://github.com/google/go-github#authentication
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: oauth2AccessToken},
	)
	httpClient = oauth2.NewClient(ctx, tokenSource)
}

func GetURL(URL string) (result string, err error) {
	response, httpGetError := http.Get(URL)
	if httpGetError != nil {
		return "", httpGetError
	}

	defer response.Body.Close()
	body, responseBodyError := ioutil.ReadAll(response.Body)
	if responseBodyError != nil {
		return "", responseBodyError
	}

	if response.StatusCode != 200 {
		return "", errors.New("Unexpected response (HTTP STATUS CODE " + strconv.Itoa(response.StatusCode) + ") " + string(body))
	}

	return string(body), nil
}
