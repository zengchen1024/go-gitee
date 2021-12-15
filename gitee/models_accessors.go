package gitee

import "time"

func (bi *BasicInfo) GetLabel() string {
	if bi == nil {
		return ""
	}

	return bi.Label
}

func (bi *BasicInfo) GetRef() string {
	if bi == nil {
		return ""
	}

	return bi.Ref
}

func (bi *BasicInfo) GetSha() string {
	if bi == nil {
		return ""
	}

	return bi.Sha
}

func (bi *BasicInfo) GetUser() *UserBasic {
	if bi == nil {
		return nil
	}

	return bi.User
}

func (bi *BasicInfo) GetRepo() *Project {
	if bi == nil {
		return nil
	}

	return bi.Repo
}

func (b *Branch) GetName() string {
	if b == nil {
		return ""
	}

	return b.Name
}

func (b *Branch) GetCommit() *BranchCommit {
	if b == nil {
		return nil
	}

	return b.Commit
}

func (b *Branch) GetProtected() bool {
	if b == nil {
		return false
	}

	return b.Protected
}

func (b *Branch) GetProtectionUrl() string {
	if b == nil {
		return ""
	}

	return b.ProtectionUrl
}

func (bc *BranchCommit) GetURL() string {
	if bc == nil {
		return ""
	}

	return bc.Url
}

func (bc *BranchCommit) GetSha() string {
	if bc == nil {
		return ""
	}

	return bc.Sha
}

func (c *Commit) GetID() string {
	if c == nil {
		return ""
	}

	return c.Id
}

func (c *Commit) GetTreeId() string {
	if c == nil {
		return ""
	}

	return c.TreeId
}

func (c *Commit) GetParentIds() []string {
	if c == nil {
		return nil
	}

	return c.ParentIds
}

func (c *Commit) GetMessage() string {
	if c == nil {
		return ""
	}

	return c.Message
}

func (c *Commit) GetTimestamp() time.Time {
	if c == nil {
		return time.Time{}
	}

	return c.Timestamp
}

func (c *Commit) GetURL() string {
	if c == nil {
		return ""
	}

	return c.Url
}

func (c *Commit) GetAuthor() *UserBasic {
	if c == nil {
		return nil
	}

	return c.Author
}

func (c *Commit) GetCommitter() *UserBasic {
	if c == nil {
		return nil
	}

	return c.Committer
}

func (c *Commit) GetDistinct() bool {
	if c == nil {
		return false
	}

	return c.Distinct
}

func (c *Commit) GetAdded() []string {
	if c == nil {
		return nil
	}

	return c.Added
}

func (c *Commit) GetRemoved() []string {
	if c == nil {
		return nil
	}

	return c.Removed
}

func (c *Commit) GetModified() []string {
	if c == nil {
		return nil
	}

	return c.Modified
}

func (ct *CommitTree) GetURL() string {
	if ct == nil {
		return ""
	}

	return ct.Url
}

func (ct *CommitTree) GetSha() string {
	if ct == nil {
		return ""
	}

	return ct.Sha
}

func (cb *ContentBasic) GetName() string {
	if cb == nil {
		return ""
	}

	return cb.Name
}

func (cb *ContentBasic) GetPath() string {
	if cb == nil {
		return ""
	}

	return cb.Path
}

func (cb *ContentBasic) GetSize() string {
	if cb == nil {
		return ""
	}

	return cb.Size
}

func (cb *ContentBasic) GetSha() string {
	if cb == nil {
		return ""
	}

	return cb.Sha
}

func (cb *ContentBasic) GetUrl() string {
	if cb == nil {
		return ""
	}

	return cb.Url
}

func (cb *ContentBasic) GetHtmlUrl() string {
	if cb == nil {
		return ""
	}

	return cb.HtmlUrl
}

func (cb *ContentBasic) GetDownloadUrl() string {
	if cb == nil {
		return ""
	}

	return cb.DownloadUrl
}

func (cb *ContentBasic) GetLinks() string {
	if cb == nil {
		return ""
	}

	return cb.Links
}

func (cb *ContentBasic) GetType() string {
	if cb == nil {
		return ""
	}

	return cb.Type_
}

func (eb *EnterpriseBasic) GetID() int32 {
	if eb == nil {
		return 0
	}

	return eb.Id
}

func (eb *EnterpriseBasic) GetPath() string {
	if eb == nil {
		return ""
	}

	return eb.Path
}

func (eb *EnterpriseBasic) GetName() string {
	if eb == nil {
		return ""
	}

	return eb.Name
}

func (eb *EnterpriseBasic) GetUrl() string {
	if eb == nil {
		return ""
	}

	return eb.Url
}

