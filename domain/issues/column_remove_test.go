package issues_test

import (
	"fickle/domain/issues"
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"testing"
)

func TestColumn_Remove(t *testing.T) {
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
			UserId:      "1",
		}},
		Archived: false,
		UserId:   "1",
	})

	type fields struct {
		Id          issues.IdColumn
		Name        string
		Color       string
		Hidden      bool
		Order       string
		Default     bool
		WorkspaceId issues.IdWorkspace
		UserId      users.IdUser
	}
	type args struct {
		r issues.IRepository
		p issues.RemoveBoardColumnParam
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
				Name:        "Backlog",
				Color:       "#2c3e50",
				Hidden:      false,
				Order:       "1",
				Default:     true,
				WorkspaceId: "1",
				UserId:      "1",
			},
			args: args{
				r: r,
				p: issues.RemoveBoardColumnParam{},
			},
			wantErr: false,
		},
		{
			name: "ok",
			fields: fields{
				Id:          "2",
				Name:        "Next up",
				Color:       "#f39c12",
				Hidden:      false,
				Order:       "11",
				Default:     false,
				WorkspaceId: "1",
				UserId:      "1",
			},
			args: args{
				r: r,
				p: issues.RemoveBoardColumnParam{},
			},
			wantErr: false,
		},
		{
			name: "fail: not found",
			fields: fields{
				Id:          "3",
				Name:        "Next up",
				Color:       "#f39c12",
				Hidden:      false,
				Order:       "11",
				Default:     false,
				WorkspaceId: "1",
				UserId:      "1",
			},
			args: args{
				r: r,
				p: issues.RemoveBoardColumnParam{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &issues.Column{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				Color:       tt.fields.Color,
				Hidden:      tt.fields.Hidden,
				Order:       tt.fields.Order,
				Default:     tt.fields.Default,
				WorkspaceId: tt.fields.WorkspaceId,
				UserId:      tt.fields.UserId,
			}
			if err := c.Remove(tt.args.r, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Column.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
