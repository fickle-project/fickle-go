package inmemory

import (
	"fickle/domain/schedules"

	"github.com/google/uuid"
)

func NewFactorySchedules() schedules.IFactory {
	return &fSchedules{}
}

type fSchedules struct{}

// NewLogId implements schedules.IFactory
func (*fSchedules) NewLogId(r schedules.IRepository) (schedules.IdLog, error) {
	return schedules.IdLog(uuid.NewString()), nil
}

// NewScheduleId implements schedules.IFactory
func (*fSchedules) NewScheduleId(r schedules.IRepository) (schedules.IdSchedule, error) {
	return schedules.IdSchedule(uuid.NewString()), nil
}
