package schedules_test

import (
	"fickle/domain/issues"
	"fickle/domain/schedules"
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"testing"
	"time"
)

func TestLog_Remove(t *testing.T) {
	r := inmemory.NewRepositorySchedules(inmemory.NewRepositoryIssues())
	r.AddLog(schedules.Log{
		Id:      "1",
		Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
		IssueId: nil,
		Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
		End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
		UserId:  "1",
	})
	r.AddLog(schedules.Log{
		Id:      "2",
		Name:    func() *string { var s *string = new(string); *s = "log 2"; return s }(),
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
		r schedules.IRepository
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
				Id:      "1",
				Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
				IssueId: nil,
				Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			args: args{
				r: r,
			},
			wantErr: false,
		},
		{
			name: "ok",
			fields: fields{
				Id:      "2",
				Name:    func() *string { var s *string = new(string); *s = "log 2"; return s }(),
				IssueId: nil,
				Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			args: args{
				r: r,
			},
			wantErr: false,
		},
		{
			name: "fail: not found",
			fields: fields{
				Id:      "3",
				Name:    func() *string { var s *string = new(string); *s = "log 3"; return s }(),
				IssueId: nil,
				Start:   time.Date(2022, time.December, 29, 14, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 15, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			args: args{
				r: r,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &schedules.Log{
				Id:      tt.fields.Id,
				Name:    tt.fields.Name,
				IssueId: tt.fields.IssueId,
				Start:   tt.fields.Start,
				End:     tt.fields.End,
				UserId:  tt.fields.UserId,
			}
			if err := l.Remove(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Log.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
