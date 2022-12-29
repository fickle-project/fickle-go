package issues_test

import (
	"fickle/domain/issues"
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"reflect"
	"testing"
)

func TestColumn_Update(t *testing.T) {
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
			Id:          "3",
			Name:        "In Progress",
			Color:       "#3498db",
			Hidden:      false,
			Order:       "111",
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
		p issues.UpdateBoardColumnParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    issues.Column
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
				p: issues.UpdateBoardColumnParam{
					Name:  func() *string { s := "updated"; return &s }(),
					Color: func() *string { s := "#FFF"; return &s }(),
				},
			},
			want: issues.Column{
				Id:          "1",
				Name:        "updated",
				Color:       "#fff",
				Hidden:      false,
				Order:       "1",
				Default:     true,
				WorkspaceId: "1",
				UserId:      "1",
			},
			wantErr: false,
		},
		{
			name: "failed to validate: 'Name' empty",
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
				p: issues.UpdateBoardColumnParam{
					Name:  func() *string { s := ""; return &s }(),
					Color: func() *string { s := "#FFF"; return &s }(),
				},
			},
			want:    issues.Column{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Color' invalid",
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
				p: issues.UpdateBoardColumnParam{
					Name:  func() *string { s := "updated"; return &s }(),
					Color: func() *string { s := "invalid"; return &s }(),
				},
			},
			want:    issues.Column{},
			wantErr: true,
		},
		{
			name: "failed to validate: no update",
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
				p: issues.UpdateBoardColumnParam{},
			},
			want:    issues.Column{},
			wantErr: true,
		},
		{
			name: "fail: not found",
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
				p: issues.UpdateBoardColumnParam{
					Name:  func() *string { s := "updated"; return &s }(),
					Color: func() *string { s := "#fff"; return &s }(),
				},
			},
			want:    issues.Column{},
			wantErr: true,
		},
		{
			name: "ok",
			fields: fields{
				Id:          "3",
				Name:        "In Progress",
				Color:       "#3498db",
				Hidden:      false,
				Order:       "111",
				Default:     false,
				WorkspaceId: "1",
				UserId:      "1",
			},
			args: args{
				r: r,
				p: issues.UpdateBoardColumnParam{
					Name:  func() *string { s := "Next up"; return &s }(),
					Order: func() *string { s := "11"; return &s }(),
				},
			},
			want: issues.Column{
				Id:          "3",
				Name:        "Next up",
				Color:       "#3498db",
				Hidden:      false,
				Order:       "11",
				Default:     false,
				WorkspaceId: "1",
				UserId:      "1",
			},
			wantErr: false,
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
			got, err := c.Update(tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Column.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Column.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
