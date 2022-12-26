package issues

// Remove implements iBoard
func (b *Board) Remove(r IRepository) error {
	return r.RemoveBoard(b.Id)
}
