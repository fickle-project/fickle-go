package issues

import (
	"fickle/domain/errors"

	"github.com/go-playground/colors"
)

func (p *UpdateBoardColumnParam) validate() error {
	if p.Name == nil && p.Color == nil && p.Hidden == nil && p.Order == nil && p.Default == nil {
		return &errors.ErrNoUpdate{}
	}
	if p.Name != nil && *p.Name == "" {
		return &errors.ErrValidation{
			Property:    "Name",
			Given:       p.Name,
			Description: "cannot be empty",
		}
	}
	if p.Color != nil {
		if *p.Color == "" {
			return &errors.ErrValidation{
				Property:    "Color",
				Given:       p.Color,
				Description: "cannot be empty",
			}
		}
		if c, err := colors.Parse(*p.Color); err != nil {
			return &errors.ErrValidation{
				Property:    "Color",
				Given:       p.Color,
				Description: "is invalid color code",
			}
		} else {
			color := c.ToHEX().String()
			p.Color = &color
		}
	}
	return nil
}

// Update implements iColumn
func (c *Column) Update(r IRepository, p UpdateBoardColumnParam) (Column, error) {
	// Validate parameter
	err := p.validate()
	if err != nil {
		return Column{}, err
	}

	// TODO: implements transaction
	if p.Default != nil && *p.Default {
		// Make default column in workspace non-default
		d := true
		newDefault := false
		err = r.UpdateColumns(
			UpdateBoardColumnParam{
				Default: &newDefault,
			},
			QueryColumnParam{
				WorkspaceId: &c.WorkspaceId,
				Default:     &d,
			},
		)
		if err != nil {
			return Column{}, err
		}
	}

	return r.UpdateColumn(c.Id, p)
}
