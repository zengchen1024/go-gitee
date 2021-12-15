package gitee

import "time"

func (u *UserHook) GetID() int32 {
	if u == nil {
		return 0
	}

	return u.Id
}

func (u *UserHook) GetName() string {
	if u == nil {
		return ""
	}

	return u.Name
}

func (u *UserHook) GetEmail() string {
	if u == nil {
		return ""
	}

	return u.Email
}

func (u *UserHook) GetUserName() string {
	if u == nil {
		return ""
	}

	return u.UserName
}

func (u *UserHook) GetURL() string {
	if u == nil {
		return ""
	}

	return u.Url
}

func (u *UserHook) GetLogin() string {
	if u == nil {
		return ""
	}

	return u.Login
}

func (u *UserHook) GetAvatarURL() string {
	if u == nil {
		return ""
	}

	return u.AvatarUrl
}

func (u *UserHook) GetHtmlURL() string {
	if u == nil {
		return ""
	}

	return u.HtmlUrl
}

func (u *UserHook) GetType() string {
	if u == nil {
		return ""
	}

	return u.Type_
}

func (u *UserHook) GetSiteAdmin() bool {
	if u == nil {
		return false
	}

	return u.SiteAdmin
}

func (u *UserHook) GetTime() time.Time {
	if u == nil {
		return time.Time{}
	}

	return u.Time
}

func (u *UserHook) GetRemark() string {
	if u == nil {
		return ""
	}

	return u.Remark
}
