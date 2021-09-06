package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type ProductsInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetPrice() float64
	GetStatus() string
}

type ProductServiceInterface interface {
	Get(id string) (ProductsInterface, error)
	Create(name string, price float64) (ProductsInterface, error)
	Enable(product ProductsInterface) (ProductsInterface, error)
	Disable(product ProductsInterface) (ProductsInterface, error)
}

// SOLID
// Interface segregation
// Separr uma interface única e muitas interfaces especificas

type ProductReader interface {
	Get(id string) (ProductsInterface, error)
}

type ProductWrite interface {
	Save(product ProductsInterface) (ProductsInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWrite
}

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Products struct {
	ID     string  `valid:"uuid"`
	Name   string  `valid:"required"`
	Status string  `valid:"required"`
	Price  float64 `valid:"float, optional "`
}

func NewProduct() *Products {
	return &Products{
		ID: uuid.NewV4().String(),
	}
}

func (p *Products) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != DISABLED && p.Status != ENABLED {
		return false, errors.New("the status must be enable or disable")
	}

	if p.Price < 0 {
		return false, errors.New("the price must greater or equal zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Products) Enable() error {
	if p.Price != 0 && p.Status != ENABLED {
		p.Status = ENABLED
		return nil
	}

	return errors.New("the price must be greater than zero and different from ENABLED to enable the product")
}

func (p *Products) Disable() error {

	if p.Price == 0 && p.Status != DISABLED {
		p.Status = DISABLED
		return nil
	}

	return errors.New("the price must be greater than zero and different from DISABLED to enable the product")
}

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
