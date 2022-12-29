package issues

import "fickle/domain/users"

type iColumn interface {
	Update(IRepository, UpdateBoardColumnParam) (Column, error)
	Remove(IRepository, RemoveBoardColumnParam) error
}

type (
	AddBoardColumnParam struct {
		Name    string
		Color   string
		Hidden  bool
		Order   string
		UserId  users.IdUser
		Default bool
	}

	UpdateBoardColumnParam struct {
		Name    *string
		Color   *string
		Hidden  *bool
		Order   *string
		Default *bool
	}

	RemoveBoardColumnParam struct {
		MoveIssuesTo *IdColumn
	}

	QueryColumnParam struct {
		Embed       QueryColumnParamEmbed
		BoardId     *IdBoard
		WorkspaceId *IdWorkspace
		Default     *bool
	}

	QueryColumnParamEmbed struct {
		Issues    bool
		Board     bool
		Workspace bool
	}

	ColumnWithEmbedDatas struct {
		Column Column
		Issues []IssueWithEmbedDatas

		Workspace Workspace
	}

	IdColumn string
)

// Implements
var _ iColumn = &Column{}

// Implements iColumn
type Column struct {
	Id      IdColumn
	Name    string
	Color   string
	Hidden  bool
	Order   string
	Default bool

	WorkspaceId IdWorkspace

	UserId users.IdUser
}
