package issues

func (p *RemoveBoardColumnParam) validate(r IRepository) error {
	if p.MoveIssuesTo != nil {
		if _, err := r.FindColumn(*p.MoveIssuesTo, QueryColumnParam{}); err != nil {
			return err
		}
	}
	return nil
}

// Remove implements iColumn
func (c *Column) Remove(r IRepository, p RemoveBoardColumnParam) error {
	// Validate parameter
	err := p.validate(r)
	if err != nil {
		return err
	}

	// Update issues
	// TODO: implements transaction
	tmpCId := &c.Id
	err = r.UpdateIssues(
		UpdateIssueParam{ColumnId: &p.MoveIssuesTo},
		QueryIssueParam{ColumnId: &tmpCId, WorkspaceId: &c.WorkspaceId},
	)
	if err != nil {
		return err
	}

	return r.RemoveColumn(c.Id)
}
