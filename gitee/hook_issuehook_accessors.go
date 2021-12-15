package gitee

import (
	"k8s.io/apimachinery/pkg/util/sets"
	"time"
)

func (ih *IssueHook) GetID() int32 {
	if ih == nil {
		return 0
	}

	return ih.Id
}

func (ih *IssueHook) GetHtmlUrl() string {
	if ih == nil {
		return ""
	}

	return ih.HtmlUrl
}

func (ih *IssueHook) GetNumber() string {
	if ih == nil {
		return ""
	}

	return ih.Number
}

func (ih *IssueHook) GetTitle() string {
	if ih == nil {
		return ""
	}

	return ih.Title
}

func (ih *IssueHook) GetUser() *UserHook {
	if ih == nil {
		return nil
	}

	return ih.User
}

func (ih *IssueHook) GetLabels() []LabelHook {
	if ih == nil {
		return nil
	}

	return ih.Labels
}

func (ih *IssueHook) GetState() string {
	if ih == nil {
		return ""
	}

	return ih.State
}

func (ih *IssueHook) GetStateName() string {
	if ih == nil {
		return ""
	}

	return ih.StateName
}

func (ih *IssueHook) GetTypeName() string {
	if ih == nil {
		return ""
	}

	return ih.TypeName
}

func (ih *IssueHook) GetAssignee() *UserHook {
	if ih == nil {
		return nil
	}

	return ih.Assignee
}

func (ih *IssueHook) GetCollaborators() []UserHook {
	if ih == nil {
		return nil
	}

	return ih.Collaborators
}

func (ih *IssueHook) GetMilestone() *MilestoneHook {
	if ih == nil {
		return nil
	}

	return ih.Milestone
}

func (ih *IssueHook) GetComments() int32 {
	if ih == nil {
		return 0
	}

	return ih.Comments
}

func (ih *IssueHook) GetCreatedAt() time.Time {
	if ih == nil {
		return time.Time{}
	}

	return ih.CreatedAt
}

func (ih *IssueHook) GetUpdatedAt() time.Time {
	if ih == nil {
		return time.Time{}
	}

	return ih.UpdatedAt
}

func (ih *IssueHook) GetBody() string {
	if ih == nil {
		return ""
	}

	return ih.Body
}

func (ih *IssueHook) LabelsToSet() sets.String {
	res := sets.NewString()

	for _, v := range ih.GetLabels() {
		res.Insert(v.GetName())
	}

	if res.Has("") {
		res.Delete("")
	}

	return res
}
