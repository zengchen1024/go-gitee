package gitee

import "strings"

const (
	PRActionOpened              = "opened"
	PRActionClosed              = "closed"
	PRActionUpdatedLabel        = "update_label"
	PRActionChangedTargetBranch = "target_branch_changed"
	PRActionChangedSourceBranch = "source_branch_changed"
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
		}

	case "close":
		return PRActionClosed
	}

	return ""
}
