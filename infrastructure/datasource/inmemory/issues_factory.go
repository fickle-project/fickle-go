package inmemory

import (
	"fickle/domain/issues"

	"github.com/google/uuid"
)

func NewFactoryIssues() issues.IFactory {
	return &fIssues{}
}

type fIssues struct{}

// NewBoardId implements issues.IFactory
func (*fIssues) NewBoardId(r issues.IRepository) (issues.IdBoard, error) {
	return issues.IdBoard(uuid.NewString()), nil
}

// NewColumnId implements issues.IFactory
func (*fIssues) NewColumnId(r issues.IRepository) (issues.IdColumn, error) {
	return issues.IdColumn(uuid.NewString()), nil
}

// NewIssueId implements issues.IFactory
func (*fIssues) NewIssueId(r issues.IRepository) (issues.IdIssue, error) {
	return issues.IdIssue(uuid.NewString()), nil
}

// NewWorkspaceId implements issues.IFactory
func (*fIssues) NewWorkspaceId(r issues.IRepository) (issues.IdWorkspace, error) {
	return issues.IdWorkspace(uuid.NewString()), nil
}
