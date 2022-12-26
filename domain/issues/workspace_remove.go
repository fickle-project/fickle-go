package issues

// Remove implements iWorkspace
func (w *Workspace) Remove(r IRepository) error {
	return r.RemoveWorkspcae(w.Id)
}
