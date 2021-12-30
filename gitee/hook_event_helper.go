package gitee

import "strings"

const (
	EventTypeNote  = "Note Hook"
	EventTypePush  = "Push Hook"
	EventTypeIssue = "Issue Hook"
	EventTypePR    = "Merge Request Hook"

	ActionOpen = "open"

	PRActionOpened              = "opened"
	PRActionClosed              = "closed"
	PRActionUpdatedLabel        = "update_label"
	PRActionChangedTargetBranch = "target_branch_changed"
	PRActionChangedSourceBranch = "source_branch_changed"
	PRActionLinkIssue           = "update_link_issue"
)

func GetPullRequestAction(e *PullRequestEvent) string {
	switch strings.ToLower(e.GetAction()) {
	case "open":
		return PRActionOpened

	case "update":
		switch strings.ToLower(e.GetActionDesc()) {
		case PRActionChangedSourceBranch: // change the pr's commits
			return PRActionChangedSourceBranch

		case PRActionChangedTargetBranch: // change the branch to which this pr will be merged
			return PRActionChangedTargetBranch

		case PRActionUpdatedLabel:
			return PRActionUpdatedLabel

		case PRActionLinkIssue:
			return PRActionLinkIssue
		}

	case "close":
		return PRActionClosed
	}

	return ""
}
