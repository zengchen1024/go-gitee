package test

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"testing"

	"github.com/opensourceways/go-gitee/gitee"
)

const (
	testOrg  = "cve-manage-test"
	testRepo = "config"
)

var client *gitee.APIClient

func init() {
	token := "your gitee token"

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	conf := gitee.NewConfiguration()
	conf.HTTPClient = oauth2.NewClient(context.Background(), ts)

	client = gitee.NewAPIClient(conf)
}

func TestPutV5ReposOwnerRepoPullsNumberLabels(t *testing.T) {
	op := gitee.PullRequestLabelPostParam{
		Body: []string{"feature", "go-gitee-test"},
	}
	labels, _, err := client.PullRequestsApi.PutV5ReposOwnerRepoPullsNumberLabels(context.Background(), testOrg, testRepo, 1, op)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(labels)
	}
}

func TestPutV5ReposOwnerRepoIssuesNumberLabels(t *testing.T) {
	op := gitee.PullRequestLabelPostParam{
		Body: []string{"feature", "go-gitee-test"},
	}
	labels, _, err := client.LabelsApi.PutV5ReposOwnerRepoIssuesNumberLabels(context.Background(), testOrg, testRepo, "I3A2AO", op)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(labels)
	}

}

func TestPatchV5ReposOwnerIssuesNumber(t *testing.T) {
	body := gitee.IssueUpdateParam{
		Repo:          testRepo,
		Collaborators: "zhangjianjun_code",
	}
	_, _, err := client.IssuesApi.PatchV5ReposOwnerIssuesNumber(context.Background(), testOrg, "I3IZ20", body)
	if err != nil {
		t.Error(err)
	}
}
