package issues_test

import (
	"fickle/domain/issues"
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"testing"
)

func TestBoard_Remove(t *testing.T) {
	r := inmemory.NewRepositoryIssues()
	r.CreateBoard(issues.Board{
		Id:          "1",
		Name:        "fickle",
		WorkspaceId: "1",
		Archived:    false,
		UserId:      "1",
	})
	r.CreateBoard(issues.Board{
		Id:          "2",
		Name:        "elkcif",
		WorkspaceId: "1",
		Archived:    false,
		UserId:      "1",
	})

	type fields struct {
		Id          issues.IdBoard
		Name        string
		WorkspaceId issues.IdWorkspace
		Archived    bool
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
				Name:        "fickle",
				WorkspaceId: "1",
				Archived:    false,
				UserId:      "1"},
			args: args{
				r: r,
			},
			wantErr: false,
		},
		{
			name: "ok",
			fields: fields{
				Id:          "2",
				Name:        "elkcif",
				WorkspaceId: "1",
				Archived:    false,
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
				Name:        "fickle",
				WorkspaceId: "1",
				Archived:    false,
				UserId:      "1"},
			args: args{
				r: r,
			},
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
			if err := b.Remove(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Board.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
