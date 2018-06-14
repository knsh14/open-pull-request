package github

import "testing"

func TestGetPullRequestURL(t *testing.T) {
	tests := []struct {
		input         [4]string
		expected      []string
		expectedError bool
	}{
		{
			[...]string{"github.com", "knsh14", "golang-testdata", "develop"},
			[]string{"https://github.com/knsh14/golang-testdata/pull/1"},
			false,
		},
	}

	for _, tt := range tests {
		urls, err := GetPullRequestURL(tt.input[0], tt.input[1], tt.input[2], tt.input[3])
		if tt.expectedError && err == nil {
			t.Error("expected to cause error but not")
		} else if tt.expectedError {
			if len(urls) != len(tt.expected) {
				t.Errorf("unexpected urls. expected=%d, got=%d", len(tt.expected), len(urls))

			}
			for i := range urls {
				if urls[i] != tt.expected[i] {
					t.Errorf("unexpected url of pull req. expected=%s, got=%s", tt.expected[i], urls[i])
				}
			}
		}
	}
}
