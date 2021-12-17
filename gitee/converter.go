package gitee

import (
	"encoding/json"
	"fmt"
)

func ConvertToNoteEvent(payload []byte) (e NoteEvent, err error) {
	if err = json.Unmarshal(payload, &e); err != nil {
		return
	}

	err = checkNoteEvent(&e)
	return
}

func checkNoteEvent(e *NoteEvent) error {
	eventType := EventTypeNote

	if e.Comment == nil {
		return fmtCheckError(eventType, "Comment")
	}

	if e.IsPullRequest() {
		err := checkPullRequestHook(e.PullRequest, eventType, "PullRequest")
		if err != nil {
			return err
		}
	}

	if e.IsIssue() && e.Issue == nil {
		return fmtCheckError(eventType, "Issue")
	}

	return checkRepository(e.Repository, eventType)
}

func ConvertToIssueEvent(payload []byte) (e IssueEvent, err error) {
	if err = json.Unmarshal(payload, &e); err != nil {
		return
	}

	err = checkIssueEvent(&e)
	return
}

func checkIssueEvent(e *IssueEvent) error {
	eventType := EventTypeIssue

	if e.Issue == nil {
		return fmtCheckError(eventType, "Issue")
	}

	return checkRepository(e.Repository, eventType)
}

func ConvertToPREvent(payload []byte) (e PullRequestEvent, err error) {
	if err = json.Unmarshal(payload, &e); err != nil {
		return
	}

	err = checkPullRequestEvent(&e)
	return
}

func checkPullRequestEvent(e *PullRequestEvent) error {
	eventType := EventTypePR

	if err := checkPullRequestHook(e.PullRequest, eventType, "PullRequest"); err != nil {
		return err
	}

	return checkRepository(e.Repository, eventType)
}

func checkPullRequestHook(pr *PullRequestHook, eventType, field string) error {
	if pr == nil {
		return fmtCheckError(eventType, field)
	}

	if pr.Head == nil || pr.Base == nil {
		return fmtCheckError(eventType, field+".Head or "+field+".Base")
	}
	return nil
}

func checkRepository(rep *ProjectHook, eventType string) error {
	field := "Repository"

	if rep == nil {
		return fmtCheckError(eventType, field)
	}

	org, repo := rep.GetOwnerAndRepo()
	if org == "" || repo == "" {
		return fmtCheckError(eventType, field+" .org or .repo")
	}
	return nil
}

func ConvertToPushEvent(payload []byte) (e PushEvent, err error) {
	if err = json.Unmarshal(payload, &e); err != nil {
		return
	}

	err = checkRepository(e.Repository, EventTypePush)
	return
}

func fmtCheckError(eventType, field string) error {
	return fmt.Errorf("%s is illegal: the field of '%s' is empty", eventType, field)
}
