package schedules_test

import (
	"fickle/domain/issues"
	"fickle/domain/schedules"
	"fickle/infrastructure/datasource/inmemory"
	"reflect"
	"testing"
	"time"
)

func Test_timeTableService_FindLog(t *testing.T) {
	ri := inmemory.NewRepositoryIssues()
	r := inmemory.NewRepositorySchedules(ri)
	q := inmemory.NewQueryServiceSchedules(r, ri)
	ri.CreateIssue(issues.Issue{
		Id:          "1",
		Name:        "issue 1",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})
	ri.CreateIssue(issues.Issue{
		Id:          "2",
		Name:        "issue 2",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})
	r.AddLog(schedules.Log{
		Id:      "1",
		Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
		IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
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

	type args struct {
		q  schedules.IQueryService
		id schedules.IdLog
		p  schedules.QueryLogParam
	}
	tests := []struct {
		name    string
		args    args
		want    schedules.LogWithEmbedDatas
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				q:  q,
				id: "1",
				p:  schedules.QueryLogParam{},
			},
			want: schedules.LogWithEmbedDatas{
				Log: schedules.Log{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
				Issue: issues.IssueWithEmbedDatas{
					Issue: issues.Issue{},
				},
			},
			wantErr: false,
		},
		{
			name: "ok: embed 'Issue'",
			args: args{
				q:  q,
				id: "1",
				p:  schedules.QueryLogParam{Embed: schedules.QueryLogParamEmbed{Issue: true}},
			},
			want: schedules.LogWithEmbedDatas{
				Log: schedules.Log{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
				Issue: issues.IssueWithEmbedDatas{
					Issue: issues.Issue{
						Id:          "1",
						Name:        "issue 1",
						BoardId:     "1",
						ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
						Order:       "1",
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ok: embed 'Issue'",
			args: args{
				q:  q,
				id: "2",
				p:  schedules.QueryLogParam{Embed: schedules.QueryLogParamEmbed{Issue: true}},
			},
			want: schedules.LogWithEmbedDatas{
				Log: schedules.Log{
					Id:      "2",
					Name:    func() *string { var s *string = new(string); *s = "log 2"; return s }(),
					IssueId: nil,
					Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
				Issue: issues.IssueWithEmbedDatas{
					Issue: issues.Issue{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := schedules.NewService()
			got, err := tr.FindLog(tt.args.q, tt.args.id, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("timeTableService.FindLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timeTableService.FindLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeTableService_FindLogs(t *testing.T) {
	ri := inmemory.NewRepositoryIssues()
	r := inmemory.NewRepositorySchedules(ri)
	q := inmemory.NewQueryServiceSchedules(r, ri)
	ri.CreateIssue(issues.Issue{
		Id:          "1",
		Name:        "issue 1",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})
	ri.CreateIssue(issues.Issue{
		Id:          "2",
		Name:        "issue 2",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})
	r.AddLog(schedules.Log{
		Id:      "1",
		Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
		IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
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

	type args struct {
		q schedules.IQueryService
		p schedules.QueryLogParam
	}
	tests := []struct {
		name    string
		args    args
		want    []schedules.LogWithEmbedDatas
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				q: q,
				p: schedules.QueryLogParam{},
			},
			want: []schedules.LogWithEmbedDatas{{
				Log: schedules.Log{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}, {
				Log: schedules.Log{
					Id:      "2",
					Name:    func() *string { var s *string = new(string); *s = "log 2"; return s }(),
					IssueId: nil,
					Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			},
			},
			wantErr: false,
		},
		{
			name: "ok: embed 'Issue'",
			args: args{
				q: q,
				p: schedules.QueryLogParam{
					Embed: schedules.QueryLogParamEmbed{Issue: true},
				},
			},
			want: []schedules.LogWithEmbedDatas{{
				Log: schedules.Log{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
				Issue: issues.IssueWithEmbedDatas{
					Issue: issues.Issue{
						Id:          "1",
						Name:        "issue 1",
						BoardId:     "1",
						ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
						Order:       "1",
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
			}, {
				Log: schedules.Log{
					Id:      "2",
					Name:    func() *string { var s *string = new(string); *s = "log 2"; return s }(),
					IssueId: nil,
					Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}},
			wantErr: false,
		},
		{
			name: "ok: query 'Name",
			args: args{
				q: q,
				p: schedules.QueryLogParam{
					Embed: schedules.QueryLogParamEmbed{Issue: true},
					Name:  func() *string { var s *string = new(string); *s = "log 1"; return s }(),
				},
			},
			want: []schedules.LogWithEmbedDatas{{
				Log: schedules.Log{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
				Issue: issues.IssueWithEmbedDatas{
					Issue: issues.Issue{
						Id:          "1",
						Name:        "issue 1",
						BoardId:     "1",
						ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
						Order:       "1",
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
			}},
			wantErr: false,
		},
		{
			name: "ok: query 'NameContain",
			args: args{
				q: q,
				p: schedules.QueryLogParam{
					Embed:       schedules.QueryLogParamEmbed{Issue: true},
					NameContain: func() *string { var s *string = new(string); *s = "log"; return s }(),
				},
			},
			want: []schedules.LogWithEmbedDatas{{
				Log: schedules.Log{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
				Issue: issues.IssueWithEmbedDatas{
					Issue: issues.Issue{
						Id:          "1",
						Name:        "issue 1",
						BoardId:     "1",
						ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
						Order:       "1",
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
			}, {
				Log: schedules.Log{
					Id:      "2",
					Name:    func() *string { var s *string = new(string); *s = "log 2"; return s }(),
					IssueId: nil,
					Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}},
			wantErr: false,
		},
		{
			name: "ok: query 'IssueId",
			args: args{
				q: q,
				p: schedules.QueryLogParam{
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
				},
			},
			want: []schedules.LogWithEmbedDatas{{
				Log: schedules.Log{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}},
			wantErr: false,
		},
		{
			name: "ok: query 'IssueId",
			args: args{
				q: q,
				p: schedules.QueryLogParam{
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "2"; return s }(),
				},
			},
			want:    []schedules.LogWithEmbedDatas{},
			wantErr: false,
		},
		{
			name: "ok: query 'From' and 'To'",
			args: args{
				q: q,
				p: schedules.QueryLogParam{
					From: func() *time.Time { t := time.Date(2022, time.December, 29, 0, 0, 0, 0, time.UTC); return &t }(),
					To:   func() *time.Time { t := time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC); return &t }(),
				},
			},
			want:    []schedules.LogWithEmbedDatas{},
			wantErr: false,
		},
		{
			name: "ok: query 'From' and 'To'",
			args: args{
				q: q,
				p: schedules.QueryLogParam{
					From: func() *time.Time { t := time.Date(2022, time.December, 29, 0, 0, 0, 0, time.UTC); return &t }(),
					To:   func() *time.Time { t := time.Date(2022, time.December, 29, 9, 30, 0, 0, time.UTC); return &t }(),
				},
			},
			want: []schedules.LogWithEmbedDatas{{
				Log: schedules.Log{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}},
			wantErr: false,
		},
		{
			name: "ok: query 'From' and 'To'",
			args: args{
				q: q,
				p: schedules.QueryLogParam{
					From: func() *time.Time { t := time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC); return &t }(),
					To:   func() *time.Time { t := time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC); return &t }(),
				},
			},
			want: []schedules.LogWithEmbedDatas{{
				Log: schedules.Log{
					Id:      "2",
					Name:    func() *string { var s *string = new(string); *s = "log 2"; return s }(),
					IssueId: nil,
					Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := schedules.NewService()
			got, err := tr.FindLogs(tt.args.q, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("timeTableService.FindLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timeTableService.FindLogs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeTableService_FindSchedule(t *testing.T) {
	ri := inmemory.NewRepositoryIssues()
	r := inmemory.NewRepositorySchedules(ri)
	q := inmemory.NewQueryServiceSchedules(r, ri)
	ri.CreateIssue(issues.Issue{
		Id:          "1",
		Name:        "issue 1",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})
	ri.CreateIssue(issues.Issue{
		Id:          "2",
		Name:        "issue 2",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})
	r.AddSchedule(schedules.Schedule{
		Id:      "1",
		Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
		IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
		Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
		End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
		UserId:  "1",
	})
	r.AddSchedule(schedules.Schedule{
		Id:      "2",
		Name:    func() *string { var s *string = new(string); *s = "log 2"; return s }(),
		IssueId: nil,
		Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
		End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
		UserId:  "1",
	})

	type args struct {
		q  schedules.IQueryService
		id schedules.IdSchedule
		p  schedules.QueryScheduleParam
	}
	tests := []struct {
		name    string
		args    args
		want    schedules.ScheduleWithEmbedDatas
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				q:  q,
				id: "1",
				p:  schedules.QueryScheduleParam{},
			},
			want: schedules.ScheduleWithEmbedDatas{
				Schedule: schedules.Schedule{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			},
			wantErr: false,
		},
		{
			name: "ok: embed 'Issue'",
			args: args{
				q:  q,
				id: "1",
				p: schedules.QueryScheduleParam{
					Embed: schedules.QueryScheduleParamEmbed{Issue: true},
				},
			},
			want: schedules.ScheduleWithEmbedDatas{
				Schedule: schedules.Schedule{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "log 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
				Issue: issues.IssueWithEmbedDatas{
					Issue: issues.Issue{
						Id:          "1",
						Name:        "issue 1",
						BoardId:     "1",
						ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
						Order:       "1",
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ok: embed 'Issue'",
			args: args{
				q:  q,
				id: "2",
				p: schedules.QueryScheduleParam{
					Embed: schedules.QueryScheduleParamEmbed{Issue: true},
				},
			},
			want: schedules.ScheduleWithEmbedDatas{
				Schedule: schedules.Schedule{
					Id:      "2",
					Name:    func() *string { var s *string = new(string); *s = "log 2"; return s }(),
					IssueId: nil,
					Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := schedules.NewService()
			got, err := tr.FindSchedule(tt.args.q, tt.args.id, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("timeTableService.FindSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timeTableService.FindSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeTableService_FindSchedules(t *testing.T) {
	ri := inmemory.NewRepositoryIssues()
	r := inmemory.NewRepositorySchedules(ri)
	q := inmemory.NewQueryServiceSchedules(r, ri)
	ri.CreateIssue(issues.Issue{
		Id:          "1",
		Name:        "issue 1",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})
	ri.CreateIssue(issues.Issue{
		Id:          "2",
		Name:        "issue 2",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})
	r.AddSchedule(schedules.Schedule{
		Id:      "1",
		Name:    func() *string { var s *string = new(string); *s = "schedules 1"; return s }(),
		IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
		Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
		End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
		UserId:  "1",
	})
	r.AddSchedule(schedules.Schedule{
		Id:      "2",
		Name:    func() *string { var s *string = new(string); *s = "schedules 2"; return s }(),
		IssueId: nil,
		Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
		End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
		UserId:  "1",
	})

	type args struct {
		q schedules.IQueryService
		p schedules.QueryScheduleParam
	}
	tests := []struct {
		name    string
		args    args
		want    []schedules.ScheduleWithEmbedDatas
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				q: q,
				p: schedules.QueryScheduleParam{},
			},
			want: []schedules.ScheduleWithEmbedDatas{{
				Schedule: schedules.Schedule{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "schedules 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}, {
				Schedule: schedules.Schedule{
					Id:      "2",
					Name:    func() *string { var s *string = new(string); *s = "schedules 2"; return s }(),
					IssueId: nil,
					Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}},
			wantErr: false,
		},
		{
			name: "ok: embed 'Issue'",
			args: args{
				q: q,
				p: schedules.QueryScheduleParam{Embed: schedules.QueryScheduleParamEmbed{Issue: true}},
			},
			want: []schedules.ScheduleWithEmbedDatas{{
				Schedule: schedules.Schedule{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "schedules 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
				Issue: issues.IssueWithEmbedDatas{
					Issue: issues.Issue{
						Id:          "1",
						Name:        "issue 1",
						BoardId:     "1",
						ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
						Order:       "1",
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
			}, {
				Schedule: schedules.Schedule{
					Id:      "2",
					Name:    func() *string { var s *string = new(string); *s = "schedules 2"; return s }(),
					IssueId: nil,
					Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}},
			wantErr: false,
		},
		{
			name: "ok: query 'Name'",
			args: args{
				q: q,
				p: schedules.QueryScheduleParam{
					Embed: schedules.QueryScheduleParamEmbed{Issue: true},
					Name:  func() *string { s := "schedules 1"; return &s }(),
				},
			},
			want: []schedules.ScheduleWithEmbedDatas{{
				Schedule: schedules.Schedule{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "schedules 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
				Issue: issues.IssueWithEmbedDatas{
					Issue: issues.Issue{
						Id:          "1",
						Name:        "issue 1",
						BoardId:     "1",
						ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
						Order:       "1",
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
			}},
			wantErr: false,
		},
		{
			name: "ok: query 'NameContain'",
			args: args{
				q: q,
				p: schedules.QueryScheduleParam{
					Embed:       schedules.QueryScheduleParamEmbed{Issue: true},
					NameContain: func() *string { s := "schedules"; return &s }(),
				},
			},
			want: []schedules.ScheduleWithEmbedDatas{{
				Schedule: schedules.Schedule{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "schedules 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
				Issue: issues.IssueWithEmbedDatas{
					Issue: issues.Issue{
						Id:          "1",
						Name:        "issue 1",
						BoardId:     "1",
						ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
						Order:       "1",
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
			}, {
				Schedule: schedules.Schedule{
					Id:      "2",
					Name:    func() *string { var s *string = new(string); *s = "schedules 2"; return s }(),
					IssueId: nil,
					Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}},
			wantErr: false,
		},
		{
			name: "ok: query 'NameContain'",
			args: args{
				q: q,
				p: schedules.QueryScheduleParam{
					Embed:       schedules.QueryScheduleParamEmbed{Issue: true},
					NameContain: func() *string { s := "2"; return &s }(),
				},
			},
			want: []schedules.ScheduleWithEmbedDatas{{
				Schedule: schedules.Schedule{
					Id:      "2",
					Name:    func() *string { var s *string = new(string); *s = "schedules 2"; return s }(),
					IssueId: nil,
					Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}},
			wantErr: false,
		},
		{
			name: "ok: query 'IssueId'",
			args: args{
				q: q,
				p: schedules.QueryScheduleParam{
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
				},
			},
			want: []schedules.ScheduleWithEmbedDatas{{
				Schedule: schedules.Schedule{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "schedules 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}},
			wantErr: false,
		},
		{
			name: "ok: query 'IssueId'",
			args: args{
				q: q,
				p: schedules.QueryScheduleParam{
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "2"; return s }(),
				},
			},
			want:    []schedules.ScheduleWithEmbedDatas{},
			wantErr: false,
		},
		{
			name: "ok: query 'From' and 'To'",
			args: args{
				q: q,
				p: schedules.QueryScheduleParam{
					From: func() *time.Time { t := time.Date(2022, time.December, 29, 0, 0, 0, 0, time.UTC); return &t }(),
					To:   func() *time.Time { t := time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC); return &t }(),
				},
			},
			want:    []schedules.ScheduleWithEmbedDatas{},
			wantErr: false,
		},
		{
			name: "ok: query 'From' and 'To'",
			args: args{
				q: q,
				p: schedules.QueryScheduleParam{
					From: func() *time.Time { t := time.Date(2022, time.December, 29, 0, 0, 0, 0, time.UTC); return &t }(),
					To:   func() *time.Time { t := time.Date(2022, time.December, 29, 9, 30, 0, 0, time.UTC); return &t }(),
				},
			},
			want: []schedules.ScheduleWithEmbedDatas{{
				Schedule: schedules.Schedule{
					Id:      "1",
					Name:    func() *string { var s *string = new(string); *s = "schedules 1"; return s }(),
					IssueId: func() *issues.IdIssue { var s *issues.IdIssue = new(issues.IdIssue); *s = "1"; return s }(),
					Start:   time.Date(2022, time.December, 29, 9, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}},
			wantErr: false,
		},
		{
			name: "ok: query 'From' and 'To'",
			args: args{
				q: q,
				p: schedules.QueryScheduleParam{
					From: func() *time.Time { t := time.Date(2022, time.December, 29, 10, 0, 0, 0, time.UTC); return &t }(),
					To:   func() *time.Time { t := time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC); return &t }(),
				},
			},
			want: []schedules.ScheduleWithEmbedDatas{{
				Schedule: schedules.Schedule{
					Id:      "2",
					Name:    func() *string { var s *string = new(string); *s = "schedules 2"; return s }(),
					IssueId: nil,
					Start:   time.Date(2022, time.December, 29, 12, 0, 0, 0, time.UTC),
					End:     time.Date(2022, time.December, 29, 13, 0, 0, 0, time.UTC),
					UserId:  "1",
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := schedules.NewService()
			got, err := tr.FindSchedules(tt.args.q, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("timeTableService.FindSchedules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timeTableService.FindSchedules() = %v, want %v", got, tt.want)
			}
		})
	}
}
