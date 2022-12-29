package inmemory

import (
	"fickle/domain/errors"
	"fickle/domain/issues"
	"fickle/domain/schedules"
	"strings"
)

func NewQueryServiceSchedules(r schedules.IRepository, ri issues.IRepository) schedules.IQueryService {
	return &qSchedules{repo: r, issuesRepo: ri}
}

type qSchedules struct {
	repo       schedules.IRepository
	issuesRepo issues.IRepository
}

// FindLog implements schedules.IRepository
func (r *qSchedules) FindLog(id schedules.IdLog, q schedules.QueryLogParam) (schedules.LogWithEmbedDatas, error) {
	ll := filter(data_logs, func(l schedules.Log) bool { return l.Id == id })

	if len(ll) == 0 {
		return schedules.LogWithEmbedDatas{}, &errors.ErrNotFound{Object: "Log", Id: string(id)}
	}

	le := schedules.LogWithEmbedDatas{Log: ll[0]}

	if q.Embed.Issue && le.Log.IssueId != nil {
		issue, err := r.issuesRepo.FindIssue(*le.Log.IssueId, issues.QueryIssueParam{})
		if err != nil {
			return schedules.LogWithEmbedDatas{}, err
		}
		le.Issue = issue
	}

	return le, nil
}

// FindLogs implements schedules.IRepository
func (r *qSchedules) FindLogs(q schedules.QueryLogParam) ([]schedules.LogWithEmbedDatas, error) {
	ll := filter(data_logs, func(l schedules.Log) bool {
		selected := true
		if q.Name != nil {
			selected = selected && l.Name != nil && *l.Name == *q.Name
		}
		if q.NameContain != nil {
			selected = selected && l.Name != nil && strings.Contains(*l.Name, *q.NameContain)
		}
		if q.From != nil {
			selected = selected && l.End.After(*q.From)
		}
		if q.To != nil {
			selected = selected && l.Start.Before(*q.To)
		}
		if q.IssueId != nil {
			selected = selected && l.IssueId != nil && *l.IssueId == *q.IssueId
		}
		return selected
	})

	embeds := []schedules.LogWithEmbedDatas{}

	for _, l := range ll {
		le := schedules.LogWithEmbedDatas{Log: l}
		if q.Embed.Issue && le.Log.IssueId != nil {
			// TODO: fix N+1
			issue, err := r.issuesRepo.FindIssue(*le.Log.IssueId, issues.QueryIssueParam{})
			if err != nil {
				return nil, err
			}
			le.Issue = issue
		}
		embeds = append(embeds, le)
	}

	return embeds, nil
}

// FindSchedule implements schedules.IRepository
func (r *qSchedules) FindSchedule(id schedules.IdSchedule, q schedules.QueryScheduleParam) (schedules.ScheduleWithEmbedDatas, error) {
	ss := filter(data_schedules, func(s schedules.Schedule) bool { return s.Id == id })

	if len(ss) == 0 {
		return schedules.ScheduleWithEmbedDatas{}, &errors.ErrNotFound{Object: "Schedule", Id: string(id)}
	}

	se := schedules.ScheduleWithEmbedDatas{Schedule: ss[0]}

	if q.Embed.Issue && se.Schedule.IssueId != nil {
		issue, err := r.issuesRepo.FindIssue(*se.Schedule.IssueId, issues.QueryIssueParam{})
		if err != nil {
			return schedules.ScheduleWithEmbedDatas{}, err
		}
		se.Issue = issue
	}

	return se, nil
}

// FindSchedules implements schedules.IRepository
func (r *qSchedules) FindSchedules(q schedules.QueryScheduleParam) ([]schedules.ScheduleWithEmbedDatas, error) {
	ss := filter(data_schedules, func(s schedules.Schedule) bool {
		selected := true
		if q.Name != nil {
			selected = selected && s.Name != nil && *s.Name == *q.Name
		}
		if q.NameContain != nil {
			selected = selected && s.Name != nil && strings.Contains(*s.Name, *q.NameContain)
		}
		if q.From != nil {
			selected = selected && s.End.After(*q.From)
		}
		if q.To != nil {
			selected = selected && s.Start.Before(*q.To)
		}
		if q.IssueId != nil {
			selected = selected && s.IssueId != nil && *s.IssueId == *q.IssueId
		}
		return selected
	})

	embeds := []schedules.ScheduleWithEmbedDatas{}

	for _, s := range ss {
		se := schedules.ScheduleWithEmbedDatas{Schedule: s}
		if q.Embed.Issue && se.Schedule.IssueId != nil {
			// TODO: fix N+1
			issue, err := r.issuesRepo.FindIssue(*se.Schedule.IssueId, issues.QueryIssueParam{})
			if err != nil {
				return nil, err
			}
			se.Issue = issue
		}
		embeds = append(embeds, se)
	}

	return embeds, nil
}
