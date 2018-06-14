package repository

import "testing"

func TestGetRemoteFromGitProtocol(t *testing.T) {
	tests := []struct {
		input         string
		expected      *RemoteInfo
		expectedError bool
	}{
		{
			input: "git@github.com:knsh14/open-pull-request.git",
			expected: &RemoteInfo{
				Domain: "github.com",
				Owner:  "knsh14",
				Repo:   "open-pull-request",
			},
			expectedError: false,
		},
		{
			input: "git@:knsh14/open-pull-request.git",
			expected: &RemoteInfo{
				Domain: "github.com",
				Owner:  "knsh14",
				Repo:   "open-pull-request",
			},
			expectedError: true,
		},
		{
			input: "git@github.com:/open-pull-request.git",
			expected: &RemoteInfo{
				Domain: "github.com",
				Owner:  "knsh14",
				Repo:   "open-pull-request",
			},
			expectedError: true,
		},
		{
			input: "git@github.com:knsh14/",
			expected: &RemoteInfo{
				Domain: "github.com",
				Owner:  "knsh14",
				Repo:   "open-pull-request",
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		remote, err := getRemoteInfoGitProtocol(tt.input)
		if tt.expectedError && err == nil {
			t.Errorf("input %s expected to cause error but not", tt.input)
		} else if !tt.expectedError {
			if remote.Domain != tt.expected.Domain {
				t.Errorf("Domain was wrong. expected=%s, got=%s", tt.expected.Domain, remote.Domain)
			}
			if remote.Owner != tt.expected.Owner {
				t.Errorf("Owner was wrong. expected=%s, got=%s", tt.expected.Owner, remote.Owner)
			}
			if remote.Repo != tt.expected.Repo {
				t.Errorf("Repo was wrong. expected=%s, got=%s", tt.expected.Repo, remote.Repo)
			}
		}
	}
}

func TestGetRemoteFromHTTPProtocol(t *testing.T) {
	tests := []struct {
		input         string
		expected      *RemoteInfo
		expectedError bool
	}{
		{
			input: "https://github.com/knsh14/open-pull-request.git",
			expected: &RemoteInfo{
				Domain: "github.com",
				Owner:  "knsh14",
				Repo:   "open-pull-request",
			},
			expectedError: false,
		},
		{
			input: "https:///knsh14/open-pull-request.git",
			expected: &RemoteInfo{
				Domain: "github.com",
				Owner:  "knsh14",
				Repo:   "open-pull-request",
			},
			expectedError: true,
		},
		{
			input: "https://github.com//open-pull-request.git",
			expected: &RemoteInfo{
				Domain: "github.com",
				Owner:  "knsh14",
				Repo:   "open-pull-request",
			},
			expectedError: true,
		},
		{
			input: "https://github.com/knsh14/",
			expected: &RemoteInfo{
				Domain: "github.com",
				Owner:  "knsh14",
				Repo:   "open-pull-request",
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		remote, err := getRemoteInfoHTTPProtocol(tt.input)
		if tt.expectedError && err == nil {
			t.Errorf("input %s expected to cause error but not", tt.input)
		} else if !tt.expectedError {
			if remote.Domain != tt.expected.Domain {
				t.Errorf("Domain was wrong. expected=%s, got=%s", tt.expected.Domain, remote.Domain)
			}
			if remote.Owner != tt.expected.Owner {
				t.Errorf("Owner was wrong. expected=%s, got=%s", tt.expected.Owner, remote.Owner)
			}
			if remote.Repo != tt.expected.Repo {
				t.Errorf("Repo was wrong. expected=%s, got=%s", tt.expected.Repo, remote.Repo)
			}
		}
	}
}
