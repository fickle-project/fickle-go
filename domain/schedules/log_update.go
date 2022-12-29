package schedules

import (
	"fickle/domain/errors"
	"fickle/domain/issues"
)

func (p *UpdateLogParam) validate(q IQueryService, ri issues.IRepository, id IdLog) error {
	if p.Name == nil && p.Start == nil && p.End == nil && p.IssueId == nil {
		return &errors.ErrNoUpdate{}
	}
	var l *Log
	if p.Name != nil && (*p.Name == nil || **p.Name == "") {
		if p.IssueId != nil && *p.IssueId == nil {
			return &errors.ErrValidation{
				Property:    "Name",
				Given:       *p.Name,
				Description: "cannot be empty if no issue",
			}
		}
		if p.IssueId == nil {
			l2, err := q.FindLog(id, QueryLogParam{})
			if err != nil {
				return err
			}
			l = &l2.Log
			if l.IssueId == nil {
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
		if l == nil {
			l2, err := q.FindLog(id, QueryLogParam{})
			if err != nil {
				return err
			}
			l = &l2.Log
		}
		start := p.Start
		end := p.End
		if start == nil {
			start = &l.Start
		}
		if end == nil {
			end = &l.End
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

// Update implements iLog
func (l *Log) Update(r IRepository, q IQueryService, ri issues.IRepository, p UpdateLogParam) (Log, error) {
	// Validate parameter
	err := p.validate(q, ri, l.Id)
	if err != nil {
		return Log{}, err
	}

	return r.UpdateLog(l.Id, p)
}
