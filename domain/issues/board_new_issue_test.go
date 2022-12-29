package issues_test

import (
	"fickle/domain/issues"
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"reflect"
	"testing"
)

func TestBoard_NewIssue(t *testing.T) {
	type fields struct {
		Id          issues.IdBoard
		Name        string
		WorkspaceId issues.IdWorkspace
		Archived    bool
		UserId      users.IdUser
	}
	type args struct {
		f issues.IFactory
		r issues.IRepository
		p issues.NewIssueParam
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
				Name:        "fickle",
				WorkspaceId: "1",
				Archived:    false,
				UserId:      "1",
			},
			args: args{
				f: inmemory.NewFactoryIssues(),
				r: inmemory.NewRepositoryIssues(),
				p: issues.NewIssueParam{
					Name:   "issue 1",
					Order:  "1",
					UserId: "1",
				},
			},
			want: issues.Issue{
				Name:        "issue 1",
				BoardId:     "1",
				Order:       "1",
				WorkspaceId: "1",
				UserId:      "1",
			},
			wantErr: false,
		},
		{
			name: "failed to validate: 'ColumnId' not found",
			fields: fields{
				Id:          "1",
				Name:        "fickle",
				WorkspaceId: "1",
				Archived:    false,
				UserId:      "1",
			},
			args: args{
				f: inmemory.NewFactoryIssues(),
				r: inmemory.NewRepositoryIssues(),
				p: issues.NewIssueParam{
					Name:     "issue 1",
					ColumnId: func() *issues.IdColumn { var id issues.IdColumn = "-"; return &id }(),
					Order:    "1",
					UserId:   "1",
				},
			},
			want:    issues.Issue{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Name' empty",
			fields: fields{
				Id:          "1",
				Name:        "fickle",
				WorkspaceId: "1",
				Archived:    false,
				UserId:      "1",
			},
			args: args{
				f: inmemory.NewFactoryIssues(),
				r: inmemory.NewRepositoryIssues(),
				p: issues.NewIssueParam{
					Name:   "",
					Order:  "1",
					UserId: "1",
				},
			},
			want:    issues.Issue{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &issues.Board{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				WorkspaceId: tt.fields.WorkspaceId,
				Archived:    tt.fields.Archived,
				UserId:      tt.fields.UserId,
			}
			got, err := b.NewIssue(tt.args.f, tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Board.NewIssue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.Id = got.Id
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.NewIssue() = %v, want %v", got, tt.want)
			}
		})
	}
}
