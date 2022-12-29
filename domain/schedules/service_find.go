package schedules

// FindLog implements iTimeTableService
func (*timeTableService) FindLog(q IQueryService, id IdLog, p QueryLogParam) (LogWithEmbedDatas, error) {
	return q.FindLog(id, p)
}

// FindLogs implements iTimeTableService
func (*timeTableService) FindLogs(q IQueryService, p QueryLogParam) ([]LogWithEmbedDatas, error) {
	return q.FindLogs(p)
}

// FindSchedule implements iTimeTableService
func (*timeTableService) FindSchedule(q IQueryService, id IdSchedule, p QueryScheduleParam) (ScheduleWithEmbedDatas, error) {
	return q.FindSchedule(id, p)
}

// FindSchedules implements iTimeTableService
func (*timeTableService) FindSchedules(q IQueryService, p QueryScheduleParam) ([]ScheduleWithEmbedDatas, error) {
	return q.FindSchedules(p)
}
