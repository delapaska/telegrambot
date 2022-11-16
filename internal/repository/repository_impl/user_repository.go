package repository_impl

type UserRepository struct {
	repository *Repository
}

func (r *UserRepository) Delete() error {
	query := `DELETE FROM users WHERE id > 0; `
	_, err := r.repository.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
