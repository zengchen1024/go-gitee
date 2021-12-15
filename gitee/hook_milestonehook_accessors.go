package gitee

import "time"

func (m *MilestoneHook) GetID() int32 {
	if m == nil {
		return 0
	}

	return m.Id
}

func (m *MilestoneHook) GetHtmlUrl() string {
	if m == nil {
		return ""
	}

	return m.HtmlUrl
}

func (m *MilestoneHook) GetNumber() int32 {
	if m == nil {
		return 0
	}

	return m.Number
}

func (m *MilestoneHook) GetTitle() string {
	if m == nil {
		return ""
	}

	return m.Title
}

func (m *MilestoneHook) GetDescription() string {
	if m == nil {
		return ""
	}

	return m.Description
}

func (m *MilestoneHook) GetOpenIssues() int32 {
	if m == nil {
		return 0
	}

	return m.OpenIssues
}

func (m *MilestoneHook) GetClosedIssues() int32 {
	if m == nil {
		return 0
	}

	return m.ClosedIssues
}

func (m *MilestoneHook) GetState() string {
	if m == nil {
		return ""
	}

	return m.State
}

func (m *MilestoneHook) GetCreatedAt() time.Time {
	if m == nil {
		return time.Time{}
	}

	return m.CreatedAt
}

func (m *MilestoneHook) GetUpdatedAt() time.Time {
	if m == nil {
		return time.Time{}
	}

	return m.UpdatedAt
}

func (m *MilestoneHook) GetDueOn() string {
	if m == nil {
		return ""
	}

	return m.DueOn
}
