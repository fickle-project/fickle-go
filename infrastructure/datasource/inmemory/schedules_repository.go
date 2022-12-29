package inmemory

import (
	"fickle/domain/errors"
	"fickle/domain/issues"
	"fickle/domain/schedules"
)

func NewRepositorySchedules(r issues.IRepository) schedules.IRepository {
	return &rSchedules{issuesRepo: r}
}

var data_schedules []schedules.Schedule
var data_logs []schedules.Log

type rSchedules struct {
	issuesRepo issues.IRepository
}

// AddLog implements schedules.IRepository
func (r *rSchedules) AddLog(l schedules.Log) (schedules.Log, error) {
	data_logs = append(data_logs, l)
	return l, nil
}

// AddSchedule implements schedules.IRepository
func (r *rSchedules) AddSchedule(s schedules.Schedule) (schedules.Schedule, error) {
	data_schedules = append(data_schedules, s)
	return s, nil
}

func (r *rSchedules) FindLog(id schedules.IdLog, q schedules.QueryLogParam) (schedules.LogWithEmbedDatas, error) {
	ll := filter(data_logs, func(l schedules.Log) bool { return l.Id == id })

	if len(ll) == 0 {
		return schedules.LogWithEmbedDatas{}, &errors.ErrNotFound{Object: "Log", Id: string(id)}
	}

	le := schedules.LogWithEmbedDatas{Log: ll[0]}

	if q.Embed.Issue {
		issue, err := r.issuesRepo.FindIssue(*le.Log.IssueId, issues.QueryIssueParam{})
		if err != nil {
			return schedules.LogWithEmbedDatas{}, err
		}
		le.Issue = issue
	}

	return le, nil
}

func (r *rSchedules) FindSchedule(id schedules.IdSchedule, q schedules.QueryScheduleParam) (schedules.ScheduleWithEmbedDatas, error) {
	ss := filter(data_schedules, func(s schedules.Schedule) bool { return s.Id == id })

	if len(ss) == 0 {
		return schedules.ScheduleWithEmbedDatas{}, &errors.ErrNotFound{Object: "Schedule", Id: string(id)}
	}

	se := schedules.ScheduleWithEmbedDatas{Schedule: ss[0]}

	if q.Embed.Issue {
		issue, err := r.issuesRepo.FindIssue(*se.Schedule.IssueId, issues.QueryIssueParam{})
		if err != nil {
			return schedules.ScheduleWithEmbedDatas{}, err
		}
		se.Issue = issue
	}

	return se, nil
}

// RemoveLog implements schedules.IRepository
func (r *rSchedules) RemoveLog(id schedules.IdLog) error {
	length := len(data_logs)
	data_logs = filter(data_logs, func(l schedules.Log) bool { return l.Id != id })
	if len(data_logs) == length {
		return &errors.ErrNotFound{Object: "Log", Id: string(id)}
	}
	return nil
}

// RemoveSchedule implements schedules.IRepository
func (r *rSchedules) RemoveSchedule(id schedules.IdSchedule) error {
	length := len(data_schedules)
	data_schedules = filter(data_schedules, func(s schedules.Schedule) bool { return s.Id != id })
	if len(data_schedules) == length {
		return &errors.ErrNotFound{Object: "Schedule", Id: string(id)}
	}
	return nil
}

// UpdateLog implements schedules.IRepository
func (r *rSchedules) UpdateLog(id schedules.IdLog, p schedules.UpdateLogParam) (schedules.Log, error) {
	le, err := r.FindLog(id, schedules.QueryLogParam{})
	if err != nil {
		return schedules.Log{}, err
	}
	l := le.Log

	err = r.RemoveLog(id)
	if err != nil {
		return schedules.Log{}, err
	}

	if p.Name != nil {
		l.Name = *p.Name
	}
	if p.Start != nil {
		l.Start = *p.Start
	}
	if p.End != nil {
		l.End = *p.End
	}
	if p.IssueId != nil {
		l.IssueId = *p.IssueId
	}
	return r.AddLog(l)
}

// UpdateSchedule implements schedules.IRepository
func (r *rSchedules) UpdateSchedule(id schedules.IdSchedule, p schedules.UpdateScheduleParam) (schedules.Schedule, error) {
	se, err := r.FindSchedule(id, schedules.QueryScheduleParam{})
	if err != nil {
		return schedules.Schedule{}, err
	}
	s := se.Schedule

	err = r.RemoveSchedule(id)
	if err != nil {
		return schedules.Schedule{}, err
	}

	if p.Name != nil {
		s.Name = *p.Name
	}
	if p.Start != nil {
		s.Start = *p.Start
	}
	if p.End != nil {
		s.End = *p.End
	}
	if p.IssueId != nil {
		s.IssueId = *p.IssueId
	}
	return r.AddSchedule(s)
}
