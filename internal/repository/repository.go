package repository

import "database/sql"

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Delete() error {
	query := `DELETE FROM users WHERE id > 0; `
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func ()  {
	
}