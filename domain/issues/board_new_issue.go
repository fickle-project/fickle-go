package issues

import "fickle/domain/errors"

func (p *NewIssueParam) validate(r IRepository) error {
	if p.ColumnId == nil && p.Name == "" {
		return &errors.ErrValidation{
			Property:    "Name",
			Given:       &p.Name,
			Description: "cannot be empty",
		}
	}
	if p.ColumnId != nil {
		if _, err := r.FindColumn(*p.ColumnId, QueryColumnParam{}); err != nil {
			return err
		}
	}
	return nil
}

// NewIssue implements iBoard
func (b *Board) NewIssue(f IFactory, r IRepository, p NewIssueParam) (Issue, error) {
	// Validate parameter
	err := p.validate(r)
	if err != nil {
		return Issue{}, err
	}

	id, err := f.NewIssueId(r)
	if err != nil {
		return Issue{}, err
	}

	// column id
	columnId := p.ColumnId
	if p.ColumnId == nil {
		// Get default column id
		isDefault := true
		c, err := r.FindColumns(QueryColumnParam{
			WorkspaceId: &b.WorkspaceId,
			Default:     &isDefault,
		})
		if err != nil {
			return Issue{}, err
		}
		if len(c) == 1 {
			columnId = &c[0].Column.Id
		}
	}

	return r.CreateIssue(Issue{
		Id:          id,
		Name:        p.Name,
		Content:     p.Content,
		BoardId:     b.Id,
		ColumnId:    columnId,
		Order:       p.Order,
		WorkspaceId: b.WorkspaceId,
		UserId:      p.UserId,
	})
}
