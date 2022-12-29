package issues

import "fickle/domain/errors"

func (p *AddBoardParam) validate() error {
	if p.Name == "" {
		return &errors.ErrValidation{
			Property:    "Name",
			Given:       &p.Name,
			Description: "cannot be empty",
		}
	}
	return nil
}

// NewBoard implements iWorkspace
func (w *Workspace) NewBoard(f IFactory, r IRepository, p AddBoardParam) (Board, error) {
	// Validate parameter
	err := p.validate()
	if err != nil {
		return Board{}, err
	}

	id, err := f.NewBoardId(r)
	if err != nil {
		return Board{}, nil
	}

	return r.CreateBoard(Board{
		Id:          id,
		Name:        p.Name,
		WorkspaceId: w.Id,
		Archived:    false,
		UserId:      w.UserId,
	})
}
