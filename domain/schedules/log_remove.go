package schedules

// Remove implements iLog
func (l *Log) Remove(r IRepository) error {
	return r.RemoveLog(l.Id)
}
