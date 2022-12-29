package schedules

import (
	"fickle/domain/issues"
	"fickle/domain/users"
	"time"
)

type iSchedule interface {
	Update(IRepository, IQueryService, issues.IRepository, UpdateScheduleParam) (Schedule, error)
	Remove(IRepository) error
}

type (
	AddScheduleParam struct {
		Name    *string
		Start   time.Time
		End     time.Time
		IssueId *issues.IdIssue
		UserId  users.IdUser
	}

	UpdateScheduleParam struct {
		Name    **string
		Start   *time.Time
		End     *time.Time
		IssueId **issues.IdIssue
	}

	QueryScheduleParam struct {
		Embed       QueryScheduleParamEmbed
		Name        *string
		NameContain *string
		From        *time.Time
		To          *time.Time
		IssueId     *issues.IdIssue
	}

	QueryScheduleParamEmbed struct {
		Issue bool
	}

	ScheduleWithEmbedDatas struct {
		Schedule Schedule
		Issue    issues.IssueWithEmbedDatas
	}

	IdSchedule string
)

// Implements
var _ iSchedule = &Schedule{}

// Implements iSchedule
type Schedule struct {
	Id      IdSchedule
	Name    *string
	IssueId *issues.IdIssue
	Start   time.Time
	End     time.Time

	UserId users.IdUser
}
