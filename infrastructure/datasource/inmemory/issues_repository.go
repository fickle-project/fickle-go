package inmemory

import (
	"fickle/domain/errors"
	"fickle/domain/issues"
	"strings"
)

func NewRepositoryIssues() issues.IRepository {
	return &rIssues{}
}

type rIssues struct {
	workspaces []issues.Workspace
	columns    []issues.Column
	boards     []issues.Board
	issues     []issues.Issue
}

// CreateBoard implements issues.IRepository
func (r *rIssues) CreateBoard(b issues.Board) (issues.Board, error) {
	r.boards = append(r.boards, b)
	return b, nil
}

// CreateColumn implements issues.IRepository
func (r *rIssues) CreateColumn(c issues.Column) (issues.Column, error) {
	r.columns = append(r.columns, c)
	return c, nil
}

// CreateIssue implements issues.IRepository
func (r *rIssues) CreateIssue(i issues.Issue) (issues.Issue, error) {
	r.issues = append(r.issues, i)
	return i, nil
}

// CreateWorkspace implements issues.IRepository
func (r *rIssues) CreateWorkspace(w issues.Workspace) (issues.Workspace, error) {
	r.workspaces = append(r.workspaces, issues.Workspace{
		Id:       w.Id,
		Name:     w.Name,
		Archived: w.Archived,
		UserId:   w.UserId,
	})
	for _, c := range w.Columns {
		_, err := r.CreateColumn(c)
		if err != nil {
			return issues.Workspace{}, err
		}
	}
	return w, nil
}

// FindBoard implements issues.IRepository
func (r *rIssues) FindBoard(id issues.IdBoard, q issues.QueryBoardParam) (issues.BoardWithEmbedDatas, error) {
	bb := filter(r.boards, func(b issues.Board) bool { return b.Id == id })

	if len(bb) == 0 {
		return issues.BoardWithEmbedDatas{}, &errors.ErrNotFound{Object: "Board", Id: string(id)}
	}

	embed := issues.BoardWithEmbedDatas{Board: bb[0]}
	if q.Embed.Workspace {
		w, err := r.FindWorkspace(bb[0].WorkspaceId, issues.QueryWorkspaceParam{})
		if err != nil {
			return issues.BoardWithEmbedDatas{}, err
		}
		embed.Workspace = w.Workspace
	}
	if q.Embed.Issues {
		c, err := r.FindColumns(issues.QueryColumnParam{
			Embed:       issues.QueryColumnParamEmbed{Issues: true},
			BoardId:     &embed.Board.Id,
			WorkspaceId: &embed.Board.WorkspaceId,
		})
		if err != nil {
			return issues.BoardWithEmbedDatas{}, err
		}
		embed.Columns = c
	}

	return embed, nil
}

// FindBoards implements issues.IRepository
func (r *rIssues) FindBoards(q issues.QueryBoardParam) ([]issues.BoardWithEmbedDatas, error) {
	bb := filter(r.boards, func(b issues.Board) bool {
		selected := true
		if q.Name != nil {
			selected = selected && b.Name == *q.Name
		}
		if q.NameContain != nil {
			selected = selected && strings.Contains(b.Name, *q.NameContain)
		}
		if q.WorkspaceId != nil {
			selected = selected && b.WorkspaceId == *q.WorkspaceId
		}
		selected = selected && (!b.Archived || q.IncludeArchived)
		return selected
	})

	embeds := []issues.BoardWithEmbedDatas{}
	for _, b := range bb {
		embed := issues.BoardWithEmbedDatas{Board: b}
		if q.Embed.Workspace {
			w, err := r.FindWorkspace(b.WorkspaceId, issues.QueryWorkspaceParam{})
			if err != nil {
				return nil, err
			}
			embed.Workspace = w.Workspace
		}

		if q.Embed.Issues {
			cc, err := r.FindColumns(issues.QueryColumnParam{
				Embed:       issues.QueryColumnParamEmbed{Issues: true},
				BoardId:     &embed.Board.Id,
				WorkspaceId: &embed.Board.WorkspaceId,
			})
			if err != nil {
				return nil, err
			}
			embed.Columns = cc
		}
		embeds = append(embeds, embed)
	}

	return embeds, nil
}

