package gitee

import "k8s.io/apimachinery/pkg/util/sets"

func (ph *PullRequestHook) GetNumber() int32 {
	if ph == nil {
		return 0
	}

	return ph.Number
}

func (ph *PullRequestHook) GetState() string {
	if ph == nil {
		return ""
	}

	return ph.State
}

func (ph *PullRequestHook) GetHtmlURL() string {
	if ph == nil {
		return ""
	}

	return ph.HtmlUrl
}

func (ph *PullRequestHook) GetDiffUrl() string {
	if ph == nil {
		return ""
	}

	return ph.DiffUrl
}

func (ph *PullRequestHook) GetPatchUrl() string {
	if ph == nil {
		return ""
	}

	return ph.PatchUrl
}

func (ph *PullRequestHook) GetTitle() string {
	if ph == nil {
		return ""
	}

	return ph.Title
}

func (ph *PullRequestHook) GetBody() string {
	if ph == nil {
		return ""
	}

	return ph.Body
}

func (ph *PullRequestHook) GetStaleLabels() []LabelHook {
	if ph == nil {
		return nil
	}

	return ph.StaleLabels
}

func (ph *PullRequestHook) GetLabels() []LabelHook {
	if ph == nil {
		return nil
	}

	return ph.Labels
}

func (ph *PullRequestHook) GetCreatedAt() string {
	if ph == nil {
		return ""
	}

	return ph.CreatedAt
}

func (ph *PullRequestHook) GetUpdatedAt() string {
	if ph == nil {
		return ""
	}

	return ph.UpdatedAt
}

func (ph *PullRequestHook) GetClosedAt() string {
	if ph == nil {
		return ""
	}

	return ph.ClosedAt
}

func (ph *PullRequestHook) GetMergedAt() string {
	if ph == nil {
		return ""
	}

	return ph.MergedAt
}

func (ph *PullRequestHook) GetMergeCommitSha() string {
	if ph == nil {
		return ""
	}

	return ph.MergeCommitSha
}

func (ph *PullRequestHook) GetMergeReferenceName() string {
	if ph == nil {
		return ""
	}

	return ph.MergeReferenceName
}

func (ph *PullRequestHook) GetNeedTest() bool {
	if ph == nil {
		return false
	}

	return ph.NeedTest
}

func (ph *PullRequestHook) GetNeedReview() bool {
	if ph == nil {
		return false
	}

	return ph.NeedReview
}

func (ph *PullRequestHook) GetMergeable() bool {
	if ph == nil {
		return false
	}

	return ph.Mergeable
}

func (ph *PullRequestHook) GetUser() *UserHook {
	if ph == nil {
		return nil
	}

	return ph.User
}

func (ph *PullRequestHook) GetAssignee() *UserHook {
	if ph == nil {
		return nil
	}

	return ph.Assignee
}

func (ph *PullRequestHook) GetAssignees() []UserHook {
	if ph == nil {
		return nil
	}

	return ph.Assignees
}

func (ph *PullRequestHook) GetTester() []UserHook {
	if ph == nil {
		return nil
	}

	return ph.Tester
}

func (ph *PullRequestHook) GetTesters() []UserHook {
	if ph == nil {
		return nil
	}

	return ph.Testers
}

func (ph *PullRequestHook) GetMilestone() *MilestoneHook {
	if ph == nil {
		return nil
	}

	return ph.Milestone
}

func (ph *PullRequestHook) GetHead() *BranchHook {
	if ph == nil {
		return nil
	}

	return ph.Head
}

func (ph *PullRequestHook) GetBase() *BranchHook {
	if ph == nil {
		return nil
	}

	return ph.Base
}

func (ph *PullRequestHook) GetUpdatedBy() *UserHook {
	if ph == nil {
		return nil
	}

	return ph.UpdatedBy
}

func (ph *PullRequestHook) GetMerged() bool {
	if ph == nil {
		return false
	}

	return ph.Merged
}

func (ph *PullRequestHook) GetMergeStatus() string {
	if ph == nil {
		return ""
	}

	return ph.MergeStatus
}

func (ph *PullRequestHook) GetComments() int32 {
	if ph != nil {
		return 0
	}

	return ph.Comments
}

func (ph *PullRequestHook) GetCommits() int32 {
	if ph != nil {
		return 0
	}

	return ph.Commits
}

func (ph *PullRequestHook) GetAdditions() int32 {
	if ph != nil {
		return 0
	}

	return ph.Additions
}

func (ph *PullRequestHook) GetDeletions() int32 {
	if ph != nil {
		return 0
	}

	return ph.Deletions
}

func (ph *PullRequestHook) GetChangedFiles() int32 {
	if ph != nil {
		return 0
	}

	return ph.ChangedFiles
}

func (ph *PullRequestHook) GetIssues() []Issue {
	if ph == nil {
		return nil
	}

	return ph.Issues
}

func (ph *PullRequestHook) GetStaleIssues() []Issue {
	if ph == nil {
		return nil
	}

	return ph.StaleIssues
}

func (ph *PullRequestHook) LabelsToSet() sets.String {
	res := sets.NewString()

	for _, v := range ph.GetLabels() {
		res.Insert(v.GetName())
	}

	if res.Has("") {
		res.Delete("")
	}

	return res
}
