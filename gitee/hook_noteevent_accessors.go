package gitee

import "k8s.io/apimachinery/pkg/util/sets"

func (ne *NoteEvent) GetAction() string {
	if ne == nil || ne.Action == nil {
		return ""
	}

	return *ne.Action
}

func (ne *NoteEvent) GetComment() *NoteHook {
	if ne == nil {
		return nil
	}

	return ne.Comment
}

func (ne *NoteEvent) GetRepository() *ProjectHook {
	if ne == nil {
		return nil
	}

	return ne.Repository
}

func (ne *NoteEvent) GetProject() *ProjectHook {
	if ne == nil {
		return nil
	}

	return ne.Project
}

func (ne *NoteEvent) GetAuthor() *UserHook {
	if ne == nil {
		return nil
	}

	return ne.Author
}

func (ne *NoteEvent) GetSender() *UserHook {
	if ne == nil {
		return nil
	}

	return ne.Sender
}

func (ne *NoteEvent) GetURL() string {
	if ne == nil || ne.URL == nil {
		return ""
	}

	return *ne.URL
}

func (ne *NoteEvent) GetNote() string {
	if ne == nil || ne.Note == nil {
		return ""
	}

	return *ne.Note
}

func (ne *NoteEvent) GetNoteableType() string {
	if ne == nil || ne.NoteableType == nil {
		return ""
	}

	return *ne.NoteableType
}

func (ne *NoteEvent) GetTitle() string {
	if ne == nil || ne.Title == nil {
		return ""
	}

	return *ne.Title
}

func (ne *NoteEvent) GetPerIID() string {
	if ne == nil || ne.PerIID == nil {
		return ""
	}

	return *ne.PerIID
}

func (ne *NoteEvent) GetShortCommitID() string {
	if ne == nil || ne.ShortCommitID == nil {
		return ""
	}

	return *ne.ShortCommitID
}

func (ne *NoteEvent) GetEnterprise() *EnterpriseHook {
	if ne == nil {
		return nil
	}

	return ne.Enterprise
}

func (ne *NoteEvent) GetPullRequest() *PullRequestHook {
	if ne == nil {
		return nil
	}

	return ne.PullRequest
}

func (ne *NoteEvent) GetIssue() *IssueHook {
	if ne == nil {
		return nil
	}

	return ne.Issue
}

func (ne *NoteEvent) GetHookName() string {
	if ne == nil || ne.HookName == nil {
		return ""
	}

	return *ne.HookName
}

func (ne *NoteEvent) GetPassword() string {
	if ne == nil || ne.Password == nil {
		return ""
	}

	return *ne.Password
}

func (ne *NoteEvent) GetOrgRepo() (string, string) {
	return ne.GetRepository().GetOwnerAndRepo()
}

func (ne *NoteEvent) IsCreatingCommentEvent() bool {
	return ne.GetAction() == "comment"
}

func (ne *NoteEvent) GetCommenter() string {
	return ne.GetComment().GetUser().GetLogin()
}

func (ne *NoteEvent) IsIssue() bool {
	return ne.GetNoteableType() == "Issue"
}

func (ne *NoteEvent) IsIssueClosed() bool {
	return ne.GetIssue().GetState() == StatusClosed
}

func (ne *NoteEvent) IsIssueOpen() bool {
	return ne.GetIssue().GetState() == StatusOpen
}

func (ne *NoteEvent) GetIssueAuthor() string {
	return ne.GetIssue().GetUser().GetLogin()
}

func (ne *NoteEvent) GetIssueNumber() string {
	return ne.GetIssue().GetNumber()
}

func (ne *NoteEvent) GetIssueLabelSet() sets.String {
	return ne.GetIssue().LabelsToSet()
}

func (ne *NoteEvent) IsPullRequest() bool {
	return ne.GetNoteableType() == "PullRequest"
}

func (ne *NoteEvent) GetPRNumber() int32 {
	return ne.GetPullRequest().GetNumber()
}

func (ne *NoteEvent) GetPRAuthor() string {
	return ne.GetPullRequest().GetUser().GetLogin()
}

func (ne *NoteEvent) IsPROpen() bool {
	return ne.GetPullRequest().GetState() == StatusOpen
}

func (ne *NoteEvent) GetPRBaseRef() string {
	return ne.GetPullRequest().GetBase().GetRef()
}

func (ne *NoteEvent) GetPRHeadSha() string {
	return ne.GetPullRequest().GetHead().GetSha()
}

func (ne *NoteEvent) GetPRLabelSet() sets.String {
	return ne.GetPullRequest().LabelsToSet()
}
