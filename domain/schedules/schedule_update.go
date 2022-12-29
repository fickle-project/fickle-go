package schedules

import (
	"fickle/domain/errors"
	"fickle/domain/issues"
)

func (p *UpdateScheduleParam) validate(q IQueryService, ri issues.IRepository, id IdSchedule) error {
	if p.Name == nil && p.Start == nil && p.End == nil && p.IssueId == nil {
		return &errors.ErrNoUpdate{}
	}
	var s *Schedule
	if p.Name != nil && (*p.Name == nil || **p.Name == "") {
		if p.IssueId != nil && *p.IssueId == nil {
			return &errors.ErrValidation{
				Property:    "Name",
				Given:       *p.Name,
				Description: "cannot be empty if no issue",
			}
		}
		if p.IssueId == nil {
			s2, err := q.FindSchedule(id, QueryScheduleParam{})
			if err != nil {
				return err
			}
			s = &s2.Schedule
			if s.IssueId == nil {
				return &errors.ErrValidation{
					Property:    "Name",
					Given:       *p.Name,
					Description: "cannot be empty if no issue",
				}
			}
		}
	}
	if p.Name != nil && *p.Name != nil && **p.Name == "" {
		*p.Name = nil
	}
	if p.Start != nil && p.End != nil {
		if !(*p.End).After(*p.Start) {
			return &errors.ErrValidation{
				Property:    "End",
				Given:       func() *string { s := p.End.String(); return &s }(),
				Description: "must be after \"Start\"",
			}
		}
	}
	if p.Start != nil && p.End == nil || p.Start == nil && p.End != nil {
		if s == nil {
			s2, err := q.FindSchedule(id, QueryScheduleParam{})
			if err != nil {
				return err
			}
			s = &s2.Schedule
		}
		start := p.Start
		end := p.End
		if start == nil {
			start = &s.Start
		}
		if end == nil {
			end = &s.End
		}
		if !(*end).After(*start) {
			return &errors.ErrValidation{
				Property:    "End",
				Given:       func() *string { s := p.End.String(); return &s }(),
				Description: "must be after \"Start\"",
			}
		}
	}
	if p.IssueId != nil && *p.IssueId != nil {
		_, err := ri.FindIssue(**p.IssueId, issues.QueryIssueParam{})
		if err != nil {
			return err
		}
	}
	return nil
}

// Update implements iSchedule
func (s *Schedule) Update(r IRepository, q IQueryService, ri issues.IRepository, p UpdateScheduleParam) (Schedule, error) {
	// Validate parameters
	err := p.validate(q, ri, s.Id)
	if err != nil {
		return Schedule{}, err
	}

	return r.UpdateSchedule(s.Id, p)
}
