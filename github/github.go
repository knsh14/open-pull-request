package github

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/knsh14/udemy/repository"
	"github.com/pkg/errors"
	gitconfig "github.com/tcnksm/go-gitconfig"
	"golang.org/x/oauth2"
)

// GetPullRequestURL returns urls which are given branch is head
func GetPullRequestURL(branch string) ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	r, err := repository.NewRepository(dir)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Repository struct")
	}
	domain, owner, repo, err := r.RemoteInfo()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get repository info")
	}

	token, err := gitconfig.GithubToken()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get github token")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	if domain != "github.com" {
		baseURL := fmt.Sprintf("https://%s/api/v3/", domain)
		client, err = github.NewEnterpriseClient(baseURL, baseURL, tc)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get github enterprise client")
		}
	}

	pulls, _, err := client.PullRequests.List(ctx, owner, repo, &github.PullRequestListOptions{
		State: "open",
		Head:  owner + ":" + branch,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get list of pull requests")
	}
	var urls []string
	for _, p := range pulls {
		urls = append(urls, p.GetHTMLURL())
	}
	return urls, nil
}
