package issues_test

import (
	"fickle/domain/issues"
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"testing"
)

func TestIssue_Remove(t *testing.T) {
	r := inmemory.NewRepositoryIssues()
	r.CreateWorkspace(issues.Workspace{
		Id:   "1",
		Name: "general",
		Columns: []issues.Column{{
			Id:          "1",
			Name:        "Backlog",
			Color:       "#2c3e50",
			Hidden:      false,
			Order:       "1",
			Default:     true,
			WorkspaceId: "1",
			UserId:      "1",
		}},
		Archived: false,
		UserId:   "1",
	})
	r.CreateBoard(issues.Board{
		Id:          "1",
		Name:        "fickle",
		WorkspaceId: "1",
		Archived:    false,
		UserId:      "1",
	})
	r.CreateIssue(issues.Issue{
		Id:          "1",
		Name:        "issue 1",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})
	r.CreateIssue(issues.Issue{
		Id:          "2",
		Name:        "issue 2",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})

	type fields struct {
		Id          issues.IdIssue
		Name        string
		Content     *string
		BoardId     issues.IdBoard
		ColumnId    *issues.IdColumn
		Order       string
		WorkspaceId issues.IdWorkspace
		UserId      users.IdUser
	}
	type args struct {
		r issues.IRepository
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Id:          "1",
				Name:        "issue 1",
				BoardId:     "1",
				ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
				Order:       "1",
				WorkspaceId: "1",
				UserId:      "1",
			},
			args: args{
				r: r,
			},
			wantErr: false,
		},
		{
			name: "ok",
			fields: fields{
				Id:          "2",
				Name:        "issue 2",
				BoardId:     "1",
				ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
				Order:       "1",
				WorkspaceId: "1",
				UserId:      "1",
			},
			args: args{
				r: r,
			},
			wantErr: false,
		},
		{
			name: "fail: not found",
			fields: fields{
				Id:          "1",
				Name:        "issue 1",
				BoardId:     "1",
				ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
				Order:       "1",
				WorkspaceId: "1",
				UserId:      "1",
			},
			args: args{
				r: r,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &issues.Issue{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				Content:     tt.fields.Content,
				BoardId:     tt.fields.BoardId,
				ColumnId:    tt.fields.ColumnId,
				Order:       tt.fields.Order,
				WorkspaceId: tt.fields.WorkspaceId,
				UserId:      tt.fields.UserId,
			}
			if err := i.Remove(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Issue.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
