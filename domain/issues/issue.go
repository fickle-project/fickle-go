package issues

import (
	"fickle/domain/users"
)

type iIssue interface {
	Update(IRepository, UpdateIssueParam) (Issue, error)
	Remove(IRepository) error
}

type (
	NewIssueParam struct {
		Name     string
		Content  *string
		ColumnId *IdColumn
		Order    string
		UserId   users.IdUser
	}

	UpdateIssueParam struct {
		Name     *string
		Content  **string
		BoardId  *IdBoard
		ColumnId **IdColumn
		Order    *string
	}

	QueryIssueParam struct {
		Embed       QueryIssueParamEmbed
		BoardId     *IdBoard
		ColumnId    **IdColumn
		WorkspaceId *IdWorkspace
		Name        *string
		NameContain *string
	}

	QueryIssueParamEmbed struct {
		Column    bool
		Board     bool
		Workspace bool
	}

	IssueWithEmbedDatas struct {
		Issue Issue

		Column    Column
		Board     Board
		Workspace Workspace
	}

	IdIssue string
)

// Implements
var _ iIssue = &Issue{}

// Implements iIssue
type Issue struct {
	Id      IdIssue
	Name    string
	Content *string

	BoardId  IdBoard
	ColumnId *IdColumn
	Order    string

	WorkspaceId IdWorkspace

	UserId users.IdUser
}
