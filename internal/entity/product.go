package entity

import (
	"errors"
	"time"

	"github.com/andreis3/api-golang/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidID       = errors.New("invalid id")
	ErrorInvalidPrice  = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0 {
		return ErrorInvalidPrice
	}

	return nil
}
