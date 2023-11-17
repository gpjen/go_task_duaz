package products

type ProductRegister struct {
	Name        string `json:"name" form:"name" validate:"required,min=3,max=100"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Stock       int    `json:"stock" form:"stock" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
}

type ProductUpdate struct {
	Name        string `json:"name" form:"name" validate:"required,min=3,max=100"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Stock       int    `json:"stock" form:"stock" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
}
