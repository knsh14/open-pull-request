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

type RemoteInfo struct {
	Domain string
	Owner  string
	Repo   string
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
func (r *Repository) RemoteInfo() (*RemoteInfo, error) {
	remote, err := r.repo.Remote("origin")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get remote origin")
	}
	if len(remote.Config().URLs) != 1 {
		return nil, errors.New("origin url is not only one")
	}
	url := remote.Config().URLs[0]
	if strings.HasPrefix(url, "git") {
		r, err := getRemoteInfoGitProtocol(url)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get remote information")
		}
		return r, nil
	} else if strings.HasPrefix(url, "https") {
		r, err := getRemoteInfoHTTPProtocol(url)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get remote information")
		}
		return r, nil
	}
	return nil, errors.New("invalid protocol")
}

func getRemoteInfoGitProtocol(url string) (*RemoteInfo, error) {
	remote := &RemoteInfo{}
	u := strings.Split(url, "@")
	if len(u) != 2 {
		return nil, errors.New("failed to divide with @")
	}
	v := strings.SplitN(u[1], ":", 2)
	if len(v) != 2 {
		return nil, errors.New("failed to divide into domain and repository path")
	}
	if v[0] == "" || v[1] == "" {
		return nil, errors.New("failed to divide into domain and repository path")
	}
	remote.Domain = v[0]
	path := v[1]
	v = strings.SplitN(path, "/", 2)
	if len(v) != 2 {
		return nil, errors.New("failed to divide into owner and repository")
	}
	if v[0] == "" || v[1] == "" {
		return nil, errors.New("failed to divide into owner and repository")
	}
	remote.Owner, remote.Repo = v[0], strings.TrimSuffix(v[1], ".git")
	return remote, nil
}

func getRemoteInfoHTTPProtocol(url string) (*RemoteInfo, error) {
	u := strings.Split(url, "/")
	if len(u) != 5 {
		return nil, errors.New("failed to divide into parts")
	}
	for _, v := range u[2:] {
		if v == "" {
			return nil, errors.New("failed to divide into parts")
		}
	}
	return &RemoteInfo{Domain: u[2], Owner: u[3], Repo: strings.TrimSuffix(u[4], ".git")}, nil
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
