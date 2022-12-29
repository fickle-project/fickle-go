package issues

import (
	"fickle/domain/errors"

	"github.com/go-playground/colors"
)

func (p *AddBoardColumnParam) validate() error {
	if p.Name == "" {
		return &errors.ErrValidation{
			Property:    "Name",
			Given:       &p.Name,
			Description: "cannot be empty",
		}
	}
	if p.Color == "" {
		return &errors.ErrValidation{
			Property:    "Color",
			Given:       &p.Color,
			Description: "cannot be empty",
		}
	}
	if c, err := colors.Parse(p.Color); err != nil {
		return &errors.ErrValidation{
			Property:    "Color",
			Given:       &p.Color,
			Description: "is invalid color code",
		}
	} else {
		p.Color = c.ToHEX().String()
	}
	return nil
}

// NewColumn implements iWorkspace
func (w *Workspace) NewColumn(f IFactory, r IRepository, p AddBoardColumnParam) (Column, error) {
	// Validate parameter
	err := p.validate()
	if err != nil {
		return Column{}, err
	}

	// TODO: implements transaction
	if p.Default {
		// Make default column in workspace non-default
		d := true
		newDefault := false
		err = r.UpdateColumns(
			UpdateBoardColumnParam{
				Default: &newDefault,
			},
			QueryColumnParam{
				WorkspaceId: &w.Id,
				Default:     &d,
			},
		)
		if err != nil {
			return Column{}, err
		}
	}

	id, err := f.NewColumnId(r)
	if err != nil {
		return Column{}, err
	}

	return r.CreateColumn(Column{
		Id:          id,
		Name:        p.Name,
		Color:       p.Color,
		Hidden:      p.Hidden,
		Order:       p.Order,
		WorkspaceId: w.Id,
		UserId:      p.UserId,
	})
}
