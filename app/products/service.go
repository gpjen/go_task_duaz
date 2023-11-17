package products

import (
	"fmt"
	"user-product-management/app/users"
)

type Service interface {
	Create(product ProductRegister, creatorID uint) (Product, error)
	FindAll() ([]Product, error)
	FindByID(ID uint) (Product, error)
	UpdateProduct(input ProductUpdate, productID uint, userRequester users.UserContex) (Product, error)
	DeleteProduct(productID uint, userRequester users.UserContex) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input ProductRegister, creatorID uint) (Product, error) {

	product := Product{
		Name:        input.Name,
		Price:       input.Price,
		Stock:       input.Stock,
		Description: input.Description,

		CreatorID: creatorID,
	}

	newProduct, err := s.repository.Create(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) FindAll() ([]Product, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) FindByID(ID uint) (Product, error) {
	product, err := s.repository.FindByID(ID)
	if err != nil || product.ID == 0 {
		return product, fmt.Errorf("product not found")
	}

	return product, nil

}

func (s *service) UpdateProduct(input ProductUpdate, productID uint, userRequester users.UserContex) (Product, error) {

	product, err := s.repository.FindByID(productID)
	if err != nil || product.ID == 0 {
		return product, fmt.Errorf("product not found")
	}

	if product.CreatorID != userRequester.ID && userRequester.Role != "admin" {
		return product, fmt.Errorf("you are not allowed to update this product")
	}

	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock
	product.Description = input.Description

	updatedProduct, err := s.repository.Update(product)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil
}

func (s *service) DeleteProduct(productID uint, userRequester users.UserContex) (Product, error) {

	product, err := s.repository.FindByID(productID)
	if err != nil || product.ID == 0 {
		return product, fmt.Errorf("product not found")
	}

	if product.CreatorID != userRequester.ID && userRequester.Role != "admin" {
		return product, fmt.Errorf("you are not allowed to delete this product")
	}

	deletedProduct, err := s.repository.Delete(product)
	if err != nil {
		return deletedProduct, err
	}

	return deletedProduct, nil
}
