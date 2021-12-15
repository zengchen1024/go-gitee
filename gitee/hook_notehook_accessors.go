package gitee

func (nh *NoteHook) GetBody() string {
	if nh == nil {
		return ""
	}

	return nh.Body
}

func (nh *NoteHook) GetID() int32 {
	if nh == nil {
		return 0
	}

	return nh.Id
}

func (nh *NoteHook) GetCreateAt() string {
	if nh == nil {
		return ""
	}

	return nh.CreatedAt
}

func (nh *NoteHook) GetUpdatedAt() string {
	if nh == nil {
		return ""
	}

	return nh.UpdatedAt
}

func (nh *NoteHook) GetHtmlUrl() string {
	if nh == nil {
		return ""
	}

	return nh.HtmlUrl
}

func (nh *NoteHook) GetPosition() string {
	if nh == nil {
		return ""
	}

	return nh.Position
}

func (nh *NoteHook) GetCommitID() string {
	if nh == nil {
		return ""
	}

	return nh.CommitId
}
