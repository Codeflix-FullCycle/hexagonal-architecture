package application

import "errors"

type ProductsInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetPrice() float64
	GetStatus() string
}

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

type Products struct {
	ID     string
	Name   string
	Status string
	Price  float64
}

// func (p *Products) IsValid() (bool, error) {

// }

func (p *Products) Enable() error {
	if p.Price == 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("the price most be greater than zero to enable the product")
}

// func (p *Products) Disable() error {

// }

func (p *Products) GetID() string {
	return p.ID
}

func (p *Products) GetName() string {
	return p.Name
}

func (p *Products) GetStatus() string {
	return p.Status
}

func (p *Products) GetPrice() float64 {
	return p.Price
}
