package repository

import (
	"strings"

	"github.com/pkg/errors"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Repository is wrapper for git repsitory
type Repository struct {
	repo *git.Repository
}

// NewRepository creates struct to get repository information
func NewRepository(path string) (*Repository, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open repository")
	}
	return &Repository{repo: r}, nil
}

// RemoteInfo returns information of origin
// TODO: work only git protocol only. need to work for http protocol
func (r *Repository) RemoteInfo() (string, string, string, error) {
	remote, err := r.repo.Remote("origin")
	if err != nil {
		return "", "", "", errors.Wrap(err, "failed to get remote origin")
	}
	if len(remote.Config().URLs) != 1 {
		return "", "", "", errors.New("origin url is not only one")
	}
	url := remote.Config().URLs[0]
	u := strings.Split(url, "@")[1]
	v := strings.SplitN(u, ":", 2)
	domain, path := v[0], v[1]
	v = strings.SplitN(path, "/", 2)
	org, repo := v[0], strings.TrimRight(v[1], ".git")
	return domain, org, repo, nil
}

// GetCurrentBranch returns branch name of head commit
func (r *Repository) GetCurrentBranch() (string, error) {

	head, err := r.repo.Head()
	if err != nil {
		return "", errors.Wrap(err, "failed to get head branch name")
	}
	return head.Name().Short(), nil
}

// GetAllBranches return all branch name in local
func (r *Repository) GetAllBranches() ([]string, error) {
	branches, err := r.repo.Branches()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all branched in local")
	}
	var brancheNames []string
	branches.ForEach(func(r *plumbing.Reference) error {
		brancheNames = append(brancheNames, r.Name().Short())
		return nil
	})
	return brancheNames, nil
}
