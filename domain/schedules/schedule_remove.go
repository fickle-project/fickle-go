package schedules

// Remove implements iSchedule
func (s *Schedule) Remove(r IRepository) error {
	return r.RemoveSchedule(s.Id)
}
