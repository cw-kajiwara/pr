package gen

import (
	"encoding/json"
	"sync"

	"github.com/google/go-github/v28/github"
)

// See: https://developer.github.com/v3/search/#search-issues-and-pull-requests
var issueSearchResultText = `{"total_count":280,"incomplete_results":false,"items":[{"url":"https://api.github.com/repos/batterseapower/pinyin-toolkit/issues/1","repository_url":"https://api.github.com/repos/batterseapower/pinyin-toolkit","labels_url":"https://api.github.com/repos/batterseapower/pinyin-toolkit/issues/1/labels{/name}","comments_url":"https://api.github.com/repos/batterseapower/pinyin-toolkit/issues/1/comments","events_url":"https://api.github.com/repos/batterseapower/pinyin-toolkit/issues/1/events","html_url":"https://github.com/batterseapower/pinyin-toolkit/issues/1","id":35802,"node_id":"MDU6SXNzdWUzNTgwMg==","number":1,"title":"Line Number Indexes Beyond 20 Not Displayed","user":{"login":"Nick3C","id":90254,"node_id":"MDQ6VXNlcjkwMjU0","avatar_url":"https://secure.gravatar.com/avatar/934442aadfe3b2f4630510de416c5718?d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-user-420.png","gravatar_id":"","url":"https://api.github.com/users/Nick3C","html_url":"https://github.com/Nick3C","followers_url":"https://api.github.com/users/Nick3C/followers","following_url":"https://api.github.com/users/Nick3C/following{/other_user}","gists_url":"https://api.github.com/users/Nick3C/gists{/gist_id}","starred_url":"https://api.github.com/users/Nick3C/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/Nick3C/subscriptions","organizations_url":"https://api.github.com/users/Nick3C/orgs","repos_url":"https://api.github.com/users/Nick3C/repos","events_url":"https://api.github.com/users/Nick3C/events{/privacy}","received_events_url":"https://api.github.com/users/Nick3C/received_events","type":"User"},"labels":[{"id":4,"node_id":"MDU6TGFiZWw0","url":"https://api.github.com/repos/batterseapower/pinyin-toolkit/labels/bug","name":"bug","color":"ff0000"}],"state":"open","assignee":null,"milestone":null,"comments":15,"created_at":"2009-07-12T20:10:41Z","updated_at":"2009-07-19T09:23:43Z","closed_at":null,"pull_request":{"html_url":null,"diff_url":null,"patch_url":null},"body":"...","score":1.3859273}]}`
var issueSearchResult *github.IssuesSearchResult
var issueSearchResultOnce sync.Once

func IssuesSearchResult() (*github.IssuesSearchResult, error) {
	var err error
	issueSearchResultOnce.Do(func() {
		err = json.Unmarshal([]byte(issueSearchResultText), &issueSearchResult)
	})
	if err != nil {
		return nil, err
	}
	i := *issueSearchResult
	return &i, nil
}
