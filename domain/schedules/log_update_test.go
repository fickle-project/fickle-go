package schedules_test

import (
	"fickle/domain/issues"
	"fickle/domain/schedules"
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"reflect"
	"testing"
	"time"
)

func TestLog_Update(t *testing.T) {
	ri := inmemory.NewRepositoryIssues()
	r := inmemory.NewRepositorySchedules(ri)
	q := inmemory.NewQueryServiceSchedules(r, ri)
	ri.CreateWorkspace(issues.Workspace{
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
			},
		},
		Archived: false,
		UserId:   "1",
	})
	ri.CreateBoard(issues.Board{
		Id:          "1",
		Name:        "fickle",
		WorkspaceId: "1",
		Archived:    false,
		UserId:      "1",
	})
	ri.CreateIssue(issues.Issue{
		Id:          "1",
		Name:        "issue 1",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})
	r.AddLog(schedules.Log{
		Id:      "1",
		Name:    func() *string { var s *string = new(string); *s = "schedule 1"; return s }(),
		IssueId: nil,
		Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
		End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
		UserId:  "1",
	})
	r.AddLog(schedules.Log{
		Id:      "2",
		Name:    func() *string { var s *string = new(string); *s = "schedule 2"; return s }(),
		IssueId: nil,
		Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
		End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
		UserId:  "1",
	})

	type fields struct {
		Id      schedules.IdLog
		Name    *string
		IssueId *issues.IdIssue
		Start   time.Time
		End     time.Time
		UserId  users.IdUser
	}
	type args struct {
		r  schedules.IRepository
		q  schedules.IQueryService
		ri issues.IRepository
		p  schedules.UpdateLogParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    schedules.Log
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Id:      "1",
				Name:    func() *string { var s *string = new(string); *s = "schedule 1"; return s }(),
				IssueId: nil,
				Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			args: args{
				r:  r,
				q:  q,
				ri: ri,
				p: schedules.UpdateLogParam{
					Name:    func() **string { var s *string = new(string); *s = "updated"; return &s }(),
					Start:   func() *time.Time { t := time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC); return &t }(),
					End:     func() *time.Time { t := time.Date(2022, time.December, 29, 11, 0, 0, 0, time.UTC); return &t }(),
					IssueId: func() **issues.IdIssue { var id *issues.IdIssue = new(issues.IdIssue); *id = "1"; return &id }(),
				},
			},
			want: schedules.Log{
				Id:      "1",
				Name:    func() *string { var s *string = new(string); *s = "updated"; return s }(),
				IssueId: func() *issues.IdIssue { var id issues.IdIssue = "1"; return &id }(),
				Start:   time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 11, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			wantErr: false,
		},
		{
			name: "ok",
			fields: fields{
				Id:      "1",
				Name:    func() *string { var s *string = new(string); *s = "updated"; return s }(),
				IssueId: func() *issues.IdIssue { var id issues.IdIssue = "1"; return &id }(),
				Start:   time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 11, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			args: args{
				r:  r,
				q:  q,
				ri: ri,
				p: schedules.UpdateLogParam{
					IssueId: func() **issues.IdIssue { var id *issues.IdIssue = nil; return &id }(),
				},
			},
			want: schedules.Log{
				Id:      "1",
				Name:    func() *string { var s *string = new(string); *s = "updated"; return s }(),
				IssueId: nil,
				Start:   time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 11, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			wantErr: false,
		},
		{
			name: "failed to validate: 'Name' cannot be empty",
			fields: fields{
				Id:      "1",
				Name:    func() *string { var s *string = new(string); *s = "updated"; return s }(),
				IssueId: nil,
				Start:   time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 11, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			args: args{
				r:  r,
				q:  q,
				ri: ri,
				p: schedules.UpdateLogParam{
					Name: func() **string { var s *string = new(string); *s = ""; return &s }(),
				},
			},
			want:    schedules.Log{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Name' cannot be empty",
			fields: fields{
				Id:      "1",
				Name:    func() *string { var s *string = new(string); *s = "updated"; return s }(),
				IssueId: nil,
				Start:   time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 11, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			args: args{
				r:  r,
				q:  q,
				ri: ri,
				p: schedules.UpdateLogParam{
					Name: func() **string { var s *string = nil; return &s }(),
				},
			},
			want:    schedules.Log{},
			wantErr: true,
		},
		{
			name: "ok",
			fields: fields{
				Id:      "1",
				Name:    func() *string { var s *string = new(string); *s = "updated"; return s }(),
				IssueId: nil,
				Start:   time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 11, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			args: args{
				r:  r,
				q:  q,
				ri: ri,
				p: schedules.UpdateLogParam{
					Name:    func() **string { var s *string = nil; return &s }(),
					IssueId: func() **issues.IdIssue { var id *issues.IdIssue = new(issues.IdIssue); *id = "1"; return &id }(),
				},
			},
			want: schedules.Log{
				Id:      "1",
				Name:    nil,
				IssueId: func() *issues.IdIssue { var id *issues.IdIssue = new(issues.IdIssue); *id = "1"; return id }(),
				Start:   time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 11, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			wantErr: false,
		},
		{
			name: "failed to validate: 'Name' cannot be empty",
			fields: fields{
				Id:      "1",
				Name:    func() *string { var s *string = new(string); *s = "updated"; return s }(),
				IssueId: nil,
				Start:   time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 11, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			args: args{
				r:  r,
				q:  q,
				ri: ri,
				p: schedules.UpdateLogParam{
					Name:    func() **string { var s *string = nil; return &s }(),
					IssueId: func() **issues.IdIssue { var id *issues.IdIssue = nil; return &id }(),
				},
			},
			want:    schedules.Log{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'IssueId' issue not found",
			fields: fields{
				Id:      "1",
				Name:    func() *string { var s *string = new(string); *s = "updated"; return s }(),
				IssueId: nil,
				Start:   time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 11, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			args: args{
				r:  r,
				q:  q,
				ri: ri,
				p: schedules.UpdateLogParam{
					Name:    func() **string { var s *string = nil; return &s }(),
					IssueId: func() **issues.IdIssue { var id *issues.IdIssue = new(issues.IdIssue); *id = "0"; return &id }(),
				},
			},
			want:    schedules.Log{},
			wantErr: true,
		},
		{
			name: "fail: not found",
			fields: fields{
				Id:      "3",
				Name:    func() *string { var s *string = new(string); *s = "schedule 3"; return s }(),
				IssueId: nil,
				Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			args: args{
				r:  r,
				q:  q,
				ri: ri,
				p: schedules.UpdateLogParam{
					Name:    func() **string { var s *string = nil; return &s }(),
					IssueId: func() **issues.IdIssue { var id *issues.IdIssue = new(issues.IdIssue); *id = "1"; return &id }(),
				},
			},
			want:    schedules.Log{},
			wantErr: true,
		},
		{
			name: "failed to validate: no update",
			fields: fields{
				Id:      "1",
				Name:    func() *string { var s *string = new(string); *s = "updated"; return s }(),
				IssueId: nil,
				Start:   time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 11, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			args: args{
				r:  r,
				q:  q,
				ri: ri,
				p:  schedules.UpdateLogParam{},
			},
			want:    schedules.Log{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &schedules.Log{
				Id:      tt.fields.Id,
				Name:    tt.fields.Name,
				IssueId: tt.fields.IssueId,
				Start:   tt.fields.Start,
				End:     tt.fields.End,
				UserId:  tt.fields.UserId,
			}
			got, err := s.Update(tt.args.r, tt.args.q, tt.args.ri, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Log.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Log.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
