package issues_test

import (
	"fickle/domain/issues"
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"reflect"
	"testing"
)

func TestIssue_Update(t *testing.T) {
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
		}, {
			Id:          "2",
			Name:        "Next up",
			Color:       "#f39c12",
			Hidden:      false,
			Order:       "11",
			Default:     false,
			WorkspaceId: "1",
			UserId:      "1"}},
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
		p issues.UpdateIssueParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    issues.Issue
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Id:          "1",
				Name:        "issue 1",
				Content:     nil,
				BoardId:     "1",
				ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
				Order:       "1",
				WorkspaceId: "1",
				UserId:      "1",
			},
			args: args{
				r: r,
				p: issues.UpdateIssueParam{
					Name:     func() *string { s := "updated"; return &s }(),
					Content:  func() **string { s := "content..."; tmp := &s; return &tmp }(),
					ColumnId: func() **issues.IdColumn { var id issues.IdColumn = "2"; tmp := &id; return &tmp }(),
					Order:    func() *string { s := "11"; return &s }(),
				},
			},
			want: issues.Issue{
				Id:          "1",
				Name:        "updated",
				Content:     func() *string { s := "content..."; return &s }(),
				BoardId:     "1",
				ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "2"; return &id }(),
				Order:       "11",
				WorkspaceId: "1",
				UserId:      "1",
			},
			wantErr: false,
		},
		{
			name: "ok",
			fields: fields{
				Id:          "2",
				Name:        "issue 2",
				Content:     func() *string { s := "content..."; return &s }(),
				BoardId:     "1",
				ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
				Order:       "1",
				WorkspaceId: "1",
				UserId:      "1",
			},
			args: args{
				r: r,
				p: issues.UpdateIssueParam{
					Content: func() **string { var s *string = nil; return &s }(),
				},
			},
			want: issues.Issue{
				Id:          "2",
				Name:        "issue 2",
				Content:     nil,
				BoardId:     "1",
				ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
				Order:       "1",
				WorkspaceId: "1",
				UserId:      "1",
			},
			wantErr: false,
		},
		{
			name: "failed to validate: 'Name' empty",
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
				p: issues.UpdateIssueParam{
					Name: func() *string { s := ""; return &s }(),
				},
			},
			want:    issues.Issue{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'ColumnId' not found",
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
				p: issues.UpdateIssueParam{
					ColumnId: func() **issues.IdColumn { var id issues.IdColumn = "3"; tmp := &id; return &tmp }(),
				},
			},
			want:    issues.Issue{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'BoardId' not found",
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
				p: issues.UpdateIssueParam{
					BoardId: func() *issues.IdBoard { var id issues.IdBoard = "2"; return &id }(),
				},
			},
			want:    issues.Issue{},
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
			got, err := i.Update(tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Issue.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Issue.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
