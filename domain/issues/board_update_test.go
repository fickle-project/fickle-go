package issues_test

import (
	"fickle/domain/issues"
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"reflect"
	"testing"
)

func TestBoard_Update(t *testing.T) {
	r := inmemory.NewRepositoryIssues()
	r.CreateBoard(issues.Board{
		Id:          "1",
		Name:        "fickle",
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
		p issues.UpdateBoardParam
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
				Id:          "1",
				Name:        "fickle",
				WorkspaceId: "1",
				Archived:    false,
				UserId:      "1",
			},
			args: args{
				r: r,
				p: issues.UpdateBoardParam{
					Name: func() *string { s := "updated"; return &s }(),
				},
			},
			want: issues.Board{
				Id:          "1",
				Name:        "updated",
				WorkspaceId: "1",
				Archived:    false,
				UserId:      "1",
			},
			wantErr: false,
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
				r: r,
				p: issues.UpdateBoardParam{
					Name: func() *string { s := ""; return &s }(),
				},
			},
			want:    issues.Board{},
			wantErr: true,
		},
		{
			name: "failed to validate: no update",
			fields: fields{
				Id:          "2",
				Name:        "fickle",
				WorkspaceId: "1",
				Archived:    false,
				UserId:      "1",
			},
			args: args{
				r: r,
				p: issues.UpdateBoardParam{},
			},
			want:    issues.Board{},
			wantErr: true,
		},
		{
			name: "fail: not found",
			fields: fields{
				Id:          "2",
				Name:        "fickle",
				WorkspaceId: "1",
				Archived:    false,
				UserId:      "1",
			},
			args: args{
				r: r,
				p: issues.UpdateBoardParam{
					Name: func() *string { s := "updated"; return &s }(),
				},
			},
			want:    issues.Board{},
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
			got, err := b.Update(tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Board.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