func (eb *EnterpriseBasic) GetAvatarUrl() string {
	if eb == nil {
		return ""
	}

	return eb.AvatarUrl
}

func (gc *GitCommit) GetUrl() string {
	if gc == nil {
		return ""
	}

	return gc.Url
}

func (gc *GitCommit) GetAuthor() *GitUser {
	if gc == nil {
		return nil
	}

	return gc.Author
}

func (gc *GitCommit) GetCommitter() *GitUser {
	if gc == nil {
		return nil
	}

	return gc.Committer
}

func (gc *GitCommit) GetMessage() string {
	if gc == nil {
		return ""
	}

	return gc.Message
}

func (gc *GitCommit) GetCommentCount() int32 {
	if gc == nil {
		return 0
	}

	return gc.CommentCount
}

func (gu *GitUser) GetName() string {
	if gu == nil {
		return ""
	}

	return gu.Name
}

func (gu *GitUser) GetEmail() string {
	if gu == nil {
		return ""
	}

	return gu.Email
}

func (gu *GitUser) GetDate() time.Time {
	if gu == nil {
		return time.Time{}
	}

	return gu.Date
}

func (g *Group) GetID() int32 {
	if g == nil {
		return 0
	}

	return g.Id
}

func (g *Group) GetLogin() string {
	if g == nil {
		return ""
	}

	return g.Login
}

func (g *Group) GetUrl() string {
	if g == nil {
		return ""
	}

	return g.Url
}

func (g *Group) GetAvatarUrl() string {
	if g == nil {
		return ""
	}

	return g.AvatarUrl
}

func (g *Group) GetReposUrl() string {
	if g == nil {
		return ""
	}

	return g.ReposUrl
}

func (g *Group) GetEventsUrl() string {
	if g == nil {
		return ""
	}

	return g.EventsUrl
}

func (g *Group) GetMembersUrl() string {
	if g == nil {
		return ""
	}

	return g.MembersUrl
}

func (g *Group) GetDescription() string {
	if g == nil {
		return ""
	}

	return g.Description
}

func (m *Milestone) GetURL() string {
	if m == nil {
		return ""
	}

	return m.Url
}

func (m *Milestone) GetHtmlUrl() string {
	if m == nil {
		return ""
	}

	return m.HtmlUrl
}

func (m *Milestone) GetState() string {
	if m == nil {
		return ""
	}

	return m.State
}

func (m *Milestone) GetTitle() string {
	if m == nil {
		return ""
	}

	return m.Title
}

func (m *Milestone) GetDescription() string {
	if m == nil {
		return ""
	}

	return m.Description
}

func (m *Milestone) GetDueOn() string {
	if m == nil {
		return ""
	}

	return m.DueOn
}

func (m *Milestone) GetNumber() int32 {
	if m == nil {
		return 0
	}

	return m.Number
}

func (m *Milestone) GetRepositoryId() int32 {
	if m == nil {
		return 0
	}

	return m.RepositoryId
}

func (m *Milestone) GetOpenIssues() int32 {
	if m == nil {
		return 0
	}

	return m.OpenIssues
}

func (m *Milestone) GetClosedIssues() int32 {
	if m == nil {
		return 0
	}

	return m.ClosedIssues
}

func (m *Milestone) GetUpdatedAt() time.Time {
	if m == nil {
		return time.Time{}
	}

	return m.UpdatedAt
}

func (m *Milestone) GetCreatedAt() time.Time {
	if m == nil {
		return time.Time{}
	}

	return m.CreatedAt
}

func (n *Namespace) GetID() int32 {
	if n == nil {
		return 0
	}

	return n.Id
}

func (n *Namespace) GetType() string {
	if n == nil {
		return ""
	}

	return n.Type_
}

func (n *Namespace) GetName() string {
	if n == nil {
		return ""
	}

	return n.Name
}

func (n *Namespace) GetPath() string {
	if n == nil {
		return ""
	}

	return n.Path
}

func (n *Namespace) GetHtmlUrl() string {
	if n == nil {
		return ""
	}

	return n.HtmlUrl
}

func (n *Namespace) GetParent() *NamespaceMini {
	if n == nil {
		return nil
	}

	return n.Parent
}

func (nm *NamespaceMini) GetID() int32 {
	if nm == nil {
		return 0
	}

	return nm.Id
}

func (nm *NamespaceMini) GetType() string {
	if nm == nil {
		return ""
	}

	return nm.Type_
}

func (nm *NamespaceMini) GetName() string {
	if nm == nil {
		return ""
	}

	return nm.Name
}

