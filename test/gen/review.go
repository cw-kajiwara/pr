package gen

import (
	"encoding/json"
	"sync"

	"github.com/google/go-github/v28/github"
)

// See: https://developer.github.com/v3/pulls/reviews/#list-reviews-on-a-pull-request
var reviewText = `{"id":80,"node_id":"MDE3OlB1bGxSZXF1ZXN0UmV2aWV3ODA=","user":{"login":"octocat","id":1,"node_id":"MDQ6VXNlcjE=","avatar_url":"https://github.com/images/error/octocat_happy.gif","gravatar_id":"","url":"https://api.github.com/users/octocat","html_url":"https://github.com/octocat","followers_url":"https://api.github.com/users/octocat/followers","following_url":"https://api.github.com/users/octocat/following{/other_user}","gists_url":"https://api.github.com/users/octocat/gists{/gist_id}","starred_url":"https://api.github.com/users/octocat/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/octocat/subscriptions","organizations_url":"https://api.github.com/users/octocat/orgs","repos_url":"https://api.github.com/users/octocat/repos","events_url":"https://api.github.com/users/octocat/events{/privacy}","received_events_url":"https://api.github.com/users/octocat/received_events","type":"User","site_admin":false},"body":"Here is the body for the review.","commit_id":"ecdd80bb57125d7ba9641ffaa4d7d2c19d3f3091","state":"APPROVED","html_url":"https://github.com/octocat/Hello-World/pull/12#pullrequestreview-80","pull_request_url":"https://api.github.com/repos/octocat/Hello-World/pulls/12","_links":{"html":{"href":"https://github.com/octocat/Hello-World/pull/12#pullrequestreview-80"},"pull_request":{"href":"https://api.github.com/repos/octocat/Hello-World/pulls/12"}}}`
var review *github.PullRequestReview
var reviewOnce sync.Once

func PullRequestReview() (*github.PullRequestReview, error) {
	var err error
	reviewOnce.Do(func() {
		err = json.Unmarshal([]byte(reviewText), &review)
	})
	if err != nil {
		return nil, err
	}
	r := *review
	return &r, nil
}

func PullRequestReviews(length int) ([]*github.PullRequestReview, error) {
	values := make([]*github.PullRequestReview, length)
	for i := 0; i < length; i++ {
		v, err := PullRequestReview()
		if err != nil {
			return nil, err
		}
		values[i] = v
	}
	return values, nil
}
