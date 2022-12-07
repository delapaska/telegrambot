package repository

import "test/internal/repository/models"

type Repository interface {
	User() UserRepository
	Price() PriceRepository
}

type UserRepository interface {
	Delete() error
}

type PriceRepository interface {
	GetAll() ([]models.Price, error)
}
