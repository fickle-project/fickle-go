package issues

// Remove implements iIssue
func (i *Issue) Remove(r IRepository) error {
	return r.RemoveIssue(i.Id)
}