// FindColumn implements issues.IRepository
func (r *rIssues) FindColumn(id issues.IdColumn, q issues.QueryColumnParam) (issues.ColumnWithEmbedDatas, error) {
	cc := filter(r.columns, func(c issues.Column) bool { return c.Id == id })

	if len(cc) == 0 {
		return issues.ColumnWithEmbedDatas{}, &errors.ErrNotFound{Object: "Column", Id: string(id)}
	}

	embed := issues.ColumnWithEmbedDatas{Column: cc[0]}
	if q.Embed.Workspace {
		w, err := r.FindWorkspace(embed.Column.WorkspaceId, issues.QueryWorkspaceParam{})
		if err != nil {
			return issues.ColumnWithEmbedDatas{}, err
		}
		embed.Workspace = w.Workspace
	}
	if q.Embed.Issues {
		tmpCId := &embed.Column.Id
		ii, err := r.FindIssues(issues.QueryIssueParam{
			Embed:       issues.QueryIssueParamEmbed{Board: q.Embed.Board},
			BoardId:     q.BoardId,
			ColumnId:    &tmpCId,
			WorkspaceId: &embed.Column.WorkspaceId,
		})
		if err != nil {
			return issues.ColumnWithEmbedDatas{}, err
		}
		embed.Issues = ii
	}

	return embed, nil
}

// FindColumns implements issues.IRepository
func (r *rIssues) FindColumns(q issues.QueryColumnParam) ([]issues.ColumnWithEmbedDatas, error) {
	cc := filter(r.columns, func(c issues.Column) bool {
		selected := true
		if q.WorkspaceId != nil {
			selected = selected && c.WorkspaceId == *q.WorkspaceId
		}
		if q.Default != nil {
			selected = selected && c.Default == *q.Default
		}
		return selected
	})

	embeds := []issues.ColumnWithEmbedDatas{}
	for _, c := range cc {
		embed := issues.ColumnWithEmbedDatas{Column: c}
		if q.Embed.Workspace {
			w, err := r.FindWorkspace(embed.Column.WorkspaceId, issues.QueryWorkspaceParam{})
			if err != nil {
				return nil, err
			}
			embed.Workspace = w.Workspace
		}
		if q.Embed.Issues {
			tmpCId := &embed.Column.Id
			ii, err := r.FindIssues(issues.QueryIssueParam{
				Embed:       issues.QueryIssueParamEmbed{Board: q.Embed.Board},
				BoardId:     q.BoardId,
				ColumnId:    &tmpCId,
				WorkspaceId: &embed.Column.WorkspaceId,
			})
			if err != nil {
				return nil, err
			}
			embed.Issues = ii
		}
		embeds = append(embeds, embed)
	}

	return embeds, nil
}

// FindIssue implements issues.IRepository
func (r *rIssues) FindIssue(id issues.IdIssue, q issues.QueryIssueParam) (issues.IssueWithEmbedDatas, error) {
	ii := filter(r.issues, func(i issues.Issue) bool { return i.Id == id })

	if len(ii) == 0 {
		return issues.IssueWithEmbedDatas{}, &errors.ErrNotFound{Object: "Issue", Id: string(id)}
	}

	embed := issues.IssueWithEmbedDatas{Issue: ii[0]}
	if q.Embed.Column && embed.Issue.ColumnId != nil {
		c, err := r.FindColumn(*embed.Issue.ColumnId, issues.QueryColumnParam{})
		if err != nil {
			return issues.IssueWithEmbedDatas{}, err
		}
		embed.Column = c.Column
	}
	if q.Embed.Board {
		b, err := r.FindBoard(embed.Issue.BoardId, issues.QueryBoardParam{})
		if err != nil {
			return issues.IssueWithEmbedDatas{}, err
		}
		embed.Board = b.Board
	}
	if q.Embed.Workspace {
		w, err := r.FindWorkspace(embed.Issue.WorkspaceId, issues.QueryWorkspaceParam{})
		if err != nil {
			return issues.IssueWithEmbedDatas{}, err
		}
		embed.Workspace = w.Workspace
	}

	return embed, nil
}

// FindIssues implements issues.IRepository
func (r *rIssues) FindIssues(q issues.QueryIssueParam) ([]issues.IssueWithEmbedDatas, error) {
	ii := filter(r.issues, func(i issues.Issue) bool {
		selected := true
		if q.Name != nil {
			selected = selected && i.Name == *q.Name
		}
		if q.NameContain != nil {
			selected = selected && strings.Contains(i.Name, *q.NameContain)
		}
		if q.BoardId != nil {
			selected = selected && i.BoardId == *q.BoardId
		}
		if q.ColumnId != nil {
			if *q.ColumnId != nil {
				selected = selected && i.ColumnId != nil && *i.ColumnId == **q.ColumnId
			} else {
				selected = selected && i.ColumnId == nil
			}
		}
		if q.WorkspaceId != nil {
			selected = selected && i.WorkspaceId == *q.WorkspaceId
		}
		return selected
	})

	embeds := []issues.IssueWithEmbedDatas{}
	for _, i := range ii {
		embed := issues.IssueWithEmbedDatas{Issue: i}
		if q.Embed.Column && embed.Issue.ColumnId != nil {
			c, err := r.FindColumn(*embed.Issue.ColumnId, issues.QueryColumnParam{})
			if err != nil {
				return nil, err
			}
			embed.Column = c.Column
		}
		if q.Embed.Board {
			b, err := r.FindBoard(embed.Issue.BoardId, issues.QueryBoardParam{})
			if err != nil {
				return nil, err
			}
			embed.Board = b.Board
		}
		if q.Embed.Workspace {
			w, err := r.FindWorkspace(embed.Issue.WorkspaceId, issues.QueryWorkspaceParam{})
			if err != nil {
				return nil, err
			}
			embed.Workspace = w.Workspace
		}
		embeds = append(embeds, embed)
	}

	return embeds, nil
}

