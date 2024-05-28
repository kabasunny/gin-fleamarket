package repositories

import "gin-fleamarket/models"

type IItemRepository interface{
	FindAll() (*[]models.Item, error)
}

type ItemMemoryRepository struct{
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository{
	return &ItemMemoryRepository{items: items}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error){
	return &r.items, nil
}