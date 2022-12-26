package issues

import (
	"fickle/domain/errors"
	"strconv"

	"github.com/go-playground/colors"
)

func (p *CreateWorkspaceParam) validate() error {
	if p.Name == "" {
		return &errors.ErrValidation{
			Property:    "Name",
			Given:       &p.Name,
			Description: "cannot be empty",
		}
	}
	defaultColumnAlreadyExist := false
	for idx, c := range p.Columns {
		if c.Name == "" {
			return &errors.ErrValidation{
				Property:    "Name",
				Given:       &c.Name,
				Description: "cannot be empty",
			}
		}
		if c.Color == "" {
			return &errors.ErrValidation{
				Property:    "Color",
				Given:       &c.Color,
				Description: "cannot be empty",
			}
		}
		if c2, err := colors.Parse(c.Color); err != nil {
			return &errors.ErrValidation{
				Property:    "Columns[" + strconv.Itoa(idx) + "].Default",
				Given:       &c.Color,
				Description: "is invalid color code",
			}
		} else {
			c.Color = c2.ToHEX().String()
		}
		if c.Default {
			if defaultColumnAlreadyExist {
				return &errors.ErrValidation{
					Property:    "Columns[" + strconv.Itoa(idx) + "].Default",
					Given:       func() *string { d := "true"; return &d }(),
					Description: "is duplicated default column",
				}
			}
			defaultColumnAlreadyExist = true
		}
	}
	if !defaultColumnAlreadyExist {
		p.Columns[0].Default = true
	}
	return nil
}

// NewWorkspace implements iBoardService
func (b *boardService) NewWorkspace(f IFactory, r IRepository, p CreateWorkspaceParam) (Workspace, error) {
	// Validate parameter
	err := p.validate()
	if err != nil {
		return Workspace{}, err
	}

	// Generate id
	id, err := f.NewWorkspaceId(r)
	if err != nil {
		return Workspace{}, err
	}

	// New `workspace`
	w := Workspace{
		Id:     id,
		Name:   p.Name,
		UserId: p.UserId,
	}

	// New `columns`
	for _, c := range p.Columns {
		// Generate id
		idC, err := f.NewColumnId(r)
		if err != nil {
			return Workspace{}, err
		}
		// Append
		w.Columns = append(w.Columns, Column{
			Id:          idC,
			Name:        c.Name,
			Color:       c.Color,
			Hidden:      c.Hidden,
			Order:       c.Order,
			Default:     c.Default,
			WorkspaceId: w.Id,
			UserId:      w.UserId,
		})
	}

	return r.CreateWorkspace(w)
}
