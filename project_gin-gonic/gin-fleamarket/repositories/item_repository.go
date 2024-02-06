package repositories

import (
	"errors"
	"gin-fleamarket/models"
	"gorm.io/gorm"
)

// IItemRepository IItemRepository 인터페이스는 ItemMemoryRepository 구조체가 구현해야 하는 메서드를 정의한다.
type IItemRepository interface {
	FindAll() (*[]models.Item, error)                    // 모든 Item을 반환한다.
	FindById(itemId uint) (*models.Item, error)          // itemId에 해당하는 Item을 반환한다.
	Create(newItem models.Item) (*models.Item, error)    // Item을 생성한다.
	Update(updateItem models.Item) (*models.Item, error) // Item을 수정한다.
	Delete(itemId uint) error                            // Item을 삭제한다.
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

func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("Item not found")
}

// Create Create() 메서드는 Item을 생성한다.
func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

// Update Update() 메서드는 Item을 수정한다.
func (r *ItemMemoryRepository) Update(updateItem models.Item) (*models.Item, error) {
	for i, v := range r.items {
		if v.ID == updateItem.ID {
			r.items[i] = updateItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("Unexpected error")
}

// Delete Delete() 메서드는 Item을 삭제한다.
func (r *ItemMemoryRepository) Delete(itemId uint) error {
	for i, v := range r.items {
		if v.ID == itemId {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("Item not found")
}

type ItemRepository struct {
	db *gorm.DB
}

func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := r.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

func (r *ItemRepository) FindById(itemId uint) (*models.Item, error) {
	var item models.Item
	result := r.db.First(&item, itemId) // First 메서드가 처리 속도 면에서 Find 메서드보다 빠르다.
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("Item not found")
		}
		return nil, result.Error
	}
	return &item, nil
}

func (r *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

func (r *ItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	result := r.db.Save(&updateItem) // Save 메서드는 updateItem의 ID가 0이 아니면 updateItem의 ID에 해당하는 레코드를 업데이트하고, 0이면 새로운 레코드를 생성한다.
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}

func (r *ItemRepository) Delete(itemId uint) error {
	deleteItem, err := r.FindById(itemId)
	if err != nil {
		return err
	}
	result := r.db.Delete(&deleteItem) // Delete 메서드는 물리 삭제가 아닌 논리 삭제를 수행한다.
	//result := r.db.Unscoped().Delete(&deleteItem) // Unscoped() 메서드를 사용하면 물리 삭제를 수행한다.
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}
