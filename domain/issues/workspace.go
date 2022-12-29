package issues

import "fickle/domain/users"

type iWorkspace interface {
	Update(IRepository, UpdateWorkspaceParam) (Workspace, error)
	Remove(IRepository) error
	NewColumn(IFactory, IRepository, AddBoardColumnParam) (Column, error)
	NewBoard(IFactory, IRepository, AddBoardParam) (Board, error)
}

type (
	CreateWorkspaceParam struct {
		Name    string
		Columns []AddBoardColumnParam
		UserId  users.IdUser
	}

	UpdateWorkspaceParam struct {
		Name     *string
		Archived *bool
	}

	QueryWorkspaceParam struct {
		Embed           QueryWorkspaceParamEmbed
		IncludeArchived bool
	}

	QueryWorkspaceParamEmbed struct {
		Boards bool
		Issues bool
	}

	WorkspaceWithEmbedDatas struct {
		Workspace Workspace
		Boards    []BoardWithEmbedDatas
	}

	IdWorkspace string
)

// Implements
var _ iWorkspace = &Workspace{}

// Implements iWorkspace
type Workspace struct {
	Id      IdWorkspace
	Name    string
	Columns []Column

	Archived bool

	UserId users.IdUser
}
