package schedules

type IFactory interface {
	NewScheduleId(r IRepository) (IdSchedule, error)
	NewLogId(r IRepository) (IdLog, error)
}
