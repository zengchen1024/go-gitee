package gitee

import "time"

func (c *CommitHook) GetID() string {
	if c == nil {
		return ""
	}

	return c.Id
}

func (c *CommitHook) GetTreeId() string {
	if c == nil {
		return ""
	}

	return c.TreeId
}

func (c *CommitHook) GetParentIds() []string {
	if c == nil {
		return nil
	}

	return c.ParentIds
}

func (c *CommitHook) GetMessage() string {
	if c == nil {
		return ""
	}

	return c.Message
}

func (c *CommitHook) GetTimestamp() time.Time {
	if c == nil {
		return time.Time{}
	}

	return c.Timestamp
}

func (c *CommitHook) GetURL() string {
	if c == nil {
		return ""
	}

	return c.Url
}
func (c *CommitHook) GetAuthor() *UserHook {
	if c == nil {
		return nil
	}

	return c.Author
}

func (c *CommitHook) GetCommitter() *UserHook {
	if c == nil {
		return nil
	}

	return c.Committer
}

func (c *CommitHook) GetDistinct() bool {
	if c == nil {
		return false
	}

	return c.Distinct
}

func (c *CommitHook) GetAdded() []string {
	if c == nil {
		return nil
	}

	return c.Added
}

func (c *CommitHook) GetRemoved() []string {
	if c == nil {
		return nil
	}

	return c.Removed
}

func (c *CommitHook) GetModified() []string {
	if c == nil {
		return nil
	}

	return c.Modified
}
