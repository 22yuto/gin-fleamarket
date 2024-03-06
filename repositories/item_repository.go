package repositories

import "gin-fleamarket/models"

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}

// factory関数
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	// ItemMemoryRepositoryの参照を返す
	return &ItemMemoryRepository{items: items}
}

// レシーバーはオブジェクト指向言語のthisやselfのようなもの
// メソッド内でその型のインスタンスにアクセスするために使用される
func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}
