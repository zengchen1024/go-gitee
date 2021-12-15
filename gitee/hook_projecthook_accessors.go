package gitee

func (pj *ProjectHook) GetID() int32 {
	if pj == nil {
		return 0
	}

	return pj.Id
}

func (pj *ProjectHook) GetName() string {
	if pj == nil {
		return ""
	}

	return pj.Name
}

func (pj *ProjectHook) GetPath() string {
	if pj == nil {
		return ""
	}

	return pj.Path
}

func (pj *ProjectHook) GetFullName() string {
	if pj == nil {
		return ""
	}

	return pj.FullName
}

func (pj *ProjectHook) GetOwner() *UserHook {
	if pj == nil {
		return nil
	}

	return pj.Owner
}

func (pj *ProjectHook) GetPrivate() bool {
	if pj == nil {
		return false
	}

	return pj.Private
}

func (pj *ProjectHook) GetHtmlUrl() string {
	if pj == nil {
		return ""
	}

	return pj.HtmlUrl
}

func (pj *ProjectHook) GetUrl() string {
	if pj == nil {
		return ""
	}

	return pj.Url
}

func (pj *ProjectHook) GetDescription() string {
	if pj == nil {
		return ""
	}

	return pj.Description
}

func (pj *ProjectHook) GetFork() bool {
	if pj == nil {
		return false
	}

	return pj.Fork
}

func (pj *ProjectHook) GetPushedAt() string {
	if pj == nil {
		return ""
	}

	return pj.PushedAt
}

func (pj *ProjectHook) GetCreatedAt() string {
	if pj == nil {
		return ""
	}

	return pj.CreatedAt
}

func (pj *ProjectHook) GetUpdatedAt() string {
	if pj == nil {
		return ""
	}

	return pj.UpdatedAt
}

func (pj *ProjectHook) GetSshUrl() string {
	if pj == nil {
		return ""
	}

	return pj.SshUrl
}

func (pj *ProjectHook) GetGitUrl() string {
	if pj == nil {
		return ""
	}

	return pj.GitUrl
}

func (pj *ProjectHook) GetCloneUrl() string {
	if pj == nil {
		return ""
	}

	return pj.CloneUrl
}

func (pj *ProjectHook) GetSvnUrl() string {
	if pj == nil {
		return ""
	}

	return pj.SvnUrl
}

func (pj *ProjectHook) GetGitHttpUrl() string {
	if pj == nil {
		return ""
	}

	return pj.GitHttpUrl
}

func (pj *ProjectHook) GetGitSshUrl() string {
	if pj == nil {
		return ""
	}

	return pj.GitSshUrl
}

func (pj *ProjectHook) GetGitSvnUrl() string {
	if pj == nil {
		return ""
	}

	return pj.GitSvnUrl
}

func (pj *ProjectHook) GetHomepage() string {
	if pj == nil {
		return ""
	}

	return pj.Homepage
}

func (pj *ProjectHook) GetStargazersCount() int32 {
	if pj == nil {
		return 0
	}

	return pj.StargazersCount
}

func (pj *ProjectHook) GetWatchersCount() int32 {
	if pj == nil {
		return 0
	}

	return pj.WatchersCount
}

func (pj *ProjectHook) GetForksCount() int32 {
	if pj == nil {
		return 0
	}

	return pj.ForksCount
}

func (pj *ProjectHook) GetLanguage() string {
	if pj == nil {
		return ""
	}

	return pj.Language
}

func (pj *ProjectHook) GetHasIssues() bool {
	if pj == nil {
		return false
	}

	return pj.HasIssues
}

func (pj *ProjectHook) GetHasWiki() bool {
	if pj == nil {
		return false
	}

	return pj.HasWiki
}

func (pj *ProjectHook) GetHasPage() bool {
	if pj == nil {
		return false
	}

	return pj.HasPage
}

func (pj *ProjectHook) GetLicense() string {
	if pj == nil {
		return ""
	}

	return pj.License
}

func (pj *ProjectHook) GetOpenIssuesCount() int32 {
	if pj == nil {
		return 0
	}

	return pj.OpenIssuesCount
}

func (pj *ProjectHook) GetDefaultBranch() string {
	if pj == nil {
		return ""
	}

	return pj.DefaultBranch
}

func (pj *ProjectHook) GetNameSpace() string {
	if pj == nil {
		return ""
	}

	return pj.Namespace
}

func (pj *ProjectHook) GetNameWithNamespace() string {
	if pj == nil {
		return ""
	}

	return pj.NameWithNamespace
}

func (pj *ProjectHook) GetPathWithNamespace() string {
	if pj == nil {
		return ""
	}

	return pj.PathWithNamespace
}

func (pj *ProjectHook) GetOwnerAndRepo() (string, string) {
	if pj == nil {
		return "", ""
	}

	return pj.Namespace, pj.Path
}
