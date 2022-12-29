package issues_test

import (
	"fickle/domain/issues"
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"testing"
)

func TestWorkspace_NewBoard(t *testing.T) {
	type fields struct {
		Id       issues.IdWorkspace
		Name     string
		Columns  []issues.Column
		Archived bool
		UserId   users.IdUser
	}
	type args struct {
		f issues.IFactory
		r issues.IRepository
		p issues.AddBoardParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    issues.Board
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
					}, {
						Id:          "2",
						Name:        "Next up",
						Color:       "#f39c12",
						Hidden:      false,
						Order:       "11",
						Default:     false,
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
					}, {
						Id:          "4",
						Name:        "Pending",
						Color:       "#e67e22",
						Hidden:      false,
						Order:       "1111",
						Default:     false,
						WorkspaceId: "1",
						UserId:      "1",
					}, {
						Id:          "5",
						Name:        "Done",
						Color:       "#2ecc71",
						Hidden:      false,
						Order:       "11111",
						Default:     false,
						WorkspaceId: "1",
						UserId:      "1",
					}, {
						Id:          "6",
						Name:        "Archived",
						Color:       "#95a5a6",
						Hidden:      true,
						Order:       "111111",
						Default:     false,
						WorkspaceId: "1",
						UserId:      "1",
					}, {
						Id:          "7",
						Name:        "Canceled",
						Color:       "#e74c3c",
						Hidden:      true,
						Order:       "1111111",
						Default:     false,
						WorkspaceId: "1",
						UserId:      "1",
					}},
				Archived: false,
				UserId:   "1",
			},
			args: args{
				f: inmemory.NewFactoryIssues(),
				r: inmemory.NewRepositoryIssues(),
				p: issues.AddBoardParam{
					Name: "fickle",
				},
			},
			want: issues.Board{
				Name:        "fickle",
				WorkspaceId: "1",
				Archived:    false,
				UserId:      "1",
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
					}, {
						Id:          "2",
						Name:        "Next up",
						Color:       "#f39c12",
						Hidden:      false,
						Order:       "11",
						Default:     false,
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
					}, {
						Id:          "4",
						Name:        "Pending",
						Color:       "#e67e22",
						Hidden:      false,
						Order:       "1111",
						Default:     false,
						WorkspaceId: "1",
						UserId:      "1",
					}, {
						Id:          "5",
						Name:        "Done",
						Color:       "#2ecc71",
						Hidden:      false,
						Order:       "11111",
						Default:     false,
						WorkspaceId: "1",
						UserId:      "1",
					}, {
						Id:          "6",
						Name:        "Archived",
						Color:       "#95a5a6",
						Hidden:      true,
						Order:       "111111",
						Default:     false,
						WorkspaceId: "1",
						UserId:      "1",
					}, {
						Id:          "7",
						Name:        "Canceled",
						Color:       "#e74c3c",
						Hidden:      true,
						Order:       "1111111",
						Default:     false,
						WorkspaceId: "1",
						UserId:      "1",
					}},
				Archived: false,
				UserId:   "1",
			},
			args: args{
				f: inmemory.NewFactoryIssues(),
				r: inmemory.NewRepositoryIssues(),
				p: issues.AddBoardParam{
					Name: "",
				},
			},
			want:    issues.Board{},
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
			got, err := w.NewBoard(tt.args.f, tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workspace.NewBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if got.Name != tt.want.Name || got.WorkspaceId != tt.want.WorkspaceId || got.Archived != tt.fields.Archived || got.UserId != tt.fields.UserId {
				t.Errorf("Workspace.NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
