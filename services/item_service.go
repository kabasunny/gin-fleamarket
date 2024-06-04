package services

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint, userId uint) (*models.Item, error)
	Create(createItemInuput dto.CreateItemInuput, useId uint) (*models.Item, error)
	Update(itemId uint, useId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
	Delete(itemId uint, useId uint) error
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

func (s *ItemService) FindById(itemId uint, useId uint) (*models.Item, error){
	return s.repository.FindById(itemId, useId)
}

func (s *ItemService) Create(createItemInuput dto.CreateItemInuput, userId uint) (*models.Item, error){
	newItem := models.Item{
		Name: createItemInuput.Name,
		Price: createItemInuput.Price,
		Description: createItemInuput.Description,
		SoldOut: false,
		UserID: userId,
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) Update(itemId uint, userId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error){
	targetItem, err := s.FindById(itemId, userId)
	if err != nil {
		return nil, err
	}
	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}
	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}
	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}
	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}
	return s.repository.Update(*targetItem)
}

func (s *ItemService) Delete(itemId uint, useId uint) error{
	return s.repository.Delete(itemId, useId)
}