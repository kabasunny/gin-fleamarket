package services

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(createItemInuput dto.CreateItemInuput) (*models.Item, error)
}

type ItemService struct{
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error){
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint) (*models.Item, error){
	return s.repository.FindById(itemId)
}

func (s *ItemService) Create(createItemInuput dto.CreateItemInuput) (*models.Item, error){
	newItem := models.Item{
		Name: createItemInuput.Name,
		Price: createItemInuput.Price,
		Description: createItemInuput.Description,
		SoldOut: false,
	}
	return s.repository.Create(newItem)
}