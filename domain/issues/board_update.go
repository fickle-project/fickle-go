package issues

import "fickle/domain/errors"

func (p *UpdateBoardParam) validate() error {
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

// Update implements iBoard
func (b *Board) Update(r IRepository, p UpdateBoardParam) (Board, error) {
	// Validate parameter
	err := p.validate()
	if err != nil {
		return Board{}, err
	}

	return r.UpdateBoard(b.Id, p)
}
