package issues

type IRepository interface {
	CreateWorkspace(Workspace) (Workspace, error)
	UpdateWorkspace(IdWorkspace, UpdateWorkspaceParam) (Workspace, error)
	RemoveWorkspcae(IdWorkspace) error
	FindWorkspace(IdWorkspace, QueryWorkspaceParam) (WorkspaceWithEmbedDatas, error)
	FindWorkspaces(QueryWorkspaceParam) ([]WorkspaceWithEmbedDatas, error)

	CreateColumn(Column) (Column, error)
	UpdateColumn(IdColumn, UpdateBoardColumnParam) (Column, error)
	UpdateColumns(UpdateBoardColumnParam, QueryColumnParam) error
	RemoveColumn(IdColumn) error
	FindColumn(IdColumn, QueryColumnParam) (ColumnWithEmbedDatas, error)
	FindColumns(QueryColumnParam) ([]ColumnWithEmbedDatas, error)

	CreateBoard(Board) (Board, error)
	UpdateBoard(IdBoard, UpdateBoardParam) (Board, error)
	RemoveBoard(IdBoard) error
	FindBoard(IdBoard, QueryBoardParam) (BoardWithEmbedDatas, error)
	FindBoards(QueryBoardParam) ([]BoardWithEmbedDatas, error)

	CreateIssue(Issue) (Issue, error)
	UpdateIssue(IdIssue, UpdateIssueParam) (Issue, error)
	UpdateIssues(UpdateIssueParam, QueryIssueParam) error
	RemoveIssue(IdIssue) error
	FindIssue(IdIssue, QueryIssueParam) (IssueWithEmbedDatas, error)
	FindIssues(QueryIssueParam) ([]IssueWithEmbedDatas, error)
}
