package dto

type CreateItemInuput struct{
	Name string `json:"name" binding:"required,min=2"`
	Price uint `json:"price" binding:"required,min=1,max=999999"`
	Description string `json:"description"`
}

type UpdateItemInput struct{
	Name *string `json:"Name" binding:"omitnil,min=2"`
	Price *uint `json:"price" binding:"omitnil,min=1,max=999999"`
	Description *string `json:"description"`
	SoldOut *bool `json:"soldOut"`
}