package repository_impl

import "test/internal/repository/models"

type PriceRepository struct {
	repository *Repository
}

func (r *PriceRepository) GetAll() ([]models.Price, error) {
	query := `SELECT id, price, wallet, exchange FROM price; `
	rows, err := r.repository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	answer := make([]models.Price, 0)
	for rows.Next() {
		s := models.Price{}
		err := rows.Scan(
			&s.ID,
			&s.Price,
			&s.Wallet,
			&s.Exchange,
		)
		if err != nil {
			return nil, err
		}
		answer = append(answer, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return answer, nil
}
