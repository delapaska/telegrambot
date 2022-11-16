package repository_impl

import (
	"database/sql"
	"test/internal/repository"
)

type Repository struct {
	db              *sql.DB
	userRepository  *UserRepository
	priceRepository *PriceRepository
}

func New(db *sql.DB) repository.Repository {
	repo := new(Repository)

	repo.db = db

	return repo
}

func (s *Repository) User() repository.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		repository: s,
	}
	return s.userRepository
}

func (s *Repository) Price() repository.PriceRepository {
	if s.priceRepository != nil {
		return s.priceRepository
	}

	s.priceRepository = &PriceRepository{
		repository: s,
	}
	return s.priceRepository
}