func (nm *NamespaceMini) GetPath() string {
	if nm == nil {
		return ""
	}

	return nm.Path
}

func (nm *NamespaceMini) GetHtmlUrl() string {
	if nm == nil {
		return ""
	}

	return nm.HtmlUrl
}

func (pb *ProgramBasic) GetID() string {
	if pb == nil {
		return ""
	}

	return pb.Id
}

func (pb *ProgramBasic) GetName() string {
	if pb == nil {
		return ""
	}

	return pb.Name
}

func (pb *ProgramBasic) GetDescription() string {
	if pb == nil {
		return ""
	}

	return pb.Description
}

func (pb *ProgramBasic) GetAssignee() string {
	if pb == nil {
		return ""
	}

	return pb.Assignee
}

func (pb *ProgramBasic) GetAuthor() string {
	if pb == nil {
		return ""
	}

	return pb.Author
}

func (pro *Project) GetID() int32 {
	if pro == nil {
		return 0
	}

	return pro.Id
}

func (pro *Project) GetFullName() string {
	if pro == nil {
		return ""
	}

	return pro.FullName
}

func (pro *Project) GetHumanName() string {
	if pro == nil {
		return ""
	}

	return pro.HumanName
}

func (pro *Project) GetUrl() string {
	if pro == nil {
		return ""
	}

	return pro.Url
}

func (pro *Project) GetPath() string {
	if pro == nil {
		return ""
	}

	return pro.Path
}

func (pro *Project) GetName() string {
	if pro == nil {
		return ""
	}

	return pro.Name
}

func (pro *Project) GetDescription() string {
	if pro == nil {
		return ""
	}

	return pro.Description
}

func (pro *Project) GetHtmlUrl() string {
	if pro == nil {
		return ""
	}

	return pro.HtmlUrl
}

func (pro *Project) GetSshUrl() string {
	if pro == nil {
		return ""
	}

	return pro.SshUrl
}

func (pro *Project) GetForksUrl() string {
	if pro == nil {
		return ""
	}

	return pro.ForksUrl
}

func (pro *Project) GetKeysUrl() string {
	if pro == nil {
		return ""
	}

	return pro.KeysUrl
}

func (pro *Project) GetCollaboratorsUrl() string {
	if pro == nil {
		return ""
	}

	return pro.CollaboratorsUrl
}

func (pro *Project) GetHooksUrl() string {
	if pro == nil {
		return ""
	}

	return pro.HooksUrl
}

func (pro *Project) GetTagsUrl() string {
	if pro == nil {
		return ""
	}

	return pro.TagsUrl
}

func (pro *Project) GetBranchesUrl() string {
	if pro == nil {
		return ""
	}

	return pro.BranchesUrl
}

func (pro *Project) GetBlobsUrl() string {
	if pro == nil {
		return ""
	}

	return pro.BlobsUrl
}

func (pro *Project) GetStargazersUrl() string {
	if pro == nil {
		return ""
	}

	return pro.StargazersUrl
}

func (pro *Project) GetContributorsUrl() string {
	if pro == nil {
		return ""
	}

	return pro.ContributorsUrl
}

func (pro *Project) GetCommitsUrl() string {
	if pro == nil {
		return ""
	}

	return pro.CommitsUrl
}

func (pro *Project) GetCommentsUrl() string {
	if pro == nil {
		return ""
	}

	return pro.CommentsUrl
}

func (pro *Project) GetIssueCommentUrl() string {
	if pro == nil {
		return ""
	}

	return pro.IssueCommentUrl
}

func (pro *Project) GetIssuesUrl() string {
	if pro == nil {
		return ""
	}

	return pro.IssuesUrl
}

func (pro *Project) GetPullsUrl() string {
	if pro == nil {
		return ""
	}

	return pro.PullsUrl
}

func (pro *Project) GetMilestonesUrl() string {
	if pro == nil {
		return ""
	}

	return pro.MilestonesUrl
}

func (pro *Project) GetNotificationsUrl() string {
	if pro == nil {
		return ""
	}

	return pro.NotificationsUrl
}

func (pro *Project) GetLabelsUrl() string {
	if pro == nil {
		return ""
	}

	return pro.LabelsUrl
}

func (pro *Project) GetReleasesUrl() string {
	if pro == nil {
		return ""
	}

	return pro.ReleasesUrl
}

func (pro *Project) GetHomepage() string {
	if pro == nil {
		return ""
	}

	return pro.Homepage
}

func (pro *Project) GetLanguage() string {
	if pro == nil {
		return ""
	}

	return pro.Language
}

func (pro *Project) GetDefaultBranch() string {
	if pro == nil {
		return ""
	}

	return pro.DefaultBranch
}

