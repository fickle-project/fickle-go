package issues

import "fickle/domain/errors"

func (p *UpdateWorkspaceParam) validate() error {
	if p.Name == nil && p.Archived == nil {
		return &errors.ErrNoUpdate{}
	}
	if p.Name != nil && *p.Name == "" {
		return &errors.ErrValidation{
			Property:    "Name",
			Given:       p.Name,
			Description: "cannot be empty",
		}
	}
	return nil
}

// Update implements iWorkspace
func (w *Workspace) Update(r IRepository, p UpdateWorkspaceParam) (Workspace, error) {
	// Validate parameter
	err := p.validate()
	if err != nil {
		return Workspace{}, err
	}

	return r.UpdateWorkspace(w.Id, p)
}
