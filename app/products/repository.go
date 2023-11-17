package products

import "gorm.io/gorm"

type Repository interface {
	Create(product Product) (Product, error)
	FindAll() ([]Product, error)
	FindByID(ID uint) (Product, error)
	Update(product Product) (Product, error)
	Delete(product Product) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) FindByID(ID uint) (Product, error) {
	var product Product
	err := r.db.First(&product, ID).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Delete(product Product) (Product, error) {
	err := r.db.Delete(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
