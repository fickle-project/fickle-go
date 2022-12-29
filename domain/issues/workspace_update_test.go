package issues_test

import (
	"fickle/domain/issues"
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"reflect"
	"testing"
)

func TestWorkspace_Update(t *testing.T) {
	r := inmemory.NewRepositoryIssues()
	r.CreateWorkspace(issues.Workspace{
		Id:   "1",
		Name: "general",
		Columns: []issues.Column{
			{
				Id:          "1",
				Name:        "Backlog",
				Color:       "#2c3e50",
				Hidden:      false,
				Order:       "1",
				Default:     true,
				WorkspaceId: "1",
				UserId:      "1",
			},
		},
		Archived: false,
		UserId:   "1",
	})
	r.CreateWorkspace(issues.Workspace{
		Id:   "2",
		Name: "archive",
		Columns: []issues.Column{
			{
				Id:          "2",
				Name:        "Backlog",
				Color:       "#2c3e50",
				Hidden:      false,
				Order:       "1",
				Default:     true,
				WorkspaceId: "2",
				UserId:      "1",
			},
		},
		Archived: false,
		UserId:   "1",
	})

	type fields struct {
		Id       issues.IdWorkspace
		Name     string
		Columns  []issues.Column
		Archived bool
		UserId   users.IdUser
	}
	type args struct {
		r issues.IRepository
		p issues.UpdateWorkspaceParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    issues.Workspace
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Id:   "1",
				Name: "general",
				Columns: []issues.Column{
					{
						Id:          "1",
						Name:        "Backlog",
						Color:       "#2c3e50",
						Hidden:      false,
						Order:       "1",
						Default:     true,
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
				Archived: false,
				UserId:   "1",
			},
			args: args{
				r: r,
				p: issues.UpdateWorkspaceParam{
					Name:     func() *string { s := "updated"; return &s }(),
					Archived: nil,
				},
			},
			want: issues.Workspace{
				Name: "updated",
				Columns: []issues.Column{
					{
						Id:          "1",
						Name:        "Backlog",
						Color:       "#2c3e50",
						Hidden:      false,
						Order:       "1",
						Default:     true,
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
				Archived: false,
				UserId:   "1",
			},
			wantErr: false,
		},
		{
			name: "failed to validate: 'Name' empty",
			fields: fields{
				Id:   "1",
				Name: "general",
				Columns: []issues.Column{
					{
						Id:          "1",
						Name:        "Backlog",
						Color:       "#2c3e50",
						Hidden:      false,
						Order:       "1",
						Default:     true,
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
				Archived: false,
				UserId:   "1",
			},
			args: args{
				r: r,
				p: issues.UpdateWorkspaceParam{
					Name:     func() *string { s := ""; return &s }(),
					Archived: nil,
				},
			},
			want:    issues.Workspace{},
			wantErr: true,
		},
		{
			name: "failed to validate: no update",
			fields: fields{
				Id:   "1",
				Name: "general",
				Columns: []issues.Column{
					{
						Id:          "1",
						Name:        "Backlog",
						Color:       "#2c3e50",
						Hidden:      false,
						Order:       "1",
						Default:     true,
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
				Archived: false,
				UserId:   "1",
			},
			args: args{
				r: r,
				p: issues.UpdateWorkspaceParam{},
			},
			want:    issues.Workspace{},
			wantErr: true,
		},
		{
			name: "ok",
			fields: fields{
				Id:   "2",
				Name: "archive",
				Columns: []issues.Column{
					{
						Id:          "2",
						Name:        "Backlog",
						Color:       "#2c3e50",
						Hidden:      false,
						Order:       "1",
						Default:     true,
						WorkspaceId: "2",
						UserId:      "1",
					},
				},
				Archived: false,
				UserId:   "1",
			},
			args: args{
				r: r,
				p: issues.UpdateWorkspaceParam{
					Name: func() *string { s := "updated"; return &s }(),
				},
			},
			want: issues.Workspace{
				Name: "updated",
				Columns: []issues.Column{
					{
						Id:          "2",
						Name:        "Backlog",
						Color:       "#2c3e50",
						Hidden:      false,
						Order:       "1",
						Default:     true,
						WorkspaceId: "2",
						UserId:      "1",
					},
				},
				Archived: false,
				UserId:   "1",
			},
			wantErr: false,
		},
		{
			name: "fail: not found",
			fields: fields{
				Id:   "0",
				Name: "general",
				Columns: []issues.Column{
					{
						Id:          "1",
						Name:        "Backlog",
						Color:       "#2c3e50",
						Hidden:      false,
						Order:       "1",
						Default:     true,
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
				Archived: false,
				UserId:   "1",
			},
			args: args{
				r: r,
				p: issues.UpdateWorkspaceParam{
					Name:     func() *string { s := "updated"; return &s }(),
					Archived: nil,
				},
			},
			want:    issues.Workspace{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &issues.Workspace{
				Id:       tt.fields.Id,
				Name:     tt.fields.Name,
				Columns:  tt.fields.Columns,
				Archived: tt.fields.Archived,
				UserId:   tt.fields.UserId,
			}
			got, err := w.Update(tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workspace.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.Id = got.Id
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workspace.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
