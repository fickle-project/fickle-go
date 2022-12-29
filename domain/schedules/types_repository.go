package schedules

type IRepository interface {
	AddSchedule(Schedule) (Schedule, error)
	UpdateSchedule(IdSchedule, UpdateScheduleParam) (Schedule, error)
	RemoveSchedule(IdSchedule) error

	AddLog(Log) (Log, error)
	UpdateLog(IdLog, UpdateLogParam) (Log, error)
	RemoveLog(IdLog) error
}

type IQueryService interface {
	FindSchedule(IdSchedule, QueryScheduleParam) (ScheduleWithEmbedDatas, error)
	FindSchedules(QueryScheduleParam) ([]ScheduleWithEmbedDatas, error)

	FindLog(IdLog, QueryLogParam) (LogWithEmbedDatas, error)
	FindLogs(QueryLogParam) ([]LogWithEmbedDatas, error)
}
