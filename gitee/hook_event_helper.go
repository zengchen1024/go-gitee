package gitee

import "strings"

const (
	EventTypeNote  = "Note Hook"
	EventTypePush  = "Push Hook"
	EventTypeIssue = "Issue Hook"
	EventTypePR    = "Merge Request Hook"

	StatusOpen   = "open"   //StatusOpen gitee issue or pr status is open
	StatusClosed = "closed" //StatusClosed gitee issue or pr status is closed
	StatusMerged = "merged"

	ActionOpen  = "open"
	ActionClose = "close"

	PRActionMerge               = "merge"
	PRActionUpdatedLabel        = "update_label"
	PRActionChangedTargetBranch = "target_branch_changed"
	PRActionChangedSourceBranch = "source_branch_changed"
	PRActionLinkIssue           = "update_link_issue"
)

func GetPullRequestAction(e *PullRequestEvent) string {
	switch strings.ToLower(e.GetAction()) {
	case ActionOpen:
		return ActionOpen

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

	case PRActionMerge:
		return PRActionMerge

	case ActionClose:
		return ActionClose
	}

	return ""
}
