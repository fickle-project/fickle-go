package issues

import "fickle/domain/errors"

func (p *UpdateIssueParam) validate(r IRepository) error {
	if p.Name == nil && p.Content == nil && p.BoardId == nil && p.ColumnId == nil && p.Order == nil {
		return &errors.ErrNoUpdate{}
	}
	if p.Name != nil && *p.Name == "" {
		return &errors.ErrValidation{
			Property:    "Name",
			Given:       p.Name,
			Description: "cannot be empty",
		}
	}
	if p.BoardId != nil {
		if _, err := r.FindBoard(*p.BoardId, QueryBoardParam{}); err != nil {
			return err
		}
	}
	if p.ColumnId != nil && *p.ColumnId != nil {
		if _, err := r.FindColumn(**p.ColumnId, QueryColumnParam{}); err != nil {
			return err
		}
	}
	return nil
}

// Update implements iIssue
func (i *Issue) Update(r IRepository, p UpdateIssueParam) (Issue, error) {
	// Validate parameter
	err := p.validate(r)
	if err != nil {
		return Issue{}, err
	}

	return r.UpdateIssue(i.Id, p)
}
