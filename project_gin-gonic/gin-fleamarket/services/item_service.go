package services

import (
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

// IItemService IItemService 인터페이스는 ItemService 구조체가 구현해야 하는 메서드를 정의한다.
type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
}

// ItemService ItemService 구조체는 IItemService 인터페이스를 구현한다.
type ItemService struct {
	repository repositories.IItemRepository
}

// NewItemService NewItemService() 메서드는 ItemService 구조체를 생성한다.
func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

// FindAll FindAll() 메서드는 모든 Item을 반환한다.
func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}