// FindWorkspace implements issues.IRepository
func (r *rIssues) FindWorkspace(id issues.IdWorkspace, q issues.QueryWorkspaceParam) (issues.WorkspaceWithEmbedDatas, error) {
	ww := filter(r.workspaces, func(w issues.Workspace) bool { return w.Id == id })

	if len(ww) == 0 {
		return issues.WorkspaceWithEmbedDatas{}, &errors.ErrNotFound{Object: "Workspace", Id: string(id)}
	}

	ce, err := r.FindColumns(issues.QueryColumnParam{WorkspaceId: &ww[0].Id})
	if err != nil {
		return issues.WorkspaceWithEmbedDatas{}, err
	}
	for _, c := range ce {
		ww[0].Columns = append(ww[0].Columns, c.Column)
	}

	embed := issues.WorkspaceWithEmbedDatas{Workspace: ww[0]}
	if q.Embed.Boards {
		bb, err := r.FindBoards(issues.QueryBoardParam{
			Embed:           issues.QueryBoardParamEmbed{Issues: q.Embed.Issues},
			WorkspaceId:     &embed.Workspace.Id,
			IncludeArchived: q.IncludeArchived,
		})
		if err != nil {
			return issues.WorkspaceWithEmbedDatas{}, err
		}
		embed.Boards = bb
	}

	return embed, nil
}

// FindWorkspaces implements issues.IRepository
func (r *rIssues) FindWorkspaces(q issues.QueryWorkspaceParam) ([]issues.WorkspaceWithEmbedDatas, error) {
	ww := filter(r.workspaces, func(w issues.Workspace) bool {
		selected := true
		selected = selected && (!w.Archived || q.IncludeArchived)
		return selected
	})

	for i, w := range ww {
		ce, err := r.FindColumns(issues.QueryColumnParam{WorkspaceId: &w.Id})
		if err != nil {
			return nil, err
		}
		cc := []issues.Column{}
		for _, c := range ce {
			cc = append(cc, c.Column)
		}
		ww[i].Columns = cc
	}

	embeds := []issues.WorkspaceWithEmbedDatas{}
	for _, w := range ww {
		embed := issues.WorkspaceWithEmbedDatas{Workspace: w}
		if q.Embed.Boards {
			bb, err := r.FindBoards(issues.QueryBoardParam{
				Embed:           issues.QueryBoardParamEmbed{Issues: q.Embed.Issues},
				WorkspaceId:     &embed.Workspace.Id,
				IncludeArchived: q.IncludeArchived,
			})
			if err != nil {
				return nil, err
			}
			embed.Boards = bb
		}
		embeds = append(embeds, embed)
	}

	return embeds, nil
}

// RemoveBoard implements issues.IRepository
func (r *rIssues) RemoveBoard(id issues.IdBoard) error {
	l := len(r.boards)
	r.boards = filter(r.boards, func(b issues.Board) bool { return b.Id != id })
	if len(r.boards) == l {
		return &errors.ErrNotFound{Object: "Board", Id: string(id)}
	}
	return nil
}

// RemoveColumn implements issues.IRepository
func (r *rIssues) RemoveColumn(id issues.IdColumn) error {
	l := len(r.columns)
	r.columns = filter(r.columns, func(c issues.Column) bool { return c.Id != id })
	if len(r.columns) == l {
		return &errors.ErrNotFound{Object: "Column", Id: string(id)}
	}
	return nil
}

// RemoveIssue implements issues.IRepository
func (r *rIssues) RemoveIssue(id issues.IdIssue) error {
	l := len(r.issues)
	r.issues = filter(r.issues, func(i issues.Issue) bool { return i.Id != id })
	if len(r.issues) == l {
		return &errors.ErrNotFound{Object: "Issue", Id: string(id)}
	}
	return nil
}

