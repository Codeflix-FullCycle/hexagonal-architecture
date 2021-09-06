package application

type ProductService struct {
	ProductPersistence ProductPersistenceInterface
}

func NewProductService(Persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{
		ProductPersistence: Persistence,
	}
}

func (p *ProductService) Get(id string) (ProductsInterface, error) {
	product, err := p.ProductPersistence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) Create(name string, price float64) (ProductsInterface, error) {
	product := NewProduct()

	product.Name = name
	product.Price = price

	if _, err := product.IsValid(); err != nil {
		return &Products{}, err
	}

	result, err := p.ProductPersistence.Save(product)
	if err != nil {
		return &Products{}, err
	}

	return result, nil
}

func (p *ProductService) Enable(product ProductsInterface) (ProductsInterface, error) {
	err := product.Enable()

	if err != nil {
		return &Products{}, err
	}

	result, err := p.ProductPersistence.Save(product)

	if err != nil {
		return &Products{}, err
	}

	return result, nil
}

func (p *ProductService) Disable(product ProductsInterface) (ProductsInterface, error) {
	err := product.Disable()

	if err != nil {
		return &Products{}, err
	}

	result, err := p.ProductPersistence.Save(product)

	if err != nil {
		return &Products{}, err
	}

	return result, nil
}
