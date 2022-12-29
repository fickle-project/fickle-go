package schedules

import (
	"fickle/domain/issues"
	"fickle/domain/users"
	"time"
)

type iLog interface {
	Update(IRepository, IQueryService, issues.IRepository, UpdateLogParam) (Log, error)
	Remove(IRepository) error
}

type (
	AddLogParam struct {
		Name    *string
		Start   time.Time
		End     time.Time
		IssueId *issues.IdIssue
		UserId  users.IdUser
	}

	UpdateLogParam struct {
		Name    **string
		Start   *time.Time
		End     *time.Time
		IssueId **issues.IdIssue
	}

	QueryLogParam struct {
		Embed       QueryLogParamEmbed
		Name        *string
		NameContain *string
		From        *time.Time
		To          *time.Time
		IssueId     *issues.IdIssue
	}

	QueryLogParamEmbed struct {
		Issue bool
	}

	LogWithEmbedDatas struct {
		Log   Log
		Issue issues.IssueWithEmbedDatas
	}

	IdLog string
)

// Implements
var _ iLog = &Log{}

// Implements iLog
type Log struct {
	Id      IdLog
	Name    *string
	IssueId *issues.IdIssue
	Start   time.Time
	End     time.Time

	UserId users.IdUser
}
