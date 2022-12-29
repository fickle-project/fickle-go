package issues_test

import (
	"fickle/domain/issues"
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"testing"
)

func TestWorkspace_Remove(t *testing.T) {
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

	type fields struct {
		Id       issues.IdWorkspace
		Name     string
		Columns  []issues.Column
		Archived bool
		UserId   users.IdUser
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
			},
			wantErr: false,
		},
		{
			name: "fail: not found",
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
			},
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
			if err := w.Remove(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Workspace.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
