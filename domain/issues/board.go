package issues

import (
	"fickle/domain/users"
)

type iBoard interface {
	Update(IRepository, UpdateBoardParam) (Board, error)
	Remove(IRepository) error
	NewIssue(IFactory, IRepository, NewIssueParam) (Issue, error)
}

type (
	AddBoardParam struct {
		Name string
	}

	UpdateBoardParam struct {
		Name     *string
		Archived *bool
	}

	QueryBoardParam struct {
		Embed           []string
		Name            *string
		NameContain     *string
		WorkspaceId     *IdWorkspace
		IncludeArchived bool
	}

	BoardWithEmbedDatas struct {
		Board   Board
		Columns []ColumnWithEmbedDatas

		Workspace Workspace
	}

	IdBoard string
)

// Implements
var _ iBoard = &Board{}

// Implements iBoard
type Board struct {
	Id   IdBoard
	Name string

	WorkspaceId IdWorkspace

	Archived bool

	UserId users.IdUser
}
