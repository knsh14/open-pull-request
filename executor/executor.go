package executor

import (
	"log"
	"os"
	"sync"

	"github.com/knsh14/udemy/github"
	"github.com/knsh14/udemy/repository"
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
	currentBranch, err := repo.GetCurrentBranch()
	if err != nil {
		log.Fatal(err)
	}
	urls, err := github.GetPullRequestURL(currentBranch)
	if err != nil {
		return errors.Wrapf(err, "failed to get pullrequest url for %s", currentBranch)
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
	branches, err := repo.GetAllBranches()
	if err != nil {
		return errors.Wrap(err, "failed to get all branches")
	}
	var wg sync.WaitGroup
	for _, branch := range branches {
		wg.Add(1)
		go func(b string) {
			urls, err := github.GetPullRequestURL(b)
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
