package schedules_test

import (
	"fickle/domain/issues"
	"fickle/domain/schedules"
	"fickle/infrastructure/datasource/inmemory"
	"reflect"
	"testing"
	"time"
)

func Test_timeTableService_NewSchedule(t *testing.T) {
	f := inmemory.NewFactorySchedules()
	ri := inmemory.NewRepositoryIssues()
	r := inmemory.NewRepositorySchedules(ri)
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

	type args struct {
		f  schedules.IFactory
		r  schedules.IRepository
		ri issues.IRepository
		p  schedules.AddScheduleParam
	}
	tests := []struct {
		name    string
		args    args
		want    schedules.Schedule
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				f:  f,
				r:  r,
				ri: ri,
				p: schedules.AddScheduleParam{
					Name:    func() *string { var s *string = new(string); *s = "schedule 1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					IssueId: func() *issues.IdIssue { var id issues.IdIssue = "1"; return &id }(),
					UserId:  "1",
				},
			},
			want: schedules.Schedule{
				Name:    func() *string { var s *string = new(string); *s = "schedule 1"; return s }(),
				IssueId: func() *issues.IdIssue { var id issues.IdIssue = "1"; return &id }(),
				Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			wantErr: false,
		},
		{
			name: "ok",
			args: args{
				f:  f,
				r:  r,
				ri: ri,
				p: schedules.AddScheduleParam{
					Name:    func() *string { var s *string = new(string); return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					IssueId: func() *issues.IdIssue { var id issues.IdIssue = "1"; return &id }(),
					UserId:  "1",
				},
			},
			want: schedules.Schedule{
				Name:    nil,
				IssueId: func() *issues.IdIssue { var id issues.IdIssue = "1"; return &id }(),
				Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
				End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
				UserId:  "1",
			},
			wantErr: false,
		},
		{
			name: "failed to validate: 'Name' empty",
			args: args{
				f:  f,
				r:  r,
				ri: ri,
				p: schedules.AddScheduleParam{
					Name:    func() *string { var s *string = new(string); return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					IssueId: nil,
					UserId:  "1",
				},
			},
			want:    schedules.Schedule{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Name' empty",
			args: args{
				f:  f,
				r:  r,
				ri: ri,
				p: schedules.AddScheduleParam{
					Name:    nil,
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					IssueId: nil,
					UserId:  "1",
				},
			},
			want:    schedules.Schedule{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'End' must be after 'Start'",
			args: args{
				f:  f,
				r:  r,
				ri: ri,
				p: schedules.AddScheduleParam{
					Name:    func() *string { var s *string = new(string); *s = "schedule 1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 8, 0, 0, 0, time.UTC),
					IssueId: nil,
					UserId:  "1",
				},
			},
			want:    schedules.Schedule{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'IssueId' issue not found",
			args: args{
				f:  f,
				r:  r,
				ri: ri,
				p: schedules.AddScheduleParam{
					Name:    func() *string { var s *string = new(string); *s = "schedule 1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					IssueId: func() *issues.IdIssue { var id issues.IdIssue = "0"; return &id }(),
					UserId:  "1",
				},
			},
			want:    schedules.Schedule{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := schedules.NewService()
			got, err := tr.NewSchedule(tt.args.f, tt.args.r, tt.args.ri, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("timeTableService.NewSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.Id = got.Id
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timeTableService.NewSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}