func (pro *Project) GetLicense() string {
	if pro == nil {
		return ""
	}

	return pro.License
}

func (pro *Project) GetPushedAt() string {
	if pro == nil {
		return ""
	}

	return pro.PushedAt
}

func (pro *Project) GetCreatedAt() string {
	if pro == nil {
		return ""
	}

	return pro.CreatedAt
}

func (pro *Project) GetUpdatedAt() string {
	if pro == nil {
		return ""
	}

	return pro.UpdatedAt
}

func (pro *Project) GetPaas() string {
	if pro == nil {
		return ""
	}

	return pro.Paas
}

func (pro *Project) GetStared() string {
	if pro == nil {
		return ""
	}

	return pro.Stared
}

func (pro *Project) GetWatched() string {
	if pro == nil {
		return ""
	}

	return pro.Watched
}

func (pro *Project) GetPermission() string {
	if pro == nil {
		return ""
	}

	return pro.Permission
}

func (pro *Project) GetRelation() string {
	if pro == nil {
		return ""
	}

	return pro.Relation
}

func (pro *Project) GetPrivate() bool {
	if pro == nil {
		return false
	}

	return pro.Private
}

func (pro *Project) GetPublic() bool {
	if pro == nil {
		return false
	}

	return pro.Public
}

func (pro *Project) GetInternal() bool {
	if pro == nil {
		return false
	}

	return pro.Internal
}

func (pro *Project) GetFork() bool {
	if pro == nil {
		return false
	}

	return pro.Fork
}

func (pro *Project) GetRecommend() bool {
	if pro == nil {
		return false
	}

	return pro.Recommend
}

func (pro *Project) GetHasIssues() bool {
	if pro == nil {
		return false
	}

	return pro.HasIssues
}

func (pro *Project) GetHasWiki() bool {
	if pro == nil {
		return false
	}

	return pro.HasWiki
}

func (pro *Project) GetCanComment() bool {
	if pro == nil {
		return false
	}

	return pro.CanComment
}

func (pro *Project) GetPullRequestsEnabled() bool {
	if pro == nil {
		return false
	}

	return pro.PullRequestsEnabled
}

func (pro *Project) GetHasPage() bool {
	if pro == nil {
		return false
	}

	return pro.HasPage
}

func (pro *Project) GetOutsourced() bool {
	if pro == nil {
		return false
	}

	return pro.Outsourced
}

func (pro *Project) GetNamespace() *Namespace {
	if pro == nil {
		return nil
	}

	return pro.Namespace
}

func (pro *Project) GetOwner() *UserBasic {
	if pro == nil {
		return nil
	}

	return pro.Owner
}

func (pro *Project) GetParent() *Project {
	if pro == nil {
		return nil
	}

	return pro.Parent
}

func (pro *Project) GetMembers() []string {
	if pro == nil {
		return nil
	}

	return pro.Members
}

func (pro *Project) GetForksCount() int32 {
	if pro == nil {
		return 0
	}

	return pro.ForksCount
}

func (pro *Project) GetStargazersCount() int32 {
	if pro == nil {
		return 0
	}

	return pro.StargazersCount
}

func (pro *Project) GetWatchersCount() int32 {
	if pro == nil {
		return 0
	}

	return pro.WatchersCount
}

func (pro *Project) GetOpenIssuesCount() int32 {
	if pro == nil {
		return 0
	}

	return pro.OpenIssuesCount
}

func (pb *ProjectBasic) GetPrivate() bool {
	if pb == nil {
		return false
	}

	return pb.Private
}

func (pb *ProjectBasic) GetPublic() bool {
	if pb == nil {
		return false
	}

	return pb.Public
}

func (pb *ProjectBasic) GetInternal() bool {
	if pb == nil {
		return false
	}

	return pb.Internal
}

func (pb *ProjectBasic) GetID() int32 {
	if pb == nil {
		return 0
	}

	return pb.Id
}

func (pb *ProjectBasic) GetFullName() string {
	if pb == nil {
		return ""
	}

	return pb.FullName
}

func (pb *ProjectBasic) GetHumanName() string {
	if pb == nil {
		return ""
	}

	return pb.HumanName
}

func (pb *ProjectBasic) GetUrl() string {
	if pb == nil {
		return ""
	}

	return pb.Url
}

func (pb *ProjectBasic) GetPath() string {
	if pb == nil {
		return ""
	}

	return pb.Path
}

func (pb *ProjectBasic) GetName() string {
	if pb == nil {
		return ""
	}

	return pb.Name
}

func (pb *ProjectBasic) GetDescription() string {
	if pb == nil {
		return ""
	}

	return pb.Description
}