// RemoveWorkspcae implements issues.IRepository
func (r *rIssues) RemoveWorkspcae(id issues.IdWorkspace) error {
	l := len(r.workspaces)
	r.workspaces = filter(r.workspaces, func(w issues.Workspace) bool { return w.Id != id })
	if len(r.workspaces) == l {
		return &errors.ErrNotFound{Object: "Workspcae", Id: string(id)}
	}
	return nil
}

// UpdateBoard implements issues.IRepository
func (r *rIssues) UpdateBoard(id issues.IdBoard, p issues.UpdateBoardParam) (issues.Board, error) {
	be, err := r.FindBoard(id, issues.QueryBoardParam{})
	if err != nil {
		return issues.Board{}, err
	}

	err = r.RemoveBoard(id)
	if err != nil {
		return issues.Board{}, err
	}

	b := be.Board
	if p.Name != nil {
		b.Name = *p.Name
	}
	if p.Archived != nil {
		b.Archived = *p.Archived
	}
	return r.CreateBoard(b)
}

// UpdateColumn implements issues.IRepository
func (r *rIssues) UpdateColumn(id issues.IdColumn, p issues.UpdateBoardColumnParam) (issues.Column, error) {
	ce, err := r.FindColumn(id, issues.QueryColumnParam{})
	if err != nil {
		return issues.Column{}, err
	}

	err = r.RemoveColumn(id)
	if err != nil {
		return issues.Column{}, err
	}

	c := ce.Column
	if p.Name != nil {
		c.Name = *p.Name
	}
	if p.Color != nil {
		c.Color = *p.Color
	}
	if p.Hidden != nil {
		c.Hidden = *p.Hidden
	}
	if p.Order != nil {
		c.Order = *p.Order
	}
	if p.Default != nil {
		c.Default = *p.Default
	}
	return r.CreateColumn(c)
}

// UpdateColumns implements issues.IRepository
func (r *rIssues) UpdateColumns(p issues.UpdateBoardColumnParam, q issues.QueryColumnParam) error {
	cc, err := r.FindColumns(q)
	if err != nil {
		return err
	}

	for _, ce := range cc {
		c := ce.Column
		if p.Name != nil {
			c.Name = *p.Name
		}
		if p.Color != nil {
			c.Color = *p.Color
		}
		if p.Hidden != nil {
			c.Hidden = *p.Hidden
		}
		if p.Order != nil {
			c.Order = *p.Order
		}
		if p.Default != nil {
			c.Default = *p.Default
		}
		err := r.RemoveColumn(c.Id)
		if err != nil {
			return err
		}
		_, err = r.CreateColumn(c)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateIssue implements issues.IRepository
func (r *rIssues) UpdateIssue(id issues.IdIssue, p issues.UpdateIssueParam) (issues.Issue, error) {
	ie, err := r.FindIssue(id, issues.QueryIssueParam{})
	if err != nil {
		return issues.Issue{}, err
	}

	err = r.RemoveIssue(id)
	if err != nil {
		return issues.Issue{}, err
	}

	i := ie.Issue
	if p.Name != nil {
		i.Name = *p.Name
	}
	if p.Content != nil {
		i.Content = *p.Content
	}
	if p.BoardId != nil {
		i.BoardId = *p.BoardId
	}
	if p.ColumnId != nil {
		i.ColumnId = *p.ColumnId
	}
	if p.Order != nil {
		i.Order = *p.Order
	}
	return r.CreateIssue(i)
}

// UpdateIssues implements issues.IRepository
func (r *rIssues) UpdateIssues(p issues.UpdateIssueParam, q issues.QueryIssueParam) error {
	ii, err := r.FindIssues(q)
	if err != nil {
		return err
	}

	for _, ie := range ii {
		i := ie.Issue
		if p.Name != nil {
			i.Name = *p.Name
		}
		if p.Content != nil {
			i.Content = *p.Content
		}
		if p.BoardId != nil {
			i.BoardId = *p.BoardId
		}
		if p.ColumnId != nil {
			i.ColumnId = *p.ColumnId
		}
		if p.Order != nil {
			i.Order = *p.Order
		}
		err = r.RemoveIssue(i.Id)
		if err != nil {
			return err
		}
		_, err = r.CreateIssue(i)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateWorkspace implements issues.IRepository
func (r *rIssues) UpdateWorkspace(id issues.IdWorkspace, p issues.UpdateWorkspaceParam) (issues.Workspace, error) {
	we, err := r.FindWorkspace(id, issues.QueryWorkspaceParam{})
	if err != nil {
		return issues.Workspace{}, err
	}

	err = r.RemoveWorkspcae(id)
	if err != nil {
		return issues.Workspace{}, err
	}

	w := we.Workspace
	if p.Name != nil {
		w.Name = *p.Name
	}
	if p.Archived != nil {
		w.Archived = *p.Archived
	}
	return r.CreateWorkspace(w)
}
