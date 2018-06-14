package executor

import (
	"log"
	"os"
	"sync"

	"github.com/knsh14/open-pull-request/github"
	"github.com/knsh14/open-pull-request/repository"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
)

// OpenCurrentBranch openes pull request page for current branch
func OpenCurrentBranch() error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	repo, err := repository.NewRepository(dir)
	if err != nil {
		return errors.Wrap(err, "failed to get Repository struct")
	}
	urls, err := github.GetPullRequestURL(repo.Remote.Domain, repo.Remote.Owner, repo.Remote.Repo, repo.Head)
	if err != nil {
		return errors.Wrapf(err, "failed to get pullrequest url for %s", repo.Head)
	}
	for _, u := range urls {
		err := browser.OpenURL(u)
		if err != nil {
			return errors.Wrapf(err, "failed to open %s in browser", u)
		}
	}
	return nil
}

// OpenAllBranches openes pull request page for all branches in local
func OpenAllBranches() error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	repo, err := repository.NewRepository(dir)
	if err != nil {
		return errors.Wrap(err, "failed to get Repository struct")
	}
	var wg sync.WaitGroup
	for _, branch := range repo.Branches {
		wg.Add(1)
		go func(b string) {
			urls, err := github.GetPullRequestURL(repo.Remote.Domain, repo.Remote.Owner, repo.Remote.Repo, b)
			if err != nil {
				log.Print(err)
				wg.Done()
			}
			for _, u := range urls {
				err := browser.OpenURL(u)
				if err != nil {
					log.Fatal(err)
				}
			}
			wg.Done()
		}(branch)
	}
	wg.Wait()
	return nil
}
