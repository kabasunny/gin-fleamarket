package repositories

import (
	"errors"
	"gin-fleamarket/models"
)

type IItemRepository interface{
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemId uint) error
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

func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error){
	for _, v := range r.items{
		if v.ID == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("Item not found")
}

func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item,error){
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

func (r *ItemMemoryRepository) Update(updateItem models.Item) (*models.Item, error){
	for i, v := range r.items{
		if v.ID == updateItem.ID {
			r.items[i] = updateItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("Unexpected error")
}

func (r *ItemMemoryRepository) Delete(itemId uint) error{
	for i, v := range r.items {
		if v.ID == itemId {
			r.items = append(r.items[:i], r.items[i+1:]... )//単純に削除する仕組みがない
			return nil
		}
	}
	return errors.New("Item not found")
}