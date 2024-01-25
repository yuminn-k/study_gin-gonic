package repositories

import "gin-fleamarket/models"

// IItemRepository IItemRepository 인터페이스는 ItemMemoryRepository 구조체가 구현해야 하는 메서드를 정의한다.
type IItemRepository interface {
	FindAll() (*[]models.Item, error) // 모든 Item을 반환한다.
}

// ItemMemoryRepository ItemMemoryRepository 구조체는 IItemRepository 인터페이스를 구현한다.
type ItemMemoryRepository struct {
	items []models.Item // ItemMemoryRepository 구조체는 items 필드를 가진다.
}

// NewItemMemoryRepository NewItemMemoryRepository() 메서드는 ItemMemoryRepository 구조체를 생성한다.
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

// FindAll FindAll() 메서드는 모든 Item을 반환한다.
func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}
