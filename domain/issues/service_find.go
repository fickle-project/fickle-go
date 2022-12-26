package issues

// FindBoard implements iBoardService
func (*boardService) FindBoard(r IRepository, id IdBoard, p QueryBoardParam) (BoardWithEmbedDatas, error) {
	return r.FindBoard(id, p)
}

// FindBoards implements iBoardService
func (*boardService) FindBoards(r IRepository, p QueryBoardParam) ([]BoardWithEmbedDatas, error) {
	return r.FindBoards(p)
}

// FindIssue implements iBoardService
func (*boardService) FindIssue(r IRepository, id IdIssue, p QueryIssueParam) (IssueWithEmbedDatas, error) {
	return r.FindIssue(id, p)
}

// FindIssues implements iBoardService
func (*boardService) FindIssues(r IRepository, p QueryIssueParam) ([]IssueWithEmbedDatas, error) {
	return r.FindIssues(p)
}

// FindWorkspace implements iBoardService
func (*boardService) FindWorkspace(r IRepository, id IdWorkspace, p QueryWorkspaceParam) (WorkspaceWithEmbedDatas, error) {
	return r.FindWorkspace(id, p)
}

// FindWorkspaces implements iBoardService
func (b *boardService) FindWorkspaces(r IRepository, p QueryWorkspaceParam) ([]WorkspaceWithEmbedDatas, error) {
	return r.FindWorkspaces(p)
}
