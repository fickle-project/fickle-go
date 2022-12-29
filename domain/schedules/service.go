package schedules

import "fickle/domain/issues"

type iTimeTableService interface {
	NewSchedule(IFactory, IRepository, issues.IRepository, AddScheduleParam) (Schedule, error)
	FindSchedule(IQueryService, IdSchedule, QueryScheduleParam) (ScheduleWithEmbedDatas, error)
	FindSchedules(IQueryService, QueryScheduleParam) ([]ScheduleWithEmbedDatas, error)
	NewLog(IFactory, IRepository, issues.IRepository, AddLogParam) (Log, error)
	FindLog(IQueryService, IdLog, QueryLogParam) (LogWithEmbedDatas, error)
	FindLogs(IQueryService, QueryLogParam) ([]LogWithEmbedDatas, error)
}

func NewService() iTimeTableService {
	return &timeTableService{}
}

// Implements iTimeTableService
type timeTableService struct{}
