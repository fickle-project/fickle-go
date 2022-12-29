package schedules

import (
	"fickle/domain/errors"
	"fickle/domain/issues"
)

func (p *AddScheduleParam) validate(ri issues.IRepository) error {
	if (p.Name == nil || *p.Name == "") && p.IssueId == nil {
		return &errors.ErrValidation{
			Property:    "Name",
			Given:       p.Name,
			Description: "cannot be empty if no issue",
		}
	}
	if p.Name != nil && *p.Name == "" {
		p.Name = nil
	}
	if !p.End.After(p.Start) {
		return &errors.ErrValidation{
			Property:    "End",
			Given:       func() *string { s := p.End.String(); return &s }(),
			Description: "must be after \"Start\"",
		}
	}
	if p.IssueId != nil {
		_, err := ri.FindIssue(*p.IssueId, issues.QueryIssueParam{})
		if err != nil {
			return err
		}
	}
	return nil
}

// NewSchedule implements iTimeTableService
func (t *timeTableService) NewSchedule(f IFactory, r IRepository, ri issues.IRepository, p AddScheduleParam) (Schedule, error) {
	// Validate parameter
	err := p.validate(ri)
	if err != nil {
		return Schedule{}, err
	}

	id, err := f.NewScheduleId(r)
	if err != nil {
		return Schedule{}, err
	}

	s := Schedule{
		Id:      id,
		Name:    p.Name,
		IssueId: p.IssueId,
		Start:   p.Start,
		End:     p.End,
		UserId:  p.UserId,
	}

	return r.AddSchedule(s)
}
