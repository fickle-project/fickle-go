package issues

type IFactory interface {
	NewWorkspaceId(r IRepository) (IdWorkspace, error)
	NewColumnId(r IRepository) (IdColumn, error)
	NewBoardId(r IRepository) (IdBoard, error)
	NewIssueId(r IRepository) (IdIssue, error)
}
