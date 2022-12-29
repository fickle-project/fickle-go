package issues

type iBoardService interface {
	NewWorkspace(IFactory, IRepository, CreateWorkspaceParam) (Workspace, error)
	FindWorkspace(IRepository, IdWorkspace, QueryWorkspaceParam) (WorkspaceWithEmbedDatas, error)
	FindWorkspaces(IRepository, QueryWorkspaceParam) ([]WorkspaceWithEmbedDatas, error)
	FindBoard(IRepository, IdBoard, QueryBoardParam) (BoardWithEmbedDatas, error)
	FindBoards(IRepository, QueryBoardParam) ([]BoardWithEmbedDatas, error)
	FindIssue(IRepository, IdIssue, QueryIssueParam) (IssueWithEmbedDatas, error)
	FindIssues(IRepository, QueryIssueParam) ([]IssueWithEmbedDatas, error)
}

func NewService() iBoardService {
	return &boardService{}
}

// Implements iBoardService
type boardService struct{}