func (pb *ProjectBasic) GetHtmlUrl() string {
	if pb == nil {
		return ""
	}

	return pb.HtmlUrl
}

func (pb *ProjectBasic) GetSshUrl() string {
	if pb == nil {
		return ""
	}

	return pb.SshUrl
}

func (pb *ProjectBasic) GetOwner() *UserBasic {
	if pb == nil {
		return nil
	}

	return pb.Owner
}

func (pmp *ProjectMemberPermissionDetail) GetPull() bool {
	if pmp == nil {
		return false
	}

	return pmp.Pull
}

func (pmp *ProjectMemberPermissionDetail) GetPush() bool {
	if pmp == nil {
		return false
	}

	return pmp.Push
}

func (pmp *ProjectMemberPermissionDetail) GetAdmin() bool {
	if pmp == nil {
		return false
	}

	return pmp.Admin
}

func (pfp *PullRequestFilePath) GetNewFile() bool {
	if pfp == nil {
		return false
	}

	return pfp.NewFile
}

func (pfp *PullRequestFilePath) GetRenamedFile() bool {
	if pfp == nil {
		return false
	}

	return pfp.RenamedFile
}

func (pfp *PullRequestFilePath) GetDeletedFile() bool {
	if pfp == nil {
		return false
	}

	return pfp.DeletedFile
}

func (pfp *PullRequestFilePath) GetTooLarge() bool {
	if pfp == nil {
		return false
	}

	return pfp.TooLarge
}

func (pfp *PullRequestFilePath) GetDiff() string {
	if pfp == nil {
		return ""
	}

	return pfp.Diff
}

func (pfp *PullRequestFilePath) GetNewPath() string {
	if pfp == nil {
		return ""
	}

	return pfp.NewPath
}

func (pfp *PullRequestFilePath) GetOldPath() string {
	if pfp == nil {
		return ""
	}

	return pfp.OldPath
}

func (pfp *PullRequestFilePath) GetAMode() string {
	if pfp == nil {
		return ""
	}

	return pfp.AMode
}

func (pfp *PullRequestFilePath) GetBMode() string {
	if pfp == nil {
		return ""
	}

	return pfp.BMode
}

func (ub *UserBasic) GetID() int32 {
	if ub == nil {
		return 0
	}

	return ub.Id
}

func (ub *UserBasic) GetSiteAdmin() bool {
	if ub == nil {
		return false
	}

	return ub.SiteAdmin
}

func (ub *UserBasic) GetLogin() string {
	if ub == nil {
		return ""
	}

	return ub.Login
}
func (ub *UserBasic) GetFollowingUrl() string {
	if ub == nil {
		return ""
	}

	return ub.FollowingUrl
}

func (ub *UserBasic) GetGistsUrl() string {
	if ub == nil {
		return ""
	}

	return ub.GistsUrl
}

func (ub *UserBasic) GetStarredUrl() string {
	if ub == nil {
		return ""
	}

	return ub.StarredUrl
}

func (ub *UserBasic) GetSubscriptionsUrl() string {
	if ub == nil {
		return ""
	}

	return ub.SubscriptionsUrl
}

func (ub *UserBasic) GetOrganizationsUrl() string {
	if ub == nil {
		return ""
	}

	return ub.OrganizationsUrl
}

func (ub *UserBasic) GetReposUrl() string {
	if ub == nil {
		return ""
	}

	return ub.ReposUrl
}

func (ub *UserBasic) GetEventsUrl() string {
	if ub == nil {
		return ""
	}

	return ub.EventsUrl
}

func (ub *UserBasic) GetReceivedEventsUrl() string {
	if ub == nil {
		return ""
	}

	return ub.ReceivedEventsUrl
}

func (ub *UserBasic) GetType() string {
	if ub == nil {
		return ""
	}

	return ub.Type_
}

func (ub *UserBasic) GetEmail() string {
	if ub == nil {
		return ""
	}

	return ub.Email
}

func (ub *UserBasic) GetAvatarUrl() string {
	if ub == nil {
		return ""
	}

	return ub.AvatarUrl
}

func (ub *UserBasic) GetUrl() string {
	if ub == nil {
		return ""
	}

	return ub.Url
}

func (ub *UserBasic) GetHtmlUrl() string {
	if ub == nil {
		return ""
	}

	return ub.HtmlUrl
}

func (ub *UserBasic) GetFollowersUrl() string {
	if ub == nil {
		return ""
	}

	return ub.FollowersUrl
}

func (ub *UserBasic) GetName() string {
	if ub == nil {
		return ""
	}

	return ub.Name
}
